package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/agents/crew"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/prompt"
)

func main() {

	ctx := context.Background()

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:11434/v1")

	ragEmbeddingModelId := env.GetEnvOrDefault("RAG_MODEL", "ai/embeddinggemma:latest")

	metadataModelId := env.GetEnvOrDefault("METADATA_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")
	
	compressorModelID :=env.GetEnvOrDefault("COMPRESSOR_MODEL", "ai/qwen2.5:3B-F16")

	chatModelID := env.GetEnvOrDefault("CHAT_MODEL", "hf.co/menlo/jan-nano-128k-gguf:q4_k_m")

	//chatModelID := "hf.co/menlo/jan-nano-128k-gguf:q4_k_m"
	//chatModelID := "hf.co/menlo/jan-nano-gguf:q4_k_m"

	// === CREATE CHAT AGENT ===
	chatAgent, err := GetChatAgent(ctx, engineURL, chatModelID)
	if err != nil {
		display.Errorf("❌ Error creating chat agent: %v", err)
		return
	}

	// === CREATE METADATA EXTRACTOR AGENT ===
	metadataExtractorAgent, err := GetMetaDataExtractorAgent(ctx, engineURL, metadataModelId)
	if err != nil {
		display.Errorf("❌ Error creating metadata extractor agent: %v", err)
		return
	}

	// === CREATE/LOAD RAG AGENT ===
	ragAgent, err := GetRagAgent(ctx, engineURL, ragEmbeddingModelId)
	if err != nil {
		display.Errorf("❌ Error creating/loading RAG agent: %v", err)
		return
	}

	// === CREATE COMPRESSOR AGENT ===
	compressorAgent, err := GetCompressorAgent(ctx, engineURL, compressorModelID)
	if err != nil {
		display.Errorf("❌ Error creating compressor agent: %v", err)
		return
	}

	// === LOAD DATA INTO RAG STORE ===
	//dataPath := "./docs-for-test"
	dataPath := env.GetEnvOrDefault("DATA_PATH", "./docs")
	//storePathFile := "./store-for-test/support.json"
	storePathFile := env.GetEnvOrDefault("STORE_PATH_FILE", "./store/support.json")

	err = LoadData(dataPath, storePathFile, ragAgent, metadataExtractorAgent)
	if err != nil {
		display.Errorf("❌ Error loading data into RAG store: %v", err)
		return
	}

	//display.Infof("✅ Data loaded into RAG store successfully.")

	crewAgent, err := crew.NewSimpleAgent(
		ctx,
		map[string]*chat.Agent{
			"chat": chatAgent,
		},
		"chat",
	)
	if err != nil {
		display.Errorf("❌ Error creating crew agent: %v", err)
		return
	}
	// Attach the RAG agent to the server agent
	crewAgent.SetRagAgent(ragAgent)

	// Attach the compressor agent to the server agent
	crewAgent.SetCompressorAgent(compressorAgent)

	crewAgent.SetContextSizeLimit(8500)

	crewAgent.SetSimilarityLimit(0.4)
	crewAgent.SetMaxSimilarities(10)

	for {

		markdownParser := display.NewMarkdownChunkParser()

		input := prompt.NewWithColor("🤖 Ask me something? [" + crewAgent.GetName() + "]")
		question, err := input.RunWithEdit()

		if err != nil {
			display.Errorf("failed to get input: %v", err)
			return
		}
		if strings.HasPrefix(question, "/bye") {
			display.Infof("👋 Goodbye!")
			break
		}

		if strings.HasPrefix(question, "/messages") {
			display.Infof("💬 Current conversation messages:")
			for i, msg := range crewAgent.GetMessages() {
				display.Infof("Message %d - Role: %s, Content: \n%s", i, msg.Role, msg.Content)
				display.Separator()
			}
			continue
		}

		if strings.HasPrefix(question, "/reset") {
			display.Infof("🔄 Resetting %s context", crewAgent.GetName())
			crewAgent.ResetMessages()
			continue
		}

		display.NewLine()

		result, err := crewAgent.StreamCompletion(question, func(chunk string, finishReason string) error {

			// Use markdown chunk parser for colorized streaming output
			if chunk != "" {
				display.MarkdownChunk(markdownParser, chunk)
			}
			if finishReason == "stop" {
				markdownParser.Flush()
				markdownParser.Reset()
				//markdownParser.Flush()
				display.NewLine()
			}
			return nil
		})

		if err != nil {
			display.Errorf("[%s][%v]failed to get completion: %v", crewAgent.GetName(), crewAgent.GetContextSize(), err)
			return
		}

		display.NewLine()
		display.Separator()
		display.KeyValue("Finish reason", result.FinishReason)
		display.KeyValue("Context size", fmt.Sprintf("%d characters", crewAgent.GetContextSize()))
		display.Separator()
	}

}

// go test -v -run TestChatSheetFact
