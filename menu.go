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
		fmt.Println("\nMenu :")
		fmt.Println("1 - Afficher les informations du personnage")
		fmt.Println("2 - Accéder au contenu de l’inventaire")
		fmt.Println("3 - Boutique")
		fmt.Println("4 - Atelier")
		fmt.Println("5 - Combat")
		fmt.Println("6 - Quitter")
		fmt.Println(Money, "$")

		var choixMenu int
		fmt.Print("Faites un choix : ")
		fmt.Scan(&choixMenu)

		switch choixMenu {
		case 1:
			switch choixPerso {
			case 1:
				fmt.Println("Nom     :", perso1.Name)
				fmt.Println("Classe  :", perso1.Classe)
				fmt.Println("Niveau  :", perso1.Level)
				fmt.Println("HP      :", perso1.HP, "/", perso1.MaxHP)
				fmt.Println("Dégâts  :", perso1.Damage)
			case 2:
				fmt.Println("Nom     :", perso2.Name)
				fmt.Println("Classe  :", perso2.Classe)
				fmt.Println("Niveau  :", perso2.Level)
				fmt.Println("HP      :", perso2.HP, "/", perso2.MaxHP)
				fmt.Println("Dégâts  :", perso2.Damage)
			case 3:
				fmt.Println("Nom     :", perso3.Name)
				fmt.Println("Classe  :", perso3.Classe)
				fmt.Println("Niveau  :", perso3.Level)
				fmt.Println("HP      :", perso3.HP, "/", perso3.MaxHP)
				fmt.Println("Dégâts  :", perso3.Damage)
			default:
				fmt.Println("Choix invalide.")
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
			fmt.Println("1 - Descente Quartier")
			fmt.Println("2 - Bavure")
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
			fmt.Println("Fermeture du jeu.")
			return
		default:
			fmt.Println("Choix invalide.")
			Pause()
		}
	}
}
