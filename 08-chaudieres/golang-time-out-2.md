L’erreur que tu vois ne vient pas d’OpenAI mais du runner de tests Go : `go test` panique après 10 minutes parce que ton binaire de test (et donc `TestGenerateAllMetaData`) n’a pas fini dans le délai par défaut.[1][2]

## Ce que signifie ce stack trace

- `panic: test timed out after 10m0s` est généré par `testing.(*M).startAlarm.func1`, c’est le mécanisme interne de `go test -timeout`.[3]
- Ton goroutine 5 montre `net/http.(*Client).Do` et `github.com/openai/openai-go/v3/internal/requestconfig.(*RequestConfig).Execute`, donc la requête OpenAI est toujours en cours quand le timeout global des tests est atteint.[4][5]

## Options côté `go test`

Deux approches possibles :

1. **Augmenter le timeout global de la suite de tests** (pratique pour des tests d’intégration lents) :
   ```bash
   go test ./... -timeout 30m
   # ou pour un seul package
   go test -timeout 30m ./testing/08-chaudieres
   ```
   Le flag `-timeout` contrôle la durée maximale du binaire de tests ; `0` désactive le timeout.[2][1]

2. **Désactiver le timeout pour ces tests précis** (à utiliser avec prudence) :
   ```bash
   go test -timeout 0 ./testing/08-chaudieres
   ```
   Avec `0`, le binaire n’est jamais tué par `go test`, donc si la requête OpenAI se bloque, ton test restera bloqué.[1]

## Ce qu’il faut ajuster dans ton code de test

En parallèle, il reste utile de garder un timeout applicatif sur l’appel OpenAI, pour ne pas laisser un test coincé pendant des heures :

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
defer cancel()

resp, err := client.Chat.Completions.New(ctx, params /* ... */)
if err != nil {
    if errors.Is(err, context.DeadlineExceeded) {
        t.Fatalf("timeout OpenAI après 2min: %v", err)
    }
    t.Fatalf("erreur OpenAI: %v", err)
}
_ = resp
```

- Le timeout **du contexte** borne la durée d’un appel individuel à OpenAI.[5]
- Le timeout **`go test -timeout`** borne la durée totale du binaire de tests, tous tests/goroutines confondus.[2]

Pour ton `TestGenerateAllMetaData`, la combinaison recommandée est : `context.WithTimeout` raisonnable dans le test, plus `go test -timeout` augmenté (ex. 20–30 min) si ce test doit traiter beaucoup de données.

[1](https://stackoverflow.com/questions/77095789/how-to-avoid-timeout-in-go-test-after-10m)
[2](https://forum.golangbridge.org/t/what-does-the-timeout-flag-in-go-test-represent/17505)
[3](https://stackoverflow.com/questions/39449079/recover-from-test-timeout-panic-in-golang)
[4](https://terratest.gruntwork.io/docs/testing-best-practices/timeouts-and-logging/)
[5](https://github.com/openai/openai-go)
[6](https://github.com/golang/go/issues/61468)
[7](https://github.com/golang/go/issues/58725)
[8](https://stackoverflow.com/questions/75787638/openai-gpt-3-api-error-request-timed-out)
[9](https://github.com/openai/openai-python/issues/1596)
[10](https://github.com/golang/go/issues/48157)
[11](https://github.com/sashabaranov/go-openai/blob/master/chat_test.go)
[12](https://github.com/openai/openai-python/issues/1134)
[13](https://community.openai.com/t/openai-chat-timeout-how-to-solve/422215)
[14](https://community.openai.com/t/is-it-still-supported-to-set-a-request-timeout-for-request/1078461)
[15](https://blog.csdn.net/qq_33690342/article/details/105434198)
[16](https://pkg.go.dev/github.com/openai/openai-go/option)
[17](https://github.com/golang/go/issues/57305)
[18](https://github.com/openai/openai-go/blob/main/CHANGELOG.md)
[19](https://github.com/openai/openai-python/issues/2688)
[20](https://groups.google.com/g/golang-nuts/c/L2k5LAg85ck)