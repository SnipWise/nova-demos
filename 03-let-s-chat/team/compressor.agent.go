package team

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/compressor"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
)

func GetCompressorAgent(ctx context.Context, engineURL string) (*compressor.Agent, error) {
	compressorAgent, err := compressor.NewAgent(
		ctx,
		agents.Config{
			Name:               "compressor-agent",
			EngineURL:          engineURL,
			SystemInstructions: compressor.Instructions.Minimalist,
		},
		models.Config{
			Name:        env.GetEnvOrDefault("COMPRESSOR_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m"),
			Temperature: models.Float64(0.0),
		},
		compressor.WithCompressionPrompt(compressor.Prompts.Minimalist),
	)
	if err != nil {
		return nil, err
	}
	return compressorAgent, nil
}
