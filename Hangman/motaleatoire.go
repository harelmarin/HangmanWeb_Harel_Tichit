package hangman

import (
	"bufio"
	"math/rand"
	"os"
)

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

