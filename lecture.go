package projethangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func LireFichier(fichier string) ([]string, error) {
	texte, erreur := os.Open(fichier)
	if erreur != nil {
		return nil, erreur
	}
	defer texte.Close()
	var mots []string
	scanner := bufio.NewScanner(texte)

	for scanner.Scan() {
		mots = append(mots, strings.TrimSpace(scanner.Text()))
	}
	if erreur := scanner.Err(); erreur != nil {
		return nil, erreur
	}
	return mots, nil
}
func MotAleatoire(mots []string) string {
	var index = 0
	index = (index + 1) % len(mots)
	return mots[index]
}
