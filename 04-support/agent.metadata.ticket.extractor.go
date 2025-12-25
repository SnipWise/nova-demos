package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
)

type TicketMetadata struct {
	Number           string   `json:"number"`
	Title            string   `json:"title"`
	Keywords         []string `json:"keywords"`
	ShortDescription string   `json:"short_description"`
	Language         string   `json:"language"`
}

/*
Cet agent permet d'extraire les métadonnées d'un ticket à partir de son contenu textuel.
Les métadonnées extraites incluent :
- Number : le numéro du ticket
- Title : le titre du ticket
- Keywords : une liste de mots-clés associés au ticket
- ShortDescription

Les instructions système pour l'agent sont chargées à partir d'un fichier markdown externe.
Ce fichier est défini dans "ticket.data.extraction.instructions.md" (dans le compose file)
et contient les directives nécessaires

IMPORTANT: l'agent conserver l'historique des messages,
il faut donc penser à le réinitialiser pour éviter les effets de bord entre les appels.
=> KeepConversationHistory est mis à false
*/
func GetTicketMetaDataExtractorAgent(ctx context.Context, engineURL, dataExtractorModelId string) (*structured.Agent[TicketMetadata], error) {
	// Structured Extractor agent for tickets (to fill meta data fields)
	ticketDataExtractionInstructions, err := files.ReadTextFile("./ticket.data.extraction.instructions.md")
	if err != nil {
		panic(err)
	}

	ticketDataExtractorAgent, err := structured.NewAgent[TicketMetadata](
		ctx,
		agents.Config{
			Name:               "ticket-data-extractor-agent",
			EngineURL:          engineURL,
			SystemInstructions: ticketDataExtractionInstructions,
			KeepConversationHistory: false,
		},
		models.Config{
			Name: dataExtractorModelId,
		},
	)
	if err != nil {
		return nil, err
	}

	return ticketDataExtractorAgent, nil
}
