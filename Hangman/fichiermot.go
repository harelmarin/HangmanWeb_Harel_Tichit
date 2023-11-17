package hangman

import (
	"fmt"
	"os"
)

// Prend le mot dans créé aléatoirement grâce à la fonction GetRandomWordFromFile
// Et l'écrit dans le mot.txt avec la fonction WriteFile
func Ecriremot() {
	filename := "Hangman/dico.txt"
	word, err := GetRandomWordFromFile(filename)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}
	os.WriteFile("mot.txt", []byte(word), 0644)
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
