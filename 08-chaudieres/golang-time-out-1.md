Tu gères le timeout via le `context.Context` passé à `OpenaiClient.Chat.Completions.New`, et éventuellement via l’option `WithRequestTimeout` pour le timeout par retry.[1][2]

## Timeout simple avec `context.WithTimeout`

```go
import (
    "context"
    "time"

    "github.com/openai/openai-go"
    "github.com/openai/openai-go/option"
)

func main() {
    client := openai.NewClient(
        option.WithAPIKey("YOUR_KEY"),
    )

    // Timeout global pour l’appel (incluant les retries)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    resp, err := client.Chat.Completions.New(
        ctx,
        openai.ChatCompletionNewParams{
            Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
                openai.UserMessage("Dis bonjour en français"),
            }),
            Model: openai.F(openai.ChatModelGPT4o),
        },
    )
    if err != nil {
        // si le timeout est dépassé: context.DeadlineExceeded
        // ou erreur de type *openai.OpError si c’est le client
        panic(err)
    }

    _ = resp
}
```

Les requêtes n’ont **aucun timeout par défaut**, donc sans `context.WithTimeout` l’appel peut rester bloqué longtemps.[1]

## Timeout par retry avec `WithRequestTimeout`

Si tu utilises les retries automatiques du SDK, tu peux affiner avec un timeout par tentative :[2]

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

resp, err := client.Chat.Completions.New(
    ctx,
    openai.ChatCompletionNewParams{
        Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
            openai.UserMessage("Test timeout"),
        }),
        Model: openai.F(openai.ChatModelGPT4o),
    },
    // timeout max par retry HTTP
    option.WithRequestTimeout(20*time.Second),
)
```

- `context.WithTimeout` : borne la durée totale (toutes tentatives confondues).[1]
- `option.WithRequestTimeout` : borne la durée de chaque tentative individuelle.[2]

Si tu me colles ton snippet actuel (`OpenaiClient` exact + imports), il est possible d’adapter précisément à ta version du SDK / à ta manière d’instancier le client.

[1](https://github.com/openai/openai-go)
[2](https://pkg.go.dev/github.com/LeverageSales/openai-go)
[3](https://community.openai.com/t/timeout-for-openai-chat-completion-in-python/411252)
[4](https://github.com/openai-php/client/discussions/121)
[5](https://community.glideapps.com/t/openai-timeouts/66639)
[6](https://learn.microsoft.com/en-us/answers/questions/1690782/azure-openai-gpt-4o-deployment-has-a-2-minute-hard)
[7](https://github.com/openai/openai-python/issues/549)
[8](https://github.com/spring-projects/spring-ai/issues/354)
[9](https://stackoverflow.com/questions/76491056/i-get-httpclient-timeout-error-in-c-sharp-openai-library)
[10](https://stackoverflow.com/questions/78191390/gpt-python-sdk-introduces-massive-overhead-incorrect-timeout)


Au moment du timeout, l’appel retourne une erreur (souvent `context.DeadlineExceeded` ou une erreur HTTP “timeout”) que tu peux tester et traiter comme un cas métier normal : message utilisateur, retry, log, etc.[1][2]

## Détection de l’erreur de timeout

Typiquement avec `context.WithTimeout` :

```go
resp, err := client.Chat.Completions.New(ctx, params)
if err != nil {
    // Timeout côté contexte
    if errors.Is(err, context.DeadlineExceeded) {
        // ici tu sais que c’est un timeout
        // → renvoyer une 504 HTTP, un message à l’utilisateur, etc.
        return fmt.Errorf("timeout OpenAI: %w", err)
    }

    // Timeout réseau générique
    if ne, ok := err.(net.Error); ok && ne.Timeout() {
        return fmt.Errorf("timeout réseau OpenAI: %w", err)
    }

    // Autres erreurs
    return fmt.Errorf("erreur OpenAI: %w", err)
}
```

- `errors.Is(err, context.DeadlineExceeded)` te dit que la deadline du contexte a été dépassée.[2][1]
- `net.Error` avec `Timeout()` permet de capter un timeout au niveau du client HTTP.[1]

## Que faire côté logique applicative

Une fois l’erreur de timeout détectée, tu peux par exemple :

- Retourner une réponse de repli à l’utilisateur (“Le modèle met trop de temps, réessayez plus tard”).  
- Logguer l’évènement (latence, prompt, modèle) pour ajuster temps de timeout et monitoring.[3]
- Lancer un retry avec backoff exponentiel si ton cas d’usage le permet (au-dessus du SDK qui fait déjà quelques retries).[4][5]

Si tu me décris ton contexte d’usage (HTTP handler, worker, CLI…), il est possible de proposer le pattern exact : valeur de retour, code HTTP, et stratégie de retry.

[1](https://gosamples.dev/context-deadline-exceeded/)
[2](https://pkg.go.dev/context)
[3](https://errorsingo.com/context-context-deadline-exceeded/)
[4](https://rapidevelopers.com/ai-build-errors-debug-solutions-library/request-timed-out)
[5](https://github.com/openai/openai-go)
[6](https://uptrace.dev/glossary/context-deadline-exceeded)
[7](https://stackoverflow.com/questions/74718429/context-deadline-exceeded-for-go-routine)
[8](https://community.temporal.io/t/how-to-best-handle-mysterious-context-deadline-exceeded-502-errors/2689)
[9](https://stackoverflow.com/questions/72569384/correctly-handle-context-deadlineexceedederror-in-golang)
[10](https://www.aifreeapi.com/en/posts/gpt-5-2-reasoning-timeout-error-fix)
[11](https://community.openai.com/t/managing-timeout-when-waiting-for-the-response-from-chat-completions-request/196633)
[12](https://community.openai.com/t/openai-chat-timeout-how-to-solve/422215)
[13](https://community.openai.com/t/configuring-timeout-for-chatcompletion-python/107226)
[14](https://www.reddit.com/r/golang/comments/1anjeyk/alerting_on_context_timeout/)
[15](https://community.openai.com/t/frequent-api-timeout-errors-recently/106903)
[16](https://community.openai.com/t/timeout-for-openai-chat-completion-in-python/411252)
[17](https://github.com/go-sql-driver/mysql/issues/1631)
[18](https://community.coda.io/t/fix-the-timeout-for-long-openai-api-requests/41559)
[19](https://community.openai.com/t/frequently-getting-api-timeout-error-what-am-i-doing-wrong/611941)
[20](https://community.auth0.com/t/false-context-deadline-exceeded-errors-in-go-management-api/151896)