package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func main() {
	ctx := context.Background()

	engineURL := "http://localhost:12434/engines/llama.cpp/v1"
	ragEmbeddingModelId := "ai/embeddinggemma:latest"
	metadataModelId := "hf.co/menlo/jan-nano-gguf:q4_k_m"

	// === CREATE METADATA EXTRACTOR AGENT ===
	metadataExtractorAgent, err := getMetadataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		display.Errorf("❌ Error creating metadata extractor agent: %v", err)
		return
	}

	// === CREATE/LOAD RAG AGENT ===
	ragAgent, err := getRagAgent(ctx, engineURL, ragEmbeddingModelId, metadataExtractorAgent)
	if err != nil {
		display.Errorf("❌ Error creating/loading RAG agent: %v", err)
		return
	}

	_ = ragAgent

	// Now you can use ragAgent to handle queries with retrieval-augmented generation
	display.Infof("✅ RAG agent is ready to use.")

}
