package projethangman

import (
	"fmt"
	"strings"
)

func Jouer(mot string, nbEssais int, nbLettresDévoilées int, fichierPendu string, niveau string) {
	mot = strings.ToLower(mot)
	lettresDevinees := RevelerLettres(mot, nbLettresDévoilées)
	essaisRestants := nbEssais
	lettresProposees := map[rune]bool{}

	for essaisRestants > 0 {
		ClearTerminal()
	
		AfficherEtat(lettresDevinees, essaisRestants)

		var proposition string
		fmt.Print("Proposez une lettre ou un mot : ")
		fmt.Scanln(&proposition)

		if len(proposition) == 0 {
			fmt.Println("Vous devez entrer une lettre ou un mot.")
			continue
		}

		proposition = strings.ReplaceAll(proposition, " ", "")

		proposition = strings.ToLower(proposition)

		if len(proposition) > 1 { 
			if proposition == mot {
				fmt.Println("Félicitations, vous avez deviné le mot !")
				return
			} else {
				fmt.Println("Mauvais mot ! Vous perdez 2 essais.")
				essaisRestants -= 2
			}
		} else { 
			lettre := rune(proposition[0])

			if lettresProposees[lettre] {
				fmt.Println("Vous avez déjà proposé cette lettre.")
				continue
			}
			lettresProposees[lettre] = true

			if strings.ContainsRune(mot, lettre) {
				fmt.Println("Bonne lettre !")
				for i, r := range mot {
					if r == lettre {
						lettresDevinees[i] = lettre
					}
				}
			} else {
				fmt.Println("Mauvaise lettre !")
				essaisRestants--
			}
		}

		if string(lettresDevinees) == mot {
			fmt.Println("Félicitations, vous avez deviné le mot !")
			return
		}
	}

	fmt.Printf("Désolé, vous avez perdu. Le mot était : %s\n", mot)
}
