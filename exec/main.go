package main

import (
	"fmt"
	"os"
	"projethangman"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run exec/main.go <fichier_de_mots>")
		return
	}
	nomFichier := os.Args[1]
	mots, err := projethangman.LireFichier(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	mot := projethangman.MotAleatoire(mots)
	fmt.Println("Un mot a été choisi au hasard.")
	nbEssais := 8
	projethangman.Jouer(mot, nbEssais)
}
