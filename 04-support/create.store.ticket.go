package main

import (
	"encoding/json"
	"fmt"

	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/spinner"
)

func CreateTicketsStoreWithVectors(ragTicketsAgent *rag.Agent, ticketMetaDataExtractorAgent *structured.Agent[TicketMetadata]) error {

	thinkingSpinner := spinner.NewWithColor("").SetSuffix("loading...").SetFrames(spinner.FramesDots)
	thinkingSpinner.SetSuffixColor(spinner.ColorBrightYellow).SetFrameColor(spinner.ColorBrightYellow)

	storePathTickets := env.GetEnvOrDefault("STORE_PATH_TICKETS", "./store/ticket_embeddings.json")

	if ragTicketsAgent.StoreFileExists(storePathTickets) {
		thinkingSpinner.Start()
		defer thinkingSpinner.Stop()
		// Load the RAG store from file
		err := ragTicketsAgent.LoadStore(storePathTickets)
		if err != nil {
			display.Errorf("failed to load RAG store from %s: %v\n", storePathTickets, err)
			return err
		}
		thinkingSpinner.Stop()
		display.Colorf(display.ColorGreen, "Successfully loaded RAG store from %s\n", storePathTickets)
		return nil
	} else {
		display.Colorf(display.ColorYellow, "RAG store file %s does not exist. A new store will be created.\n", storePathTickets)
		// Chunking + chunk enrichment

		tickets, err := files.GetContentFilesWithNames(env.GetEnvOrDefault("RAG_PATH_TICKETS", "./data/tickets"), ".md")
		if err != nil {
			display.Errorf("failed to get ticket files: %v\n", err)
			return err
		}
		for idx, ticket := range tickets {
			var ticketNumber string
			ticketNumber, err = ExtractTicketNumberStr(ticket.FileName)
			if err != nil {
				display.Errorf("failed to extract ticket number from filename %s: %v\n", ticket.FileName, err)
				// create a "no_ticket_number" to avoid nil pointer dereference
				ticketNumber = "no_ticket_number"

			}
			fmt.Println("Processing ticket file:", ticketNumber)

			// Extract metadata using the structured agent
			var metadata *TicketMetadata
			metadata, _, err = ticketMetaDataExtractorAgent.GenerateStructuredData([]messages.Message{
				{Role: roles.User, Content: ticket.Content},
			})
			if err != nil {
				display.Errorf("failed to extract metadata for ticket %s: %v\n", ticketNumber, err)
				// create empty metadata to avoid nil pointer dereference
				metadata = &TicketMetadata{}
				// do not return err
			}
			metadata.Number = ticketNumber

			// Display extracted metadata
			metadataJson, _ := json.MarshalIndent(metadata, "", "  ")
			display.Colorf(display.ColorCyan, "Extracted metadata for ticket %s: %s\n", ticketNumber, string(metadataJson))

			err = ragTicketsAgent.SaveEmbedding(string(metadataJson) + "\n" + ticket.Content)
			if err != nil {
				display.Errorf("failed to save embedding for document %d: %v\n", idx, err)
				// do not return err
			} else {
				display.Colorf(display.ColorGreen, "Successfully saved embedding for ticket %s\n", ticketNumber)
			}

			// IMPORTANT: in theory I do not need to reset the messages because the structured agent does not keep state
			//ticketMetaDataExtractorAgent.ResetMessages()
		}
		// Save the RAG store to file
		err = ragTicketsAgent.PersistStore(storePathTickets)
		if err != nil {
			display.Errorf("failed to persist RAG store to %s: %v\n", storePathTickets, err)
			return err
		}
		display.Colorf(display.ColorGreen, "Successfully saved RAG Tickets store to %s\n", storePathTickets)
		return nil
	}
}
