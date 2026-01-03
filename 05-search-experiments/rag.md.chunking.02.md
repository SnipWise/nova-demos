En Go, il n’existe pas encore de « MarkdownTextSplitter » canonique, donc la voie la plus réaliste est un splitter custom Markdown-aware, éventuellement combiné avec une lib de tokenisation.

## Approche générale en Go

- Parser le Markdown ligne par ligne, détecter les **headers** (`^#{1,6} `) hors blocs de code, et constituer un `[]Chunk{}` avec `Title`, `Depth`, `Content`.[1][2]
- Ne jamais couper à l’intérieur d’un bloc de code ```…```
- Ensuite, fusionner les petites sections adjacentes tant qu’on reste sous un budget de tokens (e.g. 400–800), en utilisant un compteur de tokens (tiktoken-like) ou à défaut un proxy caractères.[3][1]

## Pseudo-code Go minimal

- Représentation d’un chunk :  

  ```go
  type Chunk struct {
      Title   string
      Depth   int
      Content string
  }
  ```

- Split par headers Markdown en respectant les blocs de code :  

  ```go
  func SplitMarkdownByHeaders(md string) []Chunk {
      var chunks []Chunk
      lines := strings.Split(md, "\n")

      var curTitle string
      var curDepth int
      var buf []string
      inCode := false

      flush := func() {
          txt := strings.TrimSpace(strings.Join(buf, "\n"))
          if txt != "" {
              chunks = append(chunks, Chunk{
                  Title:   curTitle,
                  Depth:   curDepth,
                  Content: txt,
              })
          }
          buf = nil
      }

      for _, line := range lines {
          // toggle code block
          if strings.HasPrefix(line, "```
              inCode = !inCode
              buf = append(buf, line)
              continue
          }

          if !inCode {
              if m := headerRegex.FindStringSubmatch(line); m != nil {
                  // nouveau header → flush précédent
                  flush()
                  curDepth = len(m)        // nombre de #[4]
                  curTitle = strings.TrimSpace(m)[5]
                  buf = append(buf, line)
                  continue
              }
          }

          buf = append(buf, line)
      }
      flush()
      return chunks
  }
  ```

  Avec par ex. :  

  ```
  var headerRegex = regexp.MustCompile(`^(#{1,6})\s+(.*)$`)
  ```

Cette logique est très proche de ce qui est proposé côté Mastra pour TS (split, puis merge par profondeur en respectant un seuil de longueur). [web:25][web:32]  

## Re-chunking par taille (tokens ou caractères)

- Une fois `[]Chunk` obtenu, tu peux :  
  - Calculer une « longueur » (tokens ou caractères) pour chaque chunk. [web:6][web:30]  
  - Parcourir du plus profond au moins profond pour fusionner des sections voisines tant que `len(prev)+len(cur) < maxTokens` (pattern Mastra `joinThreshold`). [web:25][web:32]  
- Si tu utilises un modèle OpenAI / Llama, tu peux brancher un tokenizer Go (par ex. wrappers tiktoken ou llama.cpp) pour calculer exactement la longueur en tokens avant merge. [web:6][web:30]  

## Métadonnées RAG côté Go

Pour chaque chunk, stocker en plus dans ta struct :

```
type Chunk struct {
    Title       string
    Depth       int
    Content     string
    Path        string            // chemin de fichier
    HeadersPath []string          // [H1, H2, H3...]
}
```

- Construire `HeadersPath` en gardant en mémoire la dernière valeur pour chaque niveau de profondeur lors du parcours. [web:24][web:30]  
- Indexer `Content` mais aussi `strings.Join(HeadersPath, " > ")` comme contexte pour améliorer les réponses. [web:24][web:33]  

## Si tu veux un snippet complet

Si tu me donnes :  
- ton backend d’embed (OpenAI, Ollama, vLLM…)  
- ton vecteur store (PGVector, Weaviate, Milvus, Qdrant…)  

une version plus complète en Go (avec token counting approximatif ou réel + struct de metadata + fonction `ChunkMarkdown(filePath string) ([]Chunk, error)`) peut être esquissée quasi prête à coller dans ton projet.

[1](https://weaviate.io/blog/chunking-strategies-for-rag)
[2](https://github.com/mastra-ai/mastra/issues/2755)
[3](https://www.pinecone.io/learn/chunking-strategies/)
[4](https://www.reddit.com/r/Rag/comments/1mwf71t/struggling_with_rag_performance_and_chunking/)
[5](https://www.ayinedjimi-consultants.fr/ia-erreurs-communes-chunking.html)
[6](https://www.reddit.com/r/Rag/comments/1lcqw1x/embeddingschunking_for_markdown_content/)
[7](https://dev.to/rajinh24/rag-chunking-strategies-that-actually-work-and-why-most-dont-49n2)
[8](https://github.com/messkan/rag-chunk)
[9](https://mastra.ai/examples/rag/chunking/chunk-markdown)
[10](https://lagnchain.readthedocs.io/en/latest/modules/indexes/text_splitters/examples/markdown.html)
[11](https://dev.to/oleh-halytskyi/optimizing-rag-context-chunking-and-summarization-for-technical-docs-3pel)
[12](https://langchain-cn.readthedocs.io/en/latest/modules/indexes/text_splitters/examples/markdown.html)
[13](https://mastra.ai/reference/rag/chunk)
[14](https://github.com/langchain-ai/langchain/issues/2836)
[15](https://mastra.ai/docs/rag/chunking-and-embedding)
[16](https://learn.microsoft.com/en-us/azure/architecture/ai-ml/guide/rag/rag-chunking-phase)
[17](https://github.com/langchain-ai/langchain/blob/master/libs/text-splitters/langchain_text_splitters/markdown.py)
[18](https://github.com/mastra-ai/mastra/blob/main/docs/src/content/en/docs/rag/chunking-and-embedding.mdx)
[19](https://ai.gopubby.com/the-definitive-guide-to-chunking-strategies-for-llms-and-rag-57e20b9d855d)
[20](https://api.python.langchain.com/en/latest/markdown/langchain_text_splitters.markdown.MarkdownTextSplitter.html)
[21](https://mastra.ai/examples/rag/chunking/chunk-text)
[22](https://docs.langchain.com/oss/python/integrations/splitters/markdown_header_metadata_splitter)