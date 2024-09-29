package main

import (
	"fmt"
	"os"
	"projethangman"
)

func main() {
	var choix int 
	for {
		projethangman.Menu()
		fmt.Scan(&choix)
		switch choix {
		case 1:
			projethangman.ClearTerminal()
			fmt.Println("Sélectionnez le niveau de difficulté :")
			fmt.Println("1. Facile")
			fmt.Println("2. Intermédiaire")
			fmt.Println("3. Difficile")
			var niveau int
			fmt.Scan(&niveau)

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

			var nbEssais int
			var nbLettresDévoilées int
			switch niveau {
			case 1: 
				nbEssais = 10
				nbLettresDévoilées = len(mot) / 2
			case 2: 
				nbEssais = 8
				nbLettresDévoilées = len(mot) / 3
			case 3: 
				nbEssais = 6
				nbLettresDévoilées = len(mot) / 4
			default:
				fmt.Println("Niveau invalide. Sélection par défaut : Intermédiaire.")
				nbEssais = 8
				nbLettresDévoilées = len(mot) / 3
			}

			projethangman.Jouer(mot, nbEssais, nbLettresDévoilées)

		case 2:
			projethangman.ClearTerminal()
			projethangman.AfficherRegles()
		case 0:
			return
		}
	}
}
