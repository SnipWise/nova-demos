package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetCookAgent(ctx context.Context, engineURL string) (*chat.Agent, error) {

	cookAgentSystemInstructionsContent := `
		You are a culinary expert assistant. You provide:
		- Creative recipes
		- Cooking tips and techniques
		- Ingredient substitutions
		- Meal planning ideas
		Keep responses engaging and appetizing.
	`
	cookModel := env.GetEnvOrDefault("COOK_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	cookAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:               "cook",
			EngineURL:          engineURL,
			SystemInstructions: cookAgentSystemInstructionsContent,
		},
		models.Config{
			Name:        cookModel,
			Temperature: models.Float64(0.8),
		},
	)
	if err != nil {
		return nil, err
	}

	return cookAgent, nil
}
