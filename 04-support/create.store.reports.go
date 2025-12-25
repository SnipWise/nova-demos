package main

import (
	"encoding/json"

	"github.com/snipwise/nova/nova-sdk/agents/rag"
	"github.com/snipwise/nova/nova-sdk/agents/rag/chunks"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/messages"
	"github.com/snipwise/nova/nova-sdk/messages/roles"
	"github.com/snipwise/nova/nova-sdk/toolbox/env"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
	"github.com/snipwise/nova/nova-sdk/ui/display"
	"github.com/snipwise/nova/nova-sdk/ui/spinner"
)

func CreateReportsStoreWithVectors(ragReportsAgent *rag.Agent, reportMetaDataExtractorAgent *structured.Agent[ReportMetaData]) error {

	thinkingSpinner := spinner.NewWithColor("").SetSuffix("loading...").SetFrames(spinner.FramesDots)
	thinkingSpinner.SetSuffixColor(spinner.ColorBrightYellow).SetFrameColor(spinner.ColorBrightYellow)

	storePathReports := env.GetEnvOrDefault("STORE_PATH_REPORTS", "./store/report_embeddings.json")

	if ragReportsAgent.StoreFileExists(storePathReports) {
		thinkingSpinner.Start()
		defer thinkingSpinner.Stop()
		// Load the RAG store from file
		err := ragReportsAgent.LoadStore(storePathReports)
		if err != nil {
			thinkingSpinner.Stop()
			display.Errorf("failed to load RAG store from %s: %v\n", storePathReports, err)
			return err
		}
		thinkingSpinner.Stop()
		display.Colorf(display.ColorGreen, "Successfully loaded RAG store from %s\n", storePathReports)
		return nil
	} else {
		display.Colorf(display.ColorYellow, "RAG store file %s does not exist. A new store will be created.\n", storePathReports)
		// Chunking + chunk enrichment

		reports, err := files.GetContentFilesWithNames(env.GetEnvOrDefault("RAG_PATH_REPORTS", "./data/reports"), ".md")
		if err != nil {
			display.Errorf("failed to get report files: %v\n", err)
			return err
		}
		for idx, report := range reports {

			var metadata *ReportMetaData
			metadata, _, err = reportMetaDataExtractorAgent.GenerateStructuredData([]messages.Message{
				{Role: roles.User, Content: report.Content},
			})
			if err != nil {
				display.Errorf("failed to extract metadata for report %s: %v\n", report.FileName, err)
				// create empty metadata to avoid nil pointer dereference
				metadata = &ReportMetaData{}
			}

			// Display extracted metadata
			metadataJson, _ := json.MarshalIndent(metadata, "", "  ")
			display.Colorf(display.ColorBrightPurple, "Extracted metadata for report %s: %s\n", report.FileName, string(metadataJson))

			// IMPORTANT: split reports into chunks because they can be very long
			contentPieces := chunks.ChunkText(report.Content, 768, 128)

			for _, piece := range contentPieces {
				err = ragReportsAgent.SaveEmbedding(string(metadataJson) + "\n" + piece)
				if err != nil {
					display.Errorf("failed to save embedding for document %d: %v\n", idx, err)
				} else {
					display.Colorf(display.ColorGreen, "Successfully saved embedding for report %s (piece)\n", report.FileName)
				}
			}

			// IMPORTANT: in theory I do not need to reset the messages because the structured agent does not keep state
			//reportMetaDataExtractorAgent.ResetMessages()

		}
		// Save the RAG store to file
		err = ragReportsAgent.PersistStore(storePathReports)
		if err != nil {
			display.Errorf("failed to persist RAG store to %s: %v\n", storePathReports, err)
			return err
		}
		display.Colorf(display.ColorGreen, "Successfully saved RAG Reports store to %s\n", storePathReports)
		return nil
	}
}
