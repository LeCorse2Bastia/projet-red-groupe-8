package game

import (
	"fmt"
)

func Pause() {
	fmt.Println("Appuyez sur Entrée pour continuer")
	fmt.Scanln()
	fmt.Scanln()
}

func Menu() {
	Money = 20
	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l’inventaire")
		fmt.Println("3 - Boutique")
		fmt.Println("4 - Atelier")
		fmt.Println("5 - Combat")
		fmt.Println("6 - Quitter")
		fmt.Printf("💰 Argent : %d€ | 🎟️ Jetons d'atelier : %d\n", argent, jetonsAtelier)

		var choixMenu int
		fmt.Print("Faites un choix : ")
		fmt.Scan(&choixMenu)

		switch choixMenu {
		case 1:
			perso := getCurrentPerso()
			if perso != nil {
				fmt.Println("\n===== INFOS PERSONNAGE =====")
				fmt.Println("Nom       :", perso.Name)
				fmt.Println("Classe    :", perso.Classe)
				fmt.Println("Niveau    :", perso.Level)
				fmt.Printf("HP        : %d / %d\n", perso.HP, perso.MaxHP)
				fmt.Println("Dégâts    :", perso.Damage)
				fmt.Println("Protection:", perso.Protection)
				fmt.Println("Équipement:", perso.Stuff)
			} else {
				fmt.Println("Aucun personnage sélectionné.")
			}
			Pause()

		case 2:
			accessInventory()
			Pause()

		case 3:
			Shop()
			Pause()

		case 4:
			craftItem(Atelier{})
			Pause()

		case 5:
			var choixCombat int
			fmt.Println("1 - Descente dans un Quartier")
			fmt.Println("2 - Bavure (bonus)")
			fmt.Print("Choisissez : ")
			fmt.Scan(&choixCombat)

			switch choixCombat {
			case 1:
				QuartierCombat()
			case 2:
				Bavure()
				Money += 20
			default:
				fmt.Println("Choix invalide.")
			}
			Pause()

		case 6:
			fmt.Println("Fermeture du jeu. À bientôt !")
			return

		default:
			fmt.Println("Choix invalide.")
			Pause()
		}
	}
}
