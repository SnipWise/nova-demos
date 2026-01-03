package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/rag/chunks"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

// getRagAgent creates or loads a RAG agent with JSON store persistence
func getRagAgent(ctx context.Context, engineURL, embeddingModelId string, metadataExtractorAgent *structured.Agent[KeywordMetadata]) (*rag.Agent, error) {

	dataPath := "./data"
	storePathFile := "./store/support.json"

	display.Infof("📦 RAG Store path: %s", storePathFile)

	// === CREATE RAG AGENT ===
	ragAgent, err := rag.NewAgent(
		ctx,
		agents.Config{
			EngineURL: engineURL,
		},
		models.Config{
			Name: embeddingModelId,
		},
	)
	if err != nil {
		display.Errorf("❌ Error creating RAG agent: %v", err)
		return nil, err
	}

	// === LOAD OR CREATE STORE ===
	if ragAgent.StoreFileExists(storePathFile) {
		// Load existing store
		err := ragAgent.LoadStore(storePathFile)
		if err != nil {
			display.Errorf("❌ Error loading store %s: %v", storePathFile, err)
			return nil, err
		}
		display.Infof("✅ RAG store loaded from %s", storePathFile)
	} else {
		display.Infof("📝 Store not found. Creating new store and indexing character sheet...")

		filesToChunk, err := files.GetContentFilesAsMap(dataPath, ".md")

		if err != nil {
			display.Errorf("❌ Error reading content files from %s: %v", dataPath, err)
			return nil, err
		}

		for fileName, fileContent := range filesToChunk {
			display.Infof("📄 Indexing: %s", fileName)
			fmt.Println(strings.Repeat("-", 40))
			//fmt.Println(fileContent)

			// === CHUNK AND INDEX CONTENT ===
			// Split markdown by sections
			contentPieces := chunks.ChunkText(fileContent, 1024, 256)
			// TODO: try another chunking strategy that preserves section titles
			display.Infof("📄 Split document into %d sections", len(contentPieces))

			// Index each section
			//for idx, piece := range contentPieces[1:] { // Skip title section
			for idx, piece := range contentPieces { // Skip title section

				// === EXTRACT METADATA FOR SECTION ===
				// Extract keywords and metadata using structured agent
				extractionPrompt := fmt.Sprintf(`Analyze the following content and extract relevant metadata.
					Content:
					%s

					Extract:
					- Keywords: only 4 keywords, important terms and concepts from the markdown section title then from the content
					- Main topic: the primary subject (use the markdown section title)
					- Category: type of content
					`,
					piece,
				)

				metadata, _, err := metadataExtractorAgent.GenerateStructuredData([]messages.Message{
					{Role: roles.User, Content: extractionPrompt},
				})
				if err != nil {
					display.Errorf("❌ Error extracting keywords from section %d: %v", idx, err)
					// Continue with embedding even if keyword extraction fails
				} else {
					display.Infof("🏷️  Keywords: %v", metadata.Keywords)
					display.Infof("📌 Topic: %s | Category: %s",
						metadata.MainTopic, metadata.Category)

					// Enrich the chunk with metadata
					enrichedPiece := fmt.Sprintf("[METADATA]\nKeywords: %v\nTopic: %s\nCategory: %s\n\nContent:\n%s \nFile: %s\n",
						metadata.Keywords, metadata.MainTopic, metadata.Category, piece, fileName,
					)

					piece = enrichedPiece
				}

				fmt.Println(strings.Repeat("-", 20))
				fmt.Println(piece)
				fmt.Println(strings.Repeat("-", 20))

				// === SAVE EMBEDDING FOR SECTION ===
				err = ragAgent.SaveEmbedding(piece)
				if err != nil {
					display.Errorf("❌ Error embedding section %d: %v", idx, err)
				} else {
					display.Infof("✅ Indexed section %d/%d", idx+1, len(contentPieces))
					fmt.Println(piece)
				}

			}

		}

		// === PERSIST STORE TO DISK ===
		err = ragAgent.PersistStore(storePathFile)
		if err != nil {
			display.Errorf("❌ Error persisting store: %v", err)
			return nil, err
		}
		display.Infof("💾 RAG store saved to %s", storePathFile)
	}

	return ragAgent, nil
}
