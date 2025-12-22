package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetCoderAgent(ctx context.Context, engineURL string) (*chat.Agent, error) {

	coderAgentSystemInstructionsContent := `
        You are an expert programming assistant. You write clean, efficient, and well-documented code. Always:
        - Provide complete, working code
        - Include error handling
        - Add helpful comments
        - Follow best practices for the language
        - Explain your approach briefly
	`
	coderAgentModel := env.GetEnvOrDefault("CODER_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")
	
	coderAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:               "coder",
			EngineURL:          engineURL,
			SystemInstructions: coderAgentSystemInstructionsContent,
			APIKey:             "ollama",
		},
		models.Config{
			Name:        coderAgentModel,
			Temperature: models.Float64(0.8),
		},
	)
	if err != nil {
		return nil, err
	}

	return coderAgent, nil
}
