package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func GetChatAgent(ctx context.Context, engineURL, chatModelID string) (*chat.Agent, error) {
	configPath := env.GetEnvOrDefault("CONFIG_PATH", "./config")
	chatInstructions, err := files.ReadTextFile(configPath+"/chat.instructions.md")
	if err != nil {
		panic(err)
	}

	chatAgent, err := chat.NewAgent(
		ctx,
		agents.Config{
			Name:                    "chat-agent",
			EngineURL:               engineURL,
			SystemInstructions:      chatInstructions,
			KeepConversationHistory: false, // IMPORTANT: we don't need to keep history for this demo
		},
		models.Config{
			Name:        chatModelID,
			Temperature: models.Float64(0.0),
		},
	)

	if err != nil {
		display.Errorf("❌ Error creating compressor agent: %v", err)
		return nil, err
	}
	display.Infof("✅ Chat agent created")

	return chatAgent, nil

}
