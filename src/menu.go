package game

import (
	"fmt"
)

const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	Yellow     = "\033[33m"
	Blue       = "\033[34m"
	Cyan       = "\033[36m"
	BrightBlue = "\033[94m"
)

func Pause() {
	fmt.Println(BrightBlue + "Appuyez sur Entrée pour continuer" + Reset)
	fmt.Scanln()
	fmt.Scanln()
}

func Menu() {
	for {
		fmt.Println(Yellow + "\n===== MENU PRINCIPAL =====" + Reset)
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l’inventaire")
		fmt.Println("3 - Boutique")
		fmt.Println("4 - Atelier")
		fmt.Println("5 - Combat")
		fmt.Println("6 - Quitter")
		fmt.Printf(Green+"💰 Argent : %d€ | 🎟️ Jetons d'atelier : %d\n"+Reset, argent, jetonsAtelier)

		var choixMenu int
		fmt.Print(BrightBlue + "Faites un choix : " + Reset)
		fmt.Scan(&choixMenu)

		switch choixMenu {
		case 1:
			perso := getCurrentPerso()
			if perso != nil {
				fmt.Println(Cyan + "\n===== INFOS PERSONNAGE =====" + Reset)
				fmt.Println("Nom       :", perso.Name)
				fmt.Println("Classe    :", perso.Classe)
				fmt.Println("Niveau    :", perso.Level)
				fmt.Printf("HP        : %d / %d\n", perso.HP, perso.MaxHP)
				fmt.Println("Dégâts    :", perso.Damage)
				fmt.Println("Protection:", perso.Protection)
				fmt.Println("Équipement:", perso.Stuff)
			} else {
				fmt.Println(Red + "Aucun personnage sélectionné." + Reset)
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
			fmt.Println(Cyan + "1 - Descente dans un Quartier (gain jeton)" + Reset)
			fmt.Println(Cyan + "2 - Bavure (gain argent)" + Reset)
			fmt.Println(Red + "0 - Quitter" + Reset)
			fmt.Print(BrightBlue + "Choisissez : " + Reset)
			fmt.Scan(&choixCombat)

			switch choixCombat {
			case 1:
				QuartierCombat()
			case 2:
				Bavure()
			case 0:
				fmt.Println("Vous avez quitté")
			default:
				fmt.Println(Red + "Choix invalide." + Reset)
			}
			Pause()

		case 6:
			fmt.Println(Green + "Fermeture du jeu. À bientôt !" + Reset)
			return

		default:
			fmt.Println(Red + "Choix invalide." + Reset)
			Pause()
		}
	}
}
