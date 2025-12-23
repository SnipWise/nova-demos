package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"let-s-chat/team"

	"github.com/joho/godotenv"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/agents/crew"
	"github.com/snipwise/nova/nova-sdk/agents/tools"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/logger"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/prompt"
)

func main() {
	// Create logger from environment variable
	log := logger.GetLoggerFromEnv()

	envFile := "ollama.env"
	//envFile := "llm-studio.env"
	//envFile := "docker.env"
	// Load environment variables from env file
	if err := godotenv.Load(envFile); err != nil {
		log.Error("Warning: Error loading env file: %v\n", err)
	}

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:11434/v1")

	ctx := context.Background()

	// ------------------------------------------------
	// Create the agent crew
	// ------------------------------------------------
	agentCrew := make(map[string]*chat.Agent)

	coderAgent, err := team.GetCoderAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}
	agentCrew["coder"] = coderAgent

	thinkerAgent, err := team.GetThinkerAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}
	agentCrew["thinker"] = thinkerAgent

	cookAgent, err := team.GetCookAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}
	agentCrew["cook"] = cookAgent

	genericAgent, err := team.GetGenericAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}
	agentCrew["generic"] = genericAgent

	// ------------------------------------------------
	// Define the function to match agent ID to topic
	// ------------------------------------------------
	matchAgentFunction := func(topic string) string {
		fmt.Println("🔵 Matching agent for topic:", topic)
		var agentId string
		switch strings.ToLower(topic) {
		case "coding", "programming", "development", "code", "software", "debugging", "technology", "software sevelopment":
			agentId = "coder"
		case "philosophy", "thinking", "ideas", "thoughts", "psychology", "relationships", "math", "mathematics", "science":
			agentId = "thinker"
		case "cooking", "recipe", "food", "culinary", "baking", "grilling", "meal":
			agentId = "cook"
		default:
			agentId = "generic"
		}
		fmt.Println("🟢 Matched agent ID:", agentId)
		return agentId
	}

	// ------------------------------------------------
	// Create the crew agent
	// ------------------------------------------------
	crewAgent, err := crew.NewAgent(
		ctx,
		agentCrew,
		"generic",
		matchAgentFunction,
		executeFunction,
		confirmationPromptFunction,
	)
	if err != nil {
		panic(err)
	}

	// Create the tools agent
	toolsAgent, err := team.GetToolsAgent(ctx, engineURL, GetToolsIndex())
	if err != nil {
		panic(err)
	}
	// Attach the tools agent to the server agent
	crewAgent.SetToolsAgent(toolsAgent)

	// Create the RAG agent
	ragAgent, err := team.GetRagAgent(ctx, engineURL, "./data")
	if err != nil {
		panic(err)
	}

	// Attach the RAG agent to the server agent
	crewAgent.SetRagAgent(ragAgent)

	// Create the compressor agent for text compression
	compressorAgent, err := team.GetCompressorAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}

	// Attach the compressor agent to the server agent
	crewAgent.SetCompressorAgent(compressorAgent)

	crewAgent.SetContextSizeLimit(8500)

	orchestratorAgent, err := team.GetOrchestratorAgent(ctx, engineURL)
	if err != nil {
		panic(err)
	}

	// Attach the orchestrator agent to the server agent
	crewAgent.SetOrchestratorAgent(orchestratorAgent)

	// ------------------------------------------------
	// Start the interactive chat loop
	// ------------------------------------------------
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
				display.NewLine()
			}
			return nil
		})

		if err != nil {
			display.Errorf("[%s][%v]failed to get completion: %v", crewAgent.GetName(), crewAgent.GetContextSize(), err)
			return
		}
		//display.NewLine()
		display.Separator()
		display.KeyValue("Finish reason", result.FinishReason)
		display.KeyValue("Context size", fmt.Sprintf("%d characters", crewAgent.GetContextSize()))
		display.Separator()
	}

}

// GetToolsIndex returns a list of available tools for the tools agent
func GetToolsIndex() []*tools.Tool {
	calculateSumTool := tools.NewTool("calculate_sum").
		SetDescription("Calculate the sum of two numbers").
		AddParameter("a", "number", "The first number", true).
		AddParameter("b", "number", "The second number", true)

	sayHelloTool := tools.NewTool("say_hello").
		SetDescription("Say hello to the given name").
		AddParameter("name", "string", "The name to greet", true)

	return []*tools.Tool{
		calculateSumTool,
		sayHelloTool,
	}
}

// executeFunction executes the given function with arguments and returns the result
func executeFunction(functionName string, arguments string) (string, error) {
	display.Colorf(display.ColorGreen, "🟢 Executing function: %s with arguments: %s\n", functionName, arguments)

	switch functionName {
	case "say_hello":
		var args struct {
			Name string `json:"name"`
		}
		if err := json.Unmarshal([]byte(arguments), &args); err != nil {
			return `{"error": "Invalid arguments for say_hello"}`, nil
		}
		hello := fmt.Sprintf("👋 Hello, %s!🙂", args.Name)
		return fmt.Sprintf(`{"message": "%s"}`, hello), nil

	case "calculate_sum":
		var args struct {
			A float64 `json:"a"`
			B float64 `json:"b"`
		}
		if err := json.Unmarshal([]byte(arguments), &args); err != nil {
			return `{"error": "Invalid arguments for calculate_sum"}`, nil
		}
		sum := args.A + args.B
		return fmt.Sprintf(`{"result": %g}`, sum), nil

	default:
		return `{"error": "Unknown function"}`, fmt.Errorf("unknown function: %s", functionName)
	}
}

// confirmationPromptFunction prompts the user for confirmation before executing a function
func confirmationPromptFunction(functionName string, arguments string) tools.ConfirmationResponse {
	display.Colorf(display.ColorGreen, "🟢 Detected function: %s with arguments: %s\n", functionName, arguments)

	choice := prompt.HumanConfirmation(fmt.Sprintf("Execute %s with %v?", functionName, arguments))
	return choice
}
