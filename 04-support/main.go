package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/prompt"
	"github.com/snipwise/nova/nova-sdk/ui/spinner"
)

func main() {
	// Start spinner for thinking/loading
	thinkingSpinner := spinner.NewWithColor("").SetSuffix("loading...").SetFrames(spinner.FramesDots)
	thinkingSpinner.SetSuffixColor(spinner.ColorBrightYellow).SetFrameColor(spinner.ColorBrightYellow)

	engineURL := env.GetEnvOrDefault("ENGINE_URL", "http://localhost:11434/v1")
	dataExtractorModelId := env.GetEnvOrDefault("DATA_EXTRACTOR_MODEL", "hf.co/menlo/jan-nano-gguf:q4_k_m")
	ragModelId := env.GetEnvOrDefault("RAG_MODEL", "huggingface.co/mixedbread-ai/mxbai-embed-large-v1:f16")
	reporterModelId := env.GetEnvOrDefault("REPORTER_MODEL", "hf.co/menlo/lucy-gguf:q4_k_m")
	//reporterBiggerModelId := env.GetEnvOrDefault("REPORTER_BIGGER_MODEL", "ai/qwen2.5:latest")


	ctx := context.Background()

	// ------------------------------------------------
	// STRUCTURED AGENT: (TicketMetadata)
	// Structured Extractor agent for ticket metadata
	// ------------------------------------------------
	ticketMetaDataExtractorAgent, err := GetTicketMetaDataExtractorAgent(ctx, engineURL, dataExtractorModelId)

	// ------------------------------------------------
	// STRUCTURED AGENT: (Tickets)
	// Structured Extractor agent for ticket numbers
	// ------------------------------------------------
	ticketNumbersExtractorAgent, err := GetTicketNumbersExtractorAgent(ctx, engineURL, dataExtractorModelId)

	// ------------------------------------------------
	// STRUCTURED AGENT: (ReportMetaData)
	// Structured Extractor agent for report metadata
	// ------------------------------------------------
	reportMetaDataExtractorAgent, err := GetReportMetaDataExtractorAgent(ctx, engineURL, dataExtractorModelId)

	// ------------------------------------------------
	// RAG AGENT: for tickets
	// ------------------------------------------------
	ragTicketsAgent, err := GetRagTicketsAgent(ctx, engineURL, ragModelId)

	// ------------------------------------------------
	// RAG AGENT: for reports
	// ------------------------------------------------
	ragReportsAgent, err := GetRagReportsAgent(ctx, engineURL, ragModelId)

	// ------------------------------------------------
	// CHAT AGENT: reporter agent for tickets
	// ------------------------------------------------
	ticketsReportUserMessageTemplate, err := files.ReadTextFile("./ticket.report.user.message.md")

	ticketsReporterAgent, err := GetTicketsReporterAgent(ctx, engineURL, reporterModelId)

	// ------------------------------------------------
	// CHAT AGENT: reporter agent for solution reports
	// ------------------------------------------------
	solutionsReportUserMessageTemplate, err := files.ReadTextFile("./solution.report.user.message.md")

	solutionReportsReporterAgent, err := GetSolutionReportsReporterAgent(ctx, engineURL, reporterModelId)
	//solutionReportsReporterAgent, err := GetSolutionReportsReporterAgent(ctx, engineURL, reporterBiggerModelId)

	

	if err != nil {
		panic(err)
	}

	// --------------------------------------------
	// STORE: Load or create the RAG store for reports
	// --------------------------------------------

	errRagStoreReports := CreateReportsStoreWithVectors(ragReportsAgent, reportMetaDataExtractorAgent)
	if errRagStoreReports != nil {
		display.Errorf("failed to create/load RAG store for reports: %v\n", errRagStoreReports)
		panic(errRagStoreReports)
	}

	// --------------------------------------------
	// STORE: Load or create the RAG store for tickets
	// --------------------------------------------

	errRagStoreTickets := CreateTicketsStoreWithVectors(ragTicketsAgent, ticketMetaDataExtractorAgent)
	if errRagStoreTickets != nil {
		display.Errorf("failed to create/load RAG store for tickets: %v\n", errRagStoreTickets)
		panic(errRagStoreTickets)
	}

	// --------------------------------------------
	// Start the interactive chat loop
	// --------------------------------------------
	for { // BEGIN: for loop

		input := prompt.NewWithColor("🤖 What is your issue?")
		question, err := input.RunWithEdit()

		if err != nil {
			display.Errorf("failed to get input: %v", err)
			return
		}
		if strings.HasPrefix(question, "/bye") {
			display.Infof("👋 Goodbye!")
			break
		}

		display.NewLine()

		// --------------------------------------------
		// Search similar tickets
		// --------------------------------------------
		thinkingSpinner.SetSuffix("searching tickets...")
		thinkingSpinner.Start()

		similarities, err := ragTicketsAgent.SearchTopN(question, 0.6, 4)

		thinkingSpinner.Stop()

		//similarities, err := ragTicketsAgent.SearchSimilar(question, 0.6)
		if err != nil {
			display.Errorf("failed to search similar tickets: %v", err)
			continue
		}

		similaritiesTicketsContext := ""
		display.Colorf(display.ColorGreen, "Top similar %v tickets:\n", len(similarities))
		for _, sim := range similarities {
			display.Colorf(display.ColorCyan, "Content: %s\n", sim.Prompt)
			display.Colorf(display.ColorYellow, "Score: %f\n", sim.Similarity)
			similaritiesTicketsContext += sim.Prompt + "\n---\n"
		}

		// Extract ticket numbers from the user question
		thinkingSpinner.SetSuffix("extracting ticket numbers...")
		thinkingSpinner.Start()
		var extractedTickets *Tickets

		extractedTickets, _, err = ticketNumbersExtractorAgent.GenerateStructuredData([]messages.Message{
			{Role: roles.User, Content: similaritiesTicketsContext},
		})
		if err != nil {
			display.Errorf("failed to extract ticket numbers from question: %v\n", err)
			continue
		}
		thinkingSpinner.Stop()

		//ticketNumbersExtractorAgent.ResetMessages()

		// --------------------------------------------
		// NOTE: Similar Ticket numbers extracted
		// --------------------------------------------
		// IMPORTANT: to be used in the solution report search (similarities)
		ticketNumbers := extractedTickets.Numbers
		ticketKeywords := extractedTickets.Keywords

		display.NewLine()
		display.Colorf(display.ColorBrightGreen, "Extracted ticket numbers: %v\n", ticketNumbers)
		display.Colorf(display.ColorGreen, "Extracted ticket keywords: %v\n", ticketKeywords)
		display.NewLine()

		markdownParser := display.NewMarkdownChunkParser()

		thinkingSpinner.SetSuffix("generating similar tickets report...")
		thinkingSpinner.Start()

		// --------------------------------------------
		// Generate ticket report from the similarity search results
		// --------------------------------------------
		userTicketReportMessage := fmt.Sprintf(
			ticketsReportUserMessageTemplate,
			similaritiesTicketsContext,
		)

		// --------------------------------------------
		// Generate and Stream the report
		// --------------------------------------------
		result, err := ticketsReporterAgent.GenerateStreamCompletion(
			[]messages.Message{
				{
					Role:    roles.User,
					Content: userTicketReportMessage,
				},
			},
			func(chunk string, finishReason string) error {
				if thinkingSpinner.IsRunning() && finishReason == "" {
					thinkingSpinner.Success("Model loaded!")
					thinkingSpinner.Stop()
				}
				// Use markdown chunk parser for colorized streaming output
				if chunk != "" {
					display.MarkdownChunk(markdownParser, chunk)
				}
				if finishReason == "stop" {
					markdownParser.Flush()
					fmt.Println()
				}
				return nil
			},
		)
		_ = result
		if err != nil {
			thinkingSpinner.Stop()
			display.NewLine()
			display.Errorf("failed to generate tickets report: %v", err)
			continue
		}

		//ticketsReporterAgent.ResetMessages()

		display.NewLine()

		// --------------------------------------------
		// Search for related reports
		// --------------------------------------------
		thinkingSpinner.SetSuffix("searching reports...")

		// IMPORTANT:
		// TODO: try the search with one passage combining all ticket numbers

		// Combine all ticket numbers and keywords into a single search string
		similaritiesReportsContext := ""

		// ---BEGIN: [WIP]-----
		ticketNumbersCombined := strings.Join(ticketNumbers, "; ") // extractedTickets.Numbers
		keywordsCombined := strings.Join(ticketKeywords, "; ")     // extractedTickets.Keywords

		display.NewLine()
		display.Colorf(display.ColorBrightYellow,
			`🔎 Searching reports related to ticket numbers: %s and keywords: %s`,
			ticketNumbersCombined,
			keywordsCombined,
		)
		// TODO: IMPORTANT: try other queries formulations here
		// QUESTION: how to handle the case where there is no ticket number extracted?the image is frozen 
		similarReports, err := ragReportsAgent.SearchTopN(
			"Resolved tickets: "+ticketNumbersCombined + " with keywords: " + keywordsCombined, 
			0.6, 
			6,
		)

		// similarReports, err := ragReportsAgent.SearchTopN(
		// 	"Find these tickets if resolved: "+ticketNumbersCombined, 
		// 	0.5, 
		// 	6,
		// )		
		
		if err != nil {
			display.Errorf("failed to search similar reports for tickets %s: %v", ticketNumbersCombined, err)
			continue
		}
		for _, sim := range similarReports {
			display.Colorf(display.ColorGreen, "\n\nRelated report:\n")
			display.Colorf(display.ColorCyan, "Content: %s\n", Truncate100(sim.Prompt))
			display.Colorf(display.ColorYellow, "Score: %f\n", sim.Similarity)
			similaritiesReportsContext += sim.Prompt + "\n---\n"

		}
		display.NewLine()

		// ---END: [WIP]-----

		// for _, ticketNumber := range extractedTickets.Numbers {

		// 	display.Colorf(display.ColorGreen, "🔎 Searching reports related to ticket number: %s\n", ticketNumber)
		// 	similarities, err := ragReportsAgent.SearchTopN("Resolved ticket: "+ticketNumber, 0.6, 4)
		// 	if err != nil {
		// 		display.Errorf("failed to search similar reports for ticket %s: %v", ticketNumber, err)
		// 		continue
		// 	}
		// 	for _, sim := range similarities {
		// 		display.Colorf(display.ColorGreen, "Related report for ticket %s:\n", ticketNumber)
		// 		display.Colorf(display.ColorCyan, "Content: %s\n", Truncate100(sim.Prompt))
		// 		display.Colorf(display.ColorYellow, "Score: %f\n", sim.Similarity)
		// 		similaritiesReportsContext += sim.Prompt + "\n---\n"

		// 	}
		// 	display.NewLine()

		// }

		thinkingSpinner.SetSuffix("generating solution report...")
		thinkingSpinner.Start()

		// --------------------------------------------
		// Generate final solution report
		// --------------------------------------------
		markdownParser = display.NewMarkdownChunkParser()

		userMessageSolutionReport := fmt.Sprintf(
			solutionsReportUserMessageTemplate,
			similaritiesReportsContext,
		)

		result, err = solutionReportsReporterAgent.GenerateStreamCompletion(
			[]messages.Message{
				{
					Role:    roles.User,
					Content: userMessageSolutionReport,
				},
			},
			func(chunk string, finishReason string) error {
				if thinkingSpinner.IsRunning() && finishReason == "" {
					thinkingSpinner.Success("Model loaded!")
					thinkingSpinner.Stop()
				}
				// Use markdown chunk parser for colorized streaming output
				if chunk != "" {
					display.MarkdownChunk(markdownParser, chunk)
				}
				if finishReason == "stop" {
					markdownParser.Flush()
					fmt.Println()
				}
				return nil
			},
		)

		if err != nil { // NOTE: error because of the context size exceeded:
			/*
				✗ failed to generate final solution report:
				POST "http://localhost:12434/engines/llama.cpp/v1/chat/completions":
				400 Bad Request
				{"code":400,"message":"the request exceeds the available context size, try increasing it","type":"exceed_context_size_error","n_prompt_tokens":5860,"n_ctx":4096}

				SOLUTION:
				- increase the context size of the model: ie: context_size: 32768
				- we need to be sure that the model used support the context size (8k, 32k, 131k, etc.)
				- use a model with a bigger context size (e.g., lucy-128k-gguf)
			*/
			thinkingSpinner.Stop()
			display.NewLine()
			display.Errorf("failed to generate final solution report: %v", err)
			continue
		}
		//solutionReportsReporterAgent.ResetMessages()

	} // END: for loop

}

