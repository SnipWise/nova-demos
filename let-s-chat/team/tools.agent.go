package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/tools"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetToolsAgent(ctx context.Context, engineURL string, toolList []*tools.Tool) (*tools.Agent, error) {

	toolsAgent, err := tools.NewAgent(
		ctx,
		agents.Config{
			Name:      "tools-agent",
			EngineURL: engineURL,
			SystemInstructions: `
			You are an AI assistant that can call tools to help answer user queries effectively.
			- Always decide when to use a tool based on the user's request.
			- Choose the most appropriate tool for the task.
			- Provide clear and concise arguments to the tool.
			- After calling a tool, use the result to formulate your final response to the user.
			`,
		},
		models.Config{
			Name:              env.GetEnvOrDefault("GENERIC_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m"),
			Temperature:       models.Float64(0.0),
			ParallelToolCalls: models.Bool(true),
		},
		tools.WithTools(toolList),
	)
	if err != nil {
		return nil, err
	}
	return toolsAgent, nil
}
