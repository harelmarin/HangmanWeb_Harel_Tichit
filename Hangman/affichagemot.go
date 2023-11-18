package hangman

import (
	"fmt"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func (p *Global) AffichageTirets() string {
	var result string
	lettresfind := LettresTrouvees
	for i := range p.MotaDeviner {
		lettreTrouvee := false
		for _, lettre := range lettresfind {
			if p.MotaDeviner[i:i+1] == lettre {
				result += fmt.Sprintf("%c ", p.MotaDeviner[i])
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
