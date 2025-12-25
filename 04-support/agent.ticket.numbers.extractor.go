package main

import (
	"context"

	"github.com/snipwise/nova/nova-sdk/agents"
	"github.com/snipwise/nova/nova-sdk/agents/structured"
	"github.com/snipwise/nova/nova-sdk/models"
	"github.com/snipwise/nova/nova-sdk/toolbox/files"
)

type Tickets struct {
	Numbers []string `json:"numbers"`
	Keywords []string `json:"keywords"`
}

/*
Cet agent extrait les numéros de ticket à partir d'un contenu.
Le contenu est le résultat d'une recherche de similarité dans la base des tickets.
Le contenu des similarités est agrégé et passé à cet agent pour extraire les numéros de ticket pertinents.
Les numéros de ticket extraits sont retournés sous forme de liste de chaînes de caractères.
Les instructions système pour l'agent sont chargées à partir d'un fichier markdown externe.
Ce fichier est défini dans "ticket.numbers.extraction.instructions.md" (dans le compose file)
et contient les directives nécessaires.

IMPORTANT: l'agent conserver l'historique des messages,
il faut donc penser à le réinitialiser pour éviter les effets de bord entre les appels.
=> KeepConversationHistory est mis à false
*/
func GetTicketNumbersExtractorAgent(ctx context.Context, engineURL, dataExtractorModelId string) (*structured.Agent[Tickets], error) {

	// Structured Extractor agent for ticket numbers (from user query)
	ticketNumbersExtractionInstructions, err := files.ReadTextFile("./ticket.numbers.extraction.instructions.md")
	if err != nil {
		panic(err)
	}
	ticketNumbersExtractorAgent, err := structured.NewAgent[Tickets](
		ctx,
		agents.Config{
			Name:               "ticket-numbers-extractor-agent",
			EngineURL:          engineURL,
			SystemInstructions: ticketNumbersExtractionInstructions,
			KeepConversationHistory: false,
		},
		models.Config{
			Name: dataExtractorModelId,
		},
	)
	if err != nil {
		return nil, err
	}

	return ticketNumbersExtractorAgent, nil
}
