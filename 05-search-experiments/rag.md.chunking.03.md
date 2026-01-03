Voici un snippet Go complet et autonome pour chunker du Markdown en respectant les headers, les blocs de code et une taille max approximative (en caractères) par chunk.[1][2]

## Snippet complet

```go
package markdownchunk

import (
	"os"
	"regexp"
	"strings"
)

type Chunk struct {
	Title       string   // dernier header rencontré
	Depth       int      // niveau de header (#=1, ##=2, etc.)
	Content     string   // texte du chunk
	Path        string   // chemin du fichier source
	HeadersPath []string // chemin hiérarchique H1 > H2 > H3...
}

// options simples de chunking
type Options struct {
	MaxChars   int // taille max d'un chunk (approx. proxy de tokens)
	MinChars   int // taille min pour éviter des micro-chunks
	Overlap    int // overlap en caractères entre chunks (sur le texte brut)
	WithPath   string
}

var headerRegex = regexp.MustCompile(`^(#{1,6})\s+(.*)$`)

// ChunkMarkdownFile lit un fichier et renvoie les chunks.
func ChunkMarkdownFile(path string, opts Options) ([]Chunk, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	opts.WithPath = path
	return ChunkMarkdown(string(b), opts), nil
}

// ChunkMarkdown prend du markdown et renvoie une liste de chunks prêts à être indexés.
func ChunkMarkdown(md string, opts Options) []Chunk {
	if opts.MaxChars <= 0 {
		opts.MaxChars = 4000
	}
	if opts.MinChars <= 0 {
		opts.MinChars = opts.MaxChars / 4
	}
	if opts.Overlap < 0 {
		opts.Overlap = 0
	}

	sections := splitByHeaders(md)
	sections = attachHeaderPaths(sections)

	var out []Chunk
	for _, sec := range sections {
		// re-chunking par taille max
		sub := splitSectionBySize(sec, opts)
		for _, c := range sub {
			c.Path = opts.WithPath
			out = append(out, c)
		}
	}
	return out
}

// Section interne, avant re-split.
type section struct {
	Title       string
	Depth       int
	Content     string
	HeadersPath []string
}

func splitByHeaders(md string) []section {
	lines := strings.Split(md, "\n")
	var sections []section

	var curTitle string
	var curDepth int
	var buf []string
	inCode := false

	flush := func() {
		txt := strings.TrimSpace(strings.Join(buf, "\n"))
		if txt == "" {
			buf = nil
			return
		}
		sections = append(sections, section{
			Title:   curTitle,
			Depth:   curDepth,
			Content: txt,
		})
		buf = nil
	}

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		// toggle bloc de code
		if strings.HasPrefix(trim, "```
			inCode = !inCode
			buf = append(buf, line)
			continue
		}

		// détecter un header uniquement hors code block
		if !inCode {
			if m := headerRegex.FindStringSubmatch(line); m != nil {
				// nouveau header -> flush section précédente
				flush()
				curDepth = len(m)[3]
				curTitle = strings.TrimSpace(m)[4]
				buf = append(buf, line)
				continue
			}
		}

		buf = append(buf, line)
	}
	flush()
	return sections
}

// construit HeadersPath (H1 > H2 > H3...) pour chaque section.
func attachHeaderPaths(sections []section) []section {
	var lastByDepth string // index 0 inutilisé[5]
	for i, s := range sections {
		if s.Depth >= 1 && s.Depth <= 6 {
			lastByDepth[s.Depth-1] = s.Title
			// reset plus profonds
			for d := s.Depth; d < 6; d++ {
				lastByDepth[d] = ""
			}
		}
		var path []string
		for d := 0; d < 6; d++ {
			if lastByDepth[d] != "" {
				path = append(path, lastByDepth[d])
			}
		}
		sections[i].HeadersPath = path
	}
	return sections
}

// découpe une section en sous-chunks par taille max de caractères, avec overlap.
func splitSectionBySize(sec section, opts Options) []Chunk {
	content := strings.TrimSpace(sec.Content)
	if content == "" {
		return nil
	}
	runes := []rune(content)
	n := len(runes)

	if n <= opts.MaxChars {
		return []Chunk{{
			Title:       sec.Title,
			Depth:       sec.Depth,
			Content:     content,
			HeadersPath: sec.HeadersPath,
		}}
	}

	var chunks []Chunk
	start := 0
	for start < n {
		end := start + opts.MaxChars
		if end > n {
			end = n
		}

		// essayer de couper sur une limite de paragraphe/ligne pour garder le sens
		sliceEnd := findGoodBreak(runes, start, end)

		chunkText := strings.TrimSpace(string(runes[start:sliceEnd]))
		if len(chunkText) > 0 {
			chunks = append(chunks, Chunk{
				Title:       sec.Title,
				Depth:       sec.Depth,
				Content:     chunkText,
				HeadersPath: sec.HeadersPath,
			})
		}

		if sliceEnd >= n {
			break
		}

		// Overlap
		start = sliceEnd - opts.Overlap
		if start < 0 {
			start = 0
		}
	}
	return chunks
}

// cherche un bon point de coupure autour de end (priorité double saut de ligne, puis simple).
func findGoodBreak(runes []rune, start, end int) int {
	n := len(runes)
	if end >= n {
		return n
	}

	// chercher en arrière un "\n\n"
	for i := end - 1; i > start+50 && i >= 1; i-- {
		if runes[i] == '\n' && runes[i-1] == '\n' {
			return i
		}
	}
	// sinon chercher un "\n"
	for i := end - 1; i > start+50; i-- {
		if runes[i] == '\n' {
			return i
		}
	}
	// sinon on coupe brutalement
	return end
}
```

## Utilisation simple

```
chunks, err := ChunkMarkdownFile("docs/guide.md", Options{
	MaxChars: 4000,
	MinChars: 500,
	Overlap:  200,
})
if err != nil {
	// handle error
}
for _, c := range chunks {
	// indexer c.Content + metadata (c.Title, c.HeadersPath, c.Path)
}
```

Ce pattern reproduit l’idée « split par headers Markdown, respecter les blocs de code, puis re-chunker par taille avec overlap », recommandée pour le RAG sur docs techniques. [web:24][web:29][web:32]

[1](https://weaviate.io/blog/chunking-strategies-for-rag)
[2](https://github.com/mastra-ai/mastra/issues/2755)
[3](https://www.reddit.com/r/Rag/comments/1mwf71t/struggling_with_rag_performance_and_chunking/)
[4](https://www.ayinedjimi-consultants.fr/ia-erreurs-communes-chunking.html)
[5](https://www.pinecone.io/learn/chunking-strategies/)