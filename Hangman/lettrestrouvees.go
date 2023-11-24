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

func (p *Global) LettreDejaChoisie(LettreJoueur string) bool {
	for _, l := range User.LettreDejaJoue {
		if l == LettreJoueur {
			return true
		}
	}
	return false
}

func (p *Global) ToutesLettresTrouvees(MotaDeviner string, LettreDejaTrouve []string) bool {
	for _, lettre := range MotaDeviner {
		if !Contains(LettreDejaTrouve, string(lettre)) {
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

func (p *Global) ChoisirLettreAleatoire(MotaDeviner string) string {
	return string(MotaDeviner[rand.Intn(len(MotaDeviner))])
}
