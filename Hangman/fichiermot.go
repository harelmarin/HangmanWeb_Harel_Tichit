package hangman

import (
	"fmt"
	"os"
	"math/rand"
)

// Prend le mot dans créé aléatoirement grâce à la fonction GetRandomWordFromFile
// Et l'écrit dans le mot.txt avec la fonction WriteFile


func (p *Global) RevealerRandomLetters(word string, count int) string {
	if count > len(word) {
	   count = len(word)
	}
 
	// Convertir le mot en tranches de runes pour gérer les caractères Unicode.
	runes := []rune(word)
 
	// Générer des indices uniques aléatoires.
	indices := rand.Perm(len(word))[:count]
 
	// Révéler les lettres aux indices spécifiés.
	for _, idx := range indices {
	   runes[idx] = []rune(word)[idx]
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
