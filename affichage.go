package projethangman

import (
	"fmt"
	"math/rand"
)

func RevelerLettres(mot string, nbLettres int) []rune {
	lettresDevinees := make([]rune, len(mot))
	for i := range lettresDevinees {
		lettresDevinees[i] = '_'
	}
	indicesDejaReveles := map[int]bool{}
	for i := 0; i < nbLettres; i++ {
		var index int
		for {
			index = rand.Intn(len(mot))
			if !indicesDejaReveles[index] {
				lettre := rune(mot[index])
				for j, c := range mot {
					if rune(c) == lettre {
						lettresDevinees[j] = lettre
						indicesDejaReveles[j] = true
					}
				}
				break
			}
		}
	}
	return lettresDevinees
}
func AfficherEtat(lettresDevinees []rune, essaisRestants int) {
	fmt.Printf("Mot à deviner : %s\n", string(lettresDevinees))
	fmt.Printf("Essais restants : %d\n", essaisRestants)
}

func Menu() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         Jeu Hangman                              ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║             1. Jouer                                             ║")
	fmt.Println("║             2. Règles                                            ║")
	fmt.Println("║             0. Quitter                                           ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}

func AfficherRegles() {
	fmt.Println("╔══════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                          Règles du jeu                           ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════════╣")
	fmt.Println("║  1. Devinez le mot en proposant des lettres une par une.         ║")
	fmt.Println("║  2. Vous pouvez également essayer de deviner le mot complet.     ║")
	fmt.Println("║  3. Selon le niveau de difficulté, vous avez un certain nombre   ║")
	fmt.Println("║     d'essais et de lettres dévoilées.                            ║")
	fmt.Println("║  4. Le jeu se termine lorsque vous devinez le mot ou lorsque     ║")
	fmt.Println("║     vous avez épuisé tous vos essais.                            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════╝")
}
