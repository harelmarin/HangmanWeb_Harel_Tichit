package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// Affiche le mot choisi en tiret avec un espace entre chaque tiret
func (p *Global) AffichageTirets() string {
	var result string
	lettresfind := User.LettreDejaTrouve
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

func (p *Global) Ecriremot() {
	filename := "Hangman/dico.txt"
	word, err := GetRandomWordFromFile(filename)
	if err != nil {
		fmt.Println("Erreur:", err)
		return
	}

	os.WriteFile("mot.txt", []byte(word), 0644)
}

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

// FONCTION POUR PRENDRE UN MOT DANS LE FICHIER DICO.TXT GRACE AU NOMBRE RANDOM
func GetRandomWordFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	randomIndex := GenerateRandomNumber(len(words))

	return words[randomIndex], nil
}

// FONCTION QUI GENERE LE NOMBRE RANDOM
func GenerateRandomNumber(max int) int {
	return rand.Intn(150)
}

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
