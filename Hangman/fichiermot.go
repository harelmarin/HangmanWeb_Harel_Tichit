package hangman

import (
	"fmt"
	"math/rand"
	"os"
)

func (p *Global) RevealerRandomLetters(MotaDeviner string, count int) string {
	if count > len(MotaDeviner) {
		count = len(MotaDeviner)
	}

	// Convertir le mot en tranches de runes pour gérer les caractères Unicode.
	runes := []rune(MotaDeviner)

	// Générer des indices uniques aléatoires.
	indices := rand.Perm(len(MotaDeviner))[:count]

	// Révéler les lettres aux indices spécifiés.
	for _, idx := range indices {
		runes[idx] = []rune(MotaDeviner)[idx]
	}

	return string(runes)
}

// Fonction pour supprimer le fichier mot.txt
func SupprimerMot() {
	os.Remove("mot.txt")
}

// Fonction pour lire le mot dans le mot.txt
func LireFichierMot() string {
	ContenueFichierMot, err := os.ReadFile("mot.txt")
	if err != nil {
		fmt.Println("Erreur de lecture")
	}
	return string(ContenueFichierMot)
}
