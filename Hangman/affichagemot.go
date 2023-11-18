package hangman

import (
	"fmt"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func AffichageTirets(Word string) string {
	var result string
	lettresfind := LettresTrouvees
	for i := range Word {
		lettreTrouvee := false
		for _, lettre := range lettresfind {
			if Word[i:i+1] == lettre {
				result += fmt.Sprintf("%c ", Word[i])
				lettreTrouvee = true
				break
			}
		}
		if !lettreTrouvee {
			result += "- "
		}
	}
	return result
}
