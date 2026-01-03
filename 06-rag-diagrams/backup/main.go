package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"diagram/config"
	"diagram/diagram"
	"diagram/validator"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/chat"
	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/rag/chunks"
	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
)

func main() {
	ctx := context.Background()

	// Enable logging
	if err := os.Setenv("NOVA_LOG_LEVEL", "INFO"); err != nil {
		panic(err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize RAG agent
	ragAgent, err := initializeRAGAgent(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// Load or create RAG store
	if err := setupRAGStore(ragAgent, cfg); err != nil {
		panic(err)
	}

	// Define query
	query := `
	Create a class diagram of a warrior that inherits from a character class. 
	The warrior has attributes strength and weapon, and methods attack() and defend(). 
	The character class has attributes name and level, and method move().
	Additionally, include a mage class that also inherits from character,
	with attributes mana and spell, and methods castSpell() and teleport().
	And add the arrow symbols (△) to indicate inheritance relationships.
	- Mage inherits from Character
	- Warrior inherits from Character
	
	`

	// Search for similar diagrams
	similarityContext, err := searchSimilarDiagrams(ragAgent, query)
	if err != nil {
		panic(err)
	}

	// Initialize chat agent
	chatAgent, err := initializeChatAgent(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// Generate diagram with streaming
	result, err := generateDiagram(chatAgent, query, similarityContext)
	if err != nil {
		panic(err)
	}

	displayContextSize(chatAgent, result)

	// Post-processing: Auto-fix alignment
	//processAndDisplayDiagram(result.Response)
}

// initializeRAGAgent creates and configures the RAG agent
func initializeRAGAgent(ctx context.Context, cfg *config.Config) (*rag.Agent, error) {
	return rag.NewAgent(
		ctx,
		agents.Config{
			EngineURL: cfg.EngineURL,
		},
		models.Config{
			Name: cfg.RAGModelID,
		},
	)
}

// setupRAGStore loads existing RAG store or creates a new one
func setupRAGStore(ragAgent *rag.Agent, cfg *config.Config) error {
	if ragAgent.StoreFileExists(cfg.StorePathFile) {
		// Load the RAG store from file
		err := ragAgent.LoadStore(cfg.StorePathFile)
		if err != nil {
			return fmt.Errorf("failed to load RAG store from %s: %w", cfg.StorePathFile, err)
		}
		fmt.Printf("Successfully loaded RAG store from %s\n", cfg.StorePathFile)
		return nil
	}

	// Create new store
	fmt.Printf("RAG store file %s does not exist. A new store will be created.\n", cfg.StorePathFile)

	// Chunking + chunk enrichment
	filesContent, err := files.GetContentFilesWithNames(cfg.DataPath, ".md")
	if err != nil {
		return fmt.Errorf("failed to get files: %w", err)
	}

	for idx, content := range filesContent {
		contentPieces := chunks.SplitMarkdownBySections(content.Content)

		for _, piece := range contentPieces {
			err = ragAgent.SaveEmbedding(piece)
			if err != nil {
				fmt.Printf("failed to save embedding for document %d: %v\n", idx, err)
			} else {
				fmt.Printf("Successfully saved embedding for file %s (chunk)\n", content.FileName)
			}
		}
	}

	// Save the RAG store to file
	err = ragAgent.PersistStore(cfg.StorePathFile)
	if err != nil {
		return fmt.Errorf("failed to persist RAG store to %s: %w", cfg.StorePathFile, err)
	}
	fmt.Printf("Successfully saved RAG store to %s\n", cfg.StorePathFile)
	return nil
}

// searchSimilarDiagrams searches for similar diagrams in the RAG store
func searchSimilarDiagrams(ragAgent *rag.Agent, query string) (string, error) {
	similarities, err := ragAgent.SearchSimilar(query, 0.5)
	if err != nil {
		return "", err
	}

	var similarityContext string
	display.Colorf(display.ColorGreen, "📝 Similar diagrams for query: %s\n", query)
	for _, sim := range similarities {
		display.Separator()
		display.Colorf(display.ColorCyan, "Content: %s\n", sim.Prompt)
		display.Colorf(display.ColorYellow, "Score: %f\n", sim.Similarity)
		similarityContext += sim.Prompt + "\n\n"
	}

	return similarityContext, nil
}

// initializeChatAgent creates and configures the chat agent
func initializeChatAgent(ctx context.Context, cfg *config.Config) (*chat.Agent, error) {
	return chat.NewAgent(
		ctx,
		agents.Config{
			Name:                    "diagram-assistant",
			EngineURL:               cfg.EngineURL,
			SystemInstructions:      config.SystemInstructions(),
			KeepConversationHistory: false,
		},
		models.Config{
			Name:        cfg.ChatModelID,
			Temperature: models.Float64(0.0), // Zero temperature for consistent alignment
		},
	)
}

// generateDiagram generates a diagram using the chat agent with streaming
func generateDiagram(agent *chat.Agent, query string, similarityContext string) (*chat.ReasoningResult, error) {
	return agent.GenerateStreamCompletionWithReasoning(
		[]messages.Message{
			{
				Role: roles.System,
				Content: fmt.Sprintf(`
					## Reference examples (retrieved by RAG)
					%s

					## Self-verification before output
					Before responding:
					1. Count the characters in EACH line of the diagram
					2. If two lines in the same box have different lengths, adjust the spaces
					3. Only send the response when ALL lines in the same box have the same length

					## Character count validation
					For each box, verify:
					- Top border length = N characters
					- Each middle line length = N characters (achieved by adding spaces before │)
					- Bottom border length = N characters

					If any line is shorter: ADD SPACES before the closing │
					If any line is longer: SHORTEN the text or use abbreviations

					`,
					similarityContext,
				),
			},
			{
				Role: roles.User,
				Content: fmt.Sprintf(`
					## User request
					%s

					## Instructions
					- COPY the STRUCTURE from the most similar example
					- You MUST COPY IDENTICALLY all structure characters (borders, corners, lines)
					- You are ONLY allowed to modify the text BETWEEN the borders
					- FORBIDDEN to modify the number of spaces or position of border characters

					- Modify ONLY the LABELS (text content)
					- Preserve alignment and spacing
					- Wrap the code with triple backticks

					- Strictly respect the width of provided examples
					- If the new label is shorter, fill with spaces
					- If the label is longer than the box, ABBREVIATE it rather than expanding the box

					## Special instruction for UML inheritance
					- Use the △ symbol to show inheritance
					- Place the parent class above
					- Draw a line with △ pointing to the parent
					- The child class should NOT repeat inherited attributes/methods

					`,
					query,
				),
			},
		},
		func(reasoningChunk string, finishReason string) error {
			display.Color(reasoningChunk, display.ColorYellow)
			if finishReason != "" {
				display.NewLine()
				display.KeyValue("Finish reason", finishReason)
			}
			return nil
		},
		func(chunk string, finishReason string) error {
			if chunk != "" {
				fmt.Print(chunk)
			}
			if finishReason == "stop" {
				fmt.Println()
			}
			return nil
		},
	)
}

// displayContextSize displays context size information
func displayContextSize(agent *chat.Agent, result *chat.ReasoningResult) {
	fmt.Println()
	fmt.Println("Finish reason:\n", result.FinishReason)
	fmt.Printf("Context size: %d characters\n", agent.GetContextSize())
	fmt.Println(strings.Repeat("-", 40))
}

// processAndDisplayDiagram extracts, validates, and optionally fixes the diagram
func processAndDisplayDiagram(response string) {
	// Extract diagram from response
	rawDiagram := diagram.ExtractCodeBlock(response)

	// Show raw output analysis
	validator.PrintDiagramAnalysis(rawDiagram, "RAW OUTPUT ANALYSIS")

	// Validate raw output
	isValid, issues := validator.ValidateAlignment(rawDiagram)
	if isValid {
		fmt.Println("\n✅ Raw output is already properly aligned!")
		fmt.Println("\nFINAL DIAGRAM:")
		fmt.Println(rawDiagram)
		return
	}

	// Display alignment issues (minor - acceptable)
	fmt.Println("\n⚠️  Raw output has minor alignment variations (acceptable):")
	for boxNum, lengths := range issues {
		fmt.Printf("  Box %d: lengths vary %v\n", boxNum, lengths)
	}

	// Display final diagram (no auto-fix to preserve complex structure)
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("FINAL DIAGRAM")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println(rawDiagram)
}
