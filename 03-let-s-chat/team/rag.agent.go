package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/rag/chunks"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func GetRagAgent(ctx context.Context, engineURL string, dataMarkdownPath string) (*rag.Agent, error) {

	ragAgent, err := rag.NewAgent(
		ctx,
		agents.Config{
			Name:      "rag-agent",
			EngineURL: engineURL,
		},
		models.Config{
			Name: env.GetEnvOrDefault("RAG_MODEL", "huggingface.co/unsloth/embeddinggemma-300m-gguf:bf16"),
		},
	)
	if err != nil {
		return nil, err
	}

	// Add data to the RAG agent
	contents, err := files.GetContentFiles(dataMarkdownPath, ".md")
	
	if err != nil {
		return nil, err
	}
	for idx, content := range contents {
		piecesOfDoc := chunks.SplitMarkdownBySections(content)

		for chunkIdx, piece := range piecesOfDoc {

			display.Colorf(display.ColorYellow, "generating vectors... (docs %d/%d) [chunks: %d/%d]\n", idx+1, len(contents), chunkIdx+1, len(piecesOfDoc))

			err := ragAgent.SaveEmbedding(piece)
			if err != nil {
				display.Errorf("failed to save embedding for document %d: %v\n", idx, err)

			}
		}
	}

	return ragAgent, nil

}
