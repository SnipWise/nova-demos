package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetOrchestratorAgent(ctx context.Context, engineURL string) (*structured.Agent[agents.Intent], error) {
	orchestratorAgentSystemInstructions := `
        You are good at identifying the topic of a conversation.
        Given a user's input, identify the main topic of discussion in only one word.
        The possible topics are: Technology, Health, Sports, Entertainment, Politics, Science, Mathematics,
        Travel, Food, Education, Finance, Environment, Fashion, History, Literature, Art,
        Music, Psychology, Relationships, Philosophy, Religion, Automotive, Gaming, Translation.
        Respond in JSON format with the field 'topic_discussion'.
	`

	orchestratorAgent, err := structured.NewAgent[agents.Intent](
		ctx,
		agents.Config{
			Name:               "orchestrator-agent",
			EngineURL:          engineURL,
			SystemInstructions: orchestratorAgentSystemInstructions,
		},
		models.Config{
			Name:        env.GetEnvOrDefault("ORCHESTRATOR_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m"),
			Temperature: models.Float64(0.0),
		},
	)
	if err != nil {
		return nil, err
	}
	return orchestratorAgent, nil
}
