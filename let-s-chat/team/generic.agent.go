package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetGenericAgent(ctx context.Context, engineURL string) (*chat.Agent, error) {

	genericAgentSystemInstructionsContent := `
        You respond appropriately to different types of questions.
        For factual questions: Give direct answers with key facts
        For how-to questions: Provide step-by-step guidance
        For opinion questions: Present balanced perspectives
        For complex topics: Break into digestible parts

        Always start with the most important information.	
	`
	genericModel := env.GetEnvOrDefault("GENERIC_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	genericAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:               "generic",
			EngineURL:          engineURL,
			SystemInstructions: genericAgentSystemInstructionsContent,
		},
		models.Config{
			Name:        genericModel,
			Temperature: models.Float64(0.8),
		},
	)
	if err != nil {
		return nil, err
	}

	return genericAgent, nil
}
