package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/compressor"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func GetCompressorAgent(ctx context.Context, engineURL, compressorModelID string) (*compressor.Agent, error) {

	compressorAgent, err := compressor.NewAgent(
		ctx,
		agents.Config{
			Name:               "compressor-agent",
			EngineURL:          engineURL,
			SystemInstructions: compressor.Instructions.Effective,
		},
		models.Config{
			Name:        compressorModelID,
			Temperature: models.Float64(0.0),
		},
		compressor.WithCompressionPrompt(compressor.Prompts.UltraShort),
	)

	if err != nil {
		display.Errorf("❌ Error creating compressor agent: %v", err)
		return nil, err
	}
	display.Infof("✅ Compressor agent created for context packing")

	return compressorAgent, nil

}
