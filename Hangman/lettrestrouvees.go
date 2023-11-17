package hangman

import (
	"math/rand"
)

var LettresTrouvees []string
var LettresChoisis []string

// Si la lettre est trouvée, l'envoi dans le slice LettresTrouvees
func LettresTrouver(choixlettres string) []string {
	LettresTrouvees = append(LettresTrouvees, choixlettres)
	return LettresTrouvees
}

func LettreDejaChoisie(lettre string) bool {
	for _, l := range LettresChoisis {
		if l == lettre {
			return true
		}
	}
	return false
}

func ToutesLettresTrouvees(word string, lettresTrouvees []string) bool {
	for _, lettre := range word {
		if !Contains(lettresTrouvees, string(lettre)) {
			return false
		}
	}
	return true
}

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func LettresChoisi(choixlettres string) []string {
	LettresChoisis = append(LettresChoisis, choixlettres)
	return LettresChoisis
}

func ChoisirLettreAleatoire(word string) string {
	return string(word[rand.Intn(len(word))])
}
