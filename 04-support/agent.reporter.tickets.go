package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
)

func GetTicketsReporterAgent(ctx context.Context, engineURL, reporterModelId string) (*chat.Agent, error) {
	// Ticket Reporter Agent
	ticketsReportInstructions, err := files.ReadTextFile("./ticket.report.instructions.md")

	if err != nil {
		panic(err)
	}
	ticketsReporterAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:                    "ticket-reporter-agent",
			EngineURL:               engineURL,
			SystemInstructions:      ticketsReportInstructions,
			KeepConversationHistory: false,
		},
		models.Config{
			Name:        reporterModelId,
			Temperature: models.Float64(0.8),
		},
	)
	if err != nil {
		return nil, err
	}
	return ticketsReporterAgent, nil
}
