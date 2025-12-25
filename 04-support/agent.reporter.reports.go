package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
)

func GetSolutionReportsReporterAgent(ctx context.Context, engineURL, reporterModelId string) (*chat.Agent, error) {
	// Solution Reporter Agent
	solutionsReportInstructions, err := files.ReadTextFile("./solution.report.instructions.md")

	if err != nil {
		panic(err)
	}
	solutionReportsReporterAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:                    "solution-reporter-agent",
			EngineURL:               engineURL,
			SystemInstructions:      solutionsReportInstructions,
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
	return solutionReportsReporterAgent, nil
}