// ExtractTicketNumber extrait le numéro d'un chemin de ticket (retourne un entier)
// Exemple: "data/tickets/ticket_089.md" -> 89
func ExtractTicketNumber(path string) (int, error) {
	// Récupérer le nom du fichier sans l'extension
	filename := filepath.Base(path)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Extraire le numéro en enlevant le préfixe "ticket_"
	numeroStr := strings.TrimPrefix(name, "ticket_")

	// Vérifier que le préfixe a bien été enlevé
	if numeroStr == name {
		return 0, fmt.Errorf("format invalide: préfixe 'ticket_' non trouvé")
	}

	// Convertir en entier
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		return 0, fmt.Errorf("impossible de convertir '%s' en nombre: %w", numeroStr, err)
	}

	return numero, nil
}

// ExtractTicketNumberStr extrait le numéro d'un chemin de ticket (retourne une string avec les zéros)
// Exemple: "data/tickets/ticket_089.md" -> "089"
func ExtractTicketNumberStr(path string) (string, error) {
	// Récupérer le nom du fichier sans l'extension
	filename := filepath.Base(path)
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	// Extraire le numéro en enlevant le préfixe "ticket_"
	numeroStr := strings.TrimPrefix(name, "ticket_")

	// Vérifier que le préfixe a bien été enlevé
	if numeroStr == name {
		return "", fmt.Errorf("format invalide: préfixe 'ticket_' non trouvé")
	}

	// Vérifier que c'est bien un nombre (sans le convertir)
	if _, err := strconv.Atoi(numeroStr); err != nil {
		return "", fmt.Errorf("format invalide: '%s' n'est pas un nombre", numeroStr)
	}

	return numeroStr, nil
}

func Truncate100(s string) string {
	runes := []rune(s)
	if len(runes) <= 100 {
		return s
	}
	return string(runes[:100]) + "..."
}
