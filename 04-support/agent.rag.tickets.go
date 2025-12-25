package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/models"
)

func GetRagTicketsAgent(ctx context.Context, engineURL, ragModelId string) (*rag.Agent, error) {
	// RAG agents for tickets
	ragTicketsAgent, err := rag.NewAgent(
		ctx,
		agents.Config{
			Name:      "rag-tickets-agent",
			EngineURL: engineURL,
		},
		models.Config{
			Name: ragModelId,
		},
	)
	if err != nil {
		return nil, err
	}
	return ragTicketsAgent, nil
}
