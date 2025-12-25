package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
)

type ReportMetaData struct {
	ReportTitle           string   `json:"report_title"`
	ReportSummary         string   `json:"report_summary"`
	ResolvedTicketNumbers []string `json:"resolved_ticket_numbers"`
	Keywords              []string `json:"keywords"`
	ShortDescription      string   `json:"short_description"`
	MainTopic             string   `json:"main_topic"`
}

/*
Cet agent permet d'extraire les méta-données d'un rapport donné.
Les méta-données extraites incluent:
- ReportTitle : le titre du rapport
- ReportSummary : un résumé concis du rapport
- ResolvedTicketNumbers : une liste des numéros de tickets résolus dans le rapport
- Keywords : une liste de mots-clés associés au rapport
- ShortDescription : une brève description du rapport
- MainTopic : le sujet principal du rapport

Les instructions système pour l'agent sont chargées à partir d'un fichier markdown externe.
Ce fichier est défini dans "report.data.extraction.instructions.md" (dans le compose file)
et contient les directives nécessaires.

IMPORTANT: l'agent conserver l'historique des messages,
il faut donc penser à le réinitialiser pour éviter les effets de bord entre les appels.
=> KeepConversationHistory est mis à false
*/
func GetReportMetaDataExtractorAgent(ctx context.Context, engineURL, dataExtractorModelId string) (*structured.Agent[ReportMetaData], error) {
	// Structured Extractor agent for reports (to fill meta data fields)
	reportMetaDataExtractionInstructions, err := files.ReadTextFile("./report.data.extraction.instructions.md")
	if err != nil {
		panic(err)
	}
	reportMetaDataExtractorAgent, err := structured.NewAgent[ReportMetaData](
		ctx,
		agents.Config{
			Name:                    "report-data-extractor-agent",
			EngineURL:               engineURL,
			SystemInstructions:      reportMetaDataExtractionInstructions,
			KeepConversationHistory: false,
		},
		models.Config{
			Name: dataExtractorModelId,
		},
	)
	if err != nil {
		return nil, err
	}

	return reportMetaDataExtractorAgent, nil
}
