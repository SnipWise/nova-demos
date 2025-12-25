package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/models"
)

func GetRagReportsAgent(ctx context.Context, engineURL, ragModelId string) (*rag.Agent, error) {
	// RAG agent for reports
	ragReportsAgent, err := rag.NewAgent(
		ctx,
		agents.Config{
			Name:      "rag-reports-agent",
			EngineURL: engineURL,
		},
		models.Config{
			Name: ragModelId,
		},
	)
	if err != nil {
		return nil, err
	}
	return ragReportsAgent, nil
}
