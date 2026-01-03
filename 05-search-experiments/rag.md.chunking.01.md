Pour des fichiers Markdown, les stratégies de chunking les plus efficaces sont celles qui tirent parti de la structure (titres, listes, blocs de code) plutôt que de découper « au kilomètre ».[1][2]

## Principes spécifiques au Markdown

- Utiliser les **titres** (`#`, `##`, `###`) comme frontières de chunks produit des unités sémantiques cohérentes et améliore nettement la pertinence vs du fixe 500 tokens.[2][1]
- Préserver les blocs de **code** et de listes dans un même chunk évite de casser le contexte technique (signature de fonction séparée de l’explication, etc.).[3][4]
- Ajouter en metadata la hiérarchie des titres (h1 > h2 > h3) rend la navigation et le re-ranking plus fiables, surtout pour la doc technique.[4][5]

## Stratégies de chunking pertinentes

- Chunking « Markdown-aware » : splitter sur les titres, blocs horizontaux et sections structurées, souvent 200–1000 tokens par chunk selon la doc.[6][3][2]
- Chunking sémantique : regrouper des phrases avec embeddings et couper là où la similarité chute (changement de sujet), utile pour sections longues peu structurées.[7][8][1]
- Hybrid / recursive : d’abord découpe par sections Markdown, puis sous‑chunking à taille raisonnable via un splitter récursif (paragraphes → phrases → mots).[9][5][10]

## Outils / libs déjà prêts

- LangChain `MarkdownTextSplitter` / `MarkdownHeaderTextSplitter` : découpe par titres Markdown, combinable avec `RecursiveCharacterTextSplitter` pour affiner.[5][3][9]
- Mastra `MDocument.fromMarkdown(...).chunk()` : gère titres, listes, blocs de code de façon spécifique au Markdown pour RAG.[11]
- Azure AI Search / autres vendors : proposent du « semantic chunking » avec output Markdown, utilisable directement pour l’indexation.[12][7]

## Bonnes pratiques RAG sur Markdown

- Viser des chunks assez petits pour la **précision** (souvent 300–800 tokens), mais enrichis avec le contexte (titre du doc, section, sous‑section).[13][14][6]
- Ajouter un header contextuel dans le contenu du chunk (titre du doc, date, URL, chemin de fichier) améliore les réponses sans surcharger la fenêtre.[13][2]
- Sur de la doc technique : éviter tout ce qui modifie la mise en forme des blocs de code (splitter custom ou lib qui respecte les fences ```). [web:14][web:18]  

## Si tu veux un pattern concret (type code)

Pour un pipeline « propre » Markdown de doc technique :

- Étape 1 : parse Markdown → segments par `#`, `##`, `###`. [web:8][web:14]  
- Étape 2 : pour chaque segment, appliquer un splitter récursif (par paragraphe puis phrase) en visant ~400–600 tokens avec overlap 10–15 %. [web:12][web:4]  
- Étape 3 : attacher en metadata la hiérarchie des titres + chemin de fichier + type de bloc (texte, code, liste) et éventuellement injecter le chemin / titres en préambule du chunk. [web:14][web:3]  

Si tu me donnes ton stack (TS/Go/Python, LangChain/Mastra/autre) et un exemple de Markdown (doc API, notes, wiki interne…), une proposition de splitter et de schéma de metadata très concret peut être ajustée.

[1](https://www.pinecone.io/learn/chunking-strategies/)
[2](https://www.snowflake.com/en/engineering-blog/impact-retrieval-chunking-finance-rag/)
[3](https://lagnchain.readthedocs.io/en/latest/modules/indexes/text_splitters/examples/markdown.html)
[4](https://dev.to/oleh-halytskyi/optimizing-rag-context-chunking-and-summarization-for-technical-docs-3pel)
[5](https://www.kaggle.com/code/ksmooi/langchain-mastering-text-splitting)
[6](https://www.axopen.com/blog/2025/08/comment-fonctionne-un-rag/)
[7](https://milvus.io/ai-quick-reference/what-chunking-strategies-work-best-for-document-indexing)
[8](https://developer.ibm.com/articles/awb-enhancing-rag-performance-chunking-strategies/)
[9](https://langchain-opentutorial.gitbook.io/langchain-opentutorial/07-textsplitter/02-recursivecharactertextsplitter)
[10](https://docs.thub.tech/langchain/text-splitters)
[11](https://mastra.ai/examples/rag/chunking/chunk-markdown)
[12](https://learn.microsoft.com/en-us/azure/search/vector-search-how-to-chunk-documents)
[13](https://numerique.gouv.fr/sinformer/blog/am%C3%A9liorer-les-performances-des-rags--aper%C3%A7u-des-techniques-de-chunking-augment%C3%A9/)
[14](https://datacorner.fr/document-chunking/)
[15](https://www.reddit.com/r/Rag/comments/1mwf71t/struggling_with_rag_performance_and_chunking/)
[16](https://www.ayinedjimi-consultants.fr/ia-erreurs-communes-chunking.html)
[17](https://www.followtribes.io/comprendre-structure-rag-pipeline-ia/)
[18](https://hnlab.huma-num.fr/blog/2025/08/26/explorer-ses-documents-avec-la-RAG/)
[19](https://docling-project.github.io/docling/concepts/chunking/)
[20](https://meritis.fr/rag-avance-techniques-et-optimisations-pour-ameliorer-les-performances-partie-2-4/)