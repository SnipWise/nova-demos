package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dataDir := "07-chaudieres/data"
	docsDir := "07-chaudieres/docs"

	// Créer le dossier docs s'il n'existe pas
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		fmt.Printf("Erreur création dossier docs: %v\n", err)
		return
	}

	// Pattern pour détecter le début d'une fiche (## FACT-CHAUD-XXX:)
	factPattern := regexp.MustCompile(`^## (FACT-CHAUD-\d+):\s*(.+)$`)

	// Lister tous les fichiers MD numérotés
	files := []string{
		"01_Combustion_Bruleur.md",
		"02_Circuit_Fumees.md",
		"03_Hydraulique.md",
		"04_ECS.md",
		"05_Regulation_Sondes.md",
		"06_Securite_Gaz.md",
		"07_Electronique_Cartes.md",
	}

	totalFacts := 0

	for _, filename := range files {
		filePath := filepath.Join(dataDir, filename)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Erreur ouverture %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var currentFact string
		var currentTitle string
		var currentContent strings.Builder
		inFact := false

		for scanner.Scan() {
			line := scanner.Text()

			// Vérifier si on trouve une nouvelle fiche
			if matches := factPattern.FindStringSubmatch(line); matches != nil {
				// Sauvegarder la fiche précédente si elle existe
				if inFact && currentFact != "" {
					saveFact(docsDir, currentFact, currentTitle, currentContent.String())
					totalFacts++
				}

				// Commencer une nouvelle fiche
				currentFact = matches[1]
				currentTitle = strings.TrimSpace(matches[2])
				currentContent.Reset()
				currentContent.WriteString(line + "\n")
				inFact = true
			} else if inFact {
				// Vérifier si on atteint la fin du fichier ou un séparateur
				if strings.HasPrefix(line, "---") {
					currentContent.WriteString(line + "\n")
					// Regarder si la ligne suivante est une nouvelle fiche
					continue
				} else if strings.HasPrefix(line, "*Fin du fichier") ||
				          strings.HasPrefix(line, "**Retour à l'index") {
					// Fin du fichier, sauvegarder la dernière fiche
					if currentFact != "" {
						saveFact(docsDir, currentFact, currentTitle, currentContent.String())
						totalFacts++
						inFact = false
					}
					break
				} else {
					currentContent.WriteString(line + "\n")
				}
			}
		}

		// Sauvegarder la dernière fiche du fichier si nécessaire
		if inFact && currentFact != "" {
			saveFact(docsDir, currentFact, currentTitle, currentContent.String())
			totalFacts++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Erreur lecture %s: %v\n", filename, err)
		}
	}

	fmt.Printf("✓ Extraction terminée: %d fiches créées dans %s\n", totalFacts, docsDir)
}

func saveFact(docsDir, factID, title, content string) {
	// Nettoyer le titre pour créer un nom de fichier valide
	sanitizedTitle := sanitizeFilename(title)
	filename := fmt.Sprintf("%s_%s.md", factID, sanitizedTitle)
	filePath := filepath.Join(docsDir, filename)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Erreur écriture %s: %v\n", filename, err)
		return
	}

	fmt.Printf("✓ Créé: %s\n", filename)
}

func sanitizeFilename(s string) string {
	// Remplacer les caractères non autorisés par des underscores
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "\\", "_")
	s = strings.ReplaceAll(s, ":", "_")
	s = strings.ReplaceAll(s, "*", "_")
	s = strings.ReplaceAll(s, "?", "_")
	s = strings.ReplaceAll(s, "\"", "_")
	s = strings.ReplaceAll(s, "<", "_")
	s = strings.ReplaceAll(s, ">", "_")
	s = strings.ReplaceAll(s, "|", "_")
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "-", "_")

	// Limiter la longueur du nom de fichier
	if len(s) > 80 {
		s = s[:80]
	}

	return s
}
