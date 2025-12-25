package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetThinkerAgent(ctx context.Context, engineURL string) (*chat.Agent, error) {

	thinkerAgentSystemInstructionsContent := `
        You are a thoughtful conversational assistant. 
        - Listen carefully to the user
        - Think before responding
        - Ask clarifying questions when needed
        - Discuss topics with curiosity and respect
        - Admit when you don't know something
        Keep responses natural and conversational.	
	`
	thinkerModel := env.GetEnvOrDefault("GENERIC_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")

	thinkerAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:               "thinker",
			EngineURL:          engineURL,
			SystemInstructions: thinkerAgentSystemInstructionsContent,
		},
		models.Config{
			Name:        thinkerModel,
			Temperature: models.Float64(0.8),
		},
	)
	if err != nil {
		return nil, err
	}

	return thinkerAgent, nil
}
