package main

import (
	"fmt"
)

type Objet struct {
	name       string
	price      int
	damage     int
	protection int
	heal       int
}

type Boutique struct {
	Soin       []Objet
	Armes      []Objet
	Protection []Objet
}

type Atelier struct {
	Recettes []Objet
}

var inventory []Objet
var jetonsAtelier = 10
var argent = 20

var playerDamage = 5
var playerProtection = 0
var playerHP = 100

func addToInventory(item Objet) {
	if len(inventory) >= 10 {
		fmt.Println("❌ Inventaire plein ! Vous ne pouvez pas porter plus de 10 objets.")
		return
	}
	inventory = append(inventory, item)

	if item.damage > 0 {
		playerDamage += item.damage
	}
	if item.protection > 0 {
		playerProtection += item.protection
	}
	if item.heal > 0 {
		playerHP += item.heal
		if playerHP > 100 {
			playerHP = 100 //
		}
	}

	fmt.Printf("✅ %s a été ajouté à votre inventaire.\n", item.name)
}

func accessInventory() {
	if len(inventory) == 0 {
		fmt.Println("\nVotre inventaire est vide.")
	} else {
		fmt.Println("\n--- Inventaire ---")
		for i, item := range inventory {
			fmt.Printf("%d. %s | Prix: %d€ | Dégâts: %d | Protection: %d | Soin: %d\n",
				i+1, item.name, item.price, item.damage, item.protection, item.heal)
		}
	}
	fmt.Printf("\n💰 Argent restant : %d€\n", argent)
	fmt.Printf("🎟️ Jetons d'atelier restants : %d\n", jetonsAtelier)
	fmt.Printf("\n📊 Stats joueur → PV: %d | Dégâts: %d | Protection: %d\n", playerHP, playerDamage, playerProtection)
}

func acheterObjet(item Objet) {
	if argent >= item.price {
		argent -= item.price
		addToInventory(item)
		fmt.Printf("🛒 Vous avez acheté %s pour %d€. (Il vous reste %d€)\n", item.name, item.price, argent)
	} else {
		fmt.Printf("⚠️ Vous n'avez pas assez d'argent pour acheter %s. Il vous faut %d€, mais vous avez %d€.\n", item.name, item.price, argent)
	}
}

func craftItem(atelier Atelier) {
	fmt.Println("\n--- Atelier ---")
	for i, item := range atelier.Recettes {
		fmt.Printf("%d. %s (coût : %d jetons | Dégâts: %d | Protection: %d | Soin: %d)\n",
			i+1, item.name, item.price, item.damage, item.protection, item.heal)
	}
	fmt.Print("Choisissez un objet à crafter (numéro) : ")
	var choix int
	fmt.Scan(&choix)

	if choix >= 1 && choix <= len(atelier.Recettes) {
		obj := atelier.Recettes[choix-1]
		if jetonsAtelier >= obj.price {
			jetonsAtelier -= obj.price
			addToInventory(obj)
			fmt.Printf("⚒️ Vous avez crafté %s. (Il vous reste %d jetons)\n", obj.name, jetonsAtelier)
		} else {
			fmt.Printf("⚠️ Pas assez de jetons pour crafter %s. Il faut %d jetons, vous avez %d.\n", obj.name, obj.price, jetonsAtelier)
		}
	} else {
		fmt.Println("Choix invalide.")
	}
}

func main() {
	listItems := Boutique{
		Soin: []Objet{
			{"Donuts", 3, 0, 0, 10},
			{"Redbull", 5, 0, 0, 20},
		},
		Armes: []Objet{
			{"Matraque renforcée", 15, 5, 0, 0},
			{"Pistolet de service", 30, 10, 0, 0},
			{"Fusil à pompe", 50, 20, 0, 0},
			{"Fusil d’assaut", 75, 30, 0, 0},
			{"Sniper", 100, 50, 0, 0},
		},
		Protection: []Objet{
			{"Casque de patrouille", 10, 0, 5, 0},
			{"Gilet pare-balles", 30, 0, 15, 0},
			{"Plaque renforcée", 35, 0, 20, 0},
		},
	}

	atelier := Atelier{
		Recettes: []Objet{
			{"Grenade fumigène", 10, 0, 0, 20},
			{"Grenade flash", 10, 20, 0, 0},
			{"Grenade", 20, 50, 0, 0},
		},
	}

	fmt.Println("Bienvenue dans la boutique et l’atelier !\n")
	fmt.Printf("💰 Vous commencez avec %d€ et %d jetons d’atelier.\n", argent, jetonsAtelier)
	fmt.Printf("📊 Stats de départ → PV: %d | Dégâts: %d | Protection: %d\n", playerHP, playerDamage, playerProtection)

	for {
		fmt.Println("\nAppuyez sur :")
		fmt.Println("1 = Soins")
		fmt.Println("2 = Armes")
		fmt.Println("3 = Protection")
		fmt.Println("4 = Voir inventaire")
		fmt.Println("5 = Atelier (crafter avec jetons)")
		fmt.Println("0 = Quitter")

		var choix int
		fmt.Print("Votre choix : ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			fmt.Println("\n--- Soins ---")
			for i, item := range listItems.Soin {
				fmt.Printf("%d. %s : %d€ (+%d PV)\n", i+1, item.name, item.price, item.heal)
			}
			fmt.Print("Choisissez un objet à acheter (numéro) : ")
			var subChoice int
			fmt.Scan(&subChoice)
			if subChoice >= 1 && subChoice <= len(listItems.Soin) {
				acheterObjet(listItems.Soin[subChoice-1])
			}

		case 2:
			fmt.Println("\n--- Armes ---")
			for i, item := range listItems.Armes {
				fmt.Printf("%d. %s : %d€ (+%d dégâts)\n", i+1, item.name, item.price, item.damage)
			}
			fmt.Print("Choisissez un objet à acheter (numéro) : ")
			var subChoice int
			fmt.Scan(&subChoice)
			if subChoice >= 1 && subChoice <= len(listItems.Armes) {
				acheterObjet(listItems.Armes[subChoice-1])
			}

		case 3:
			fmt.Println("\n--- Protection ---")
			for i, item := range listItems.Protection {
				fmt.Printf("%d. %s : %d€ (+%d protection)\n", i+1, item.name, item.price, item.protection)
			}
			fmt.Print("Choisissez un objet à acheter (numéro) : ")
			var subChoice int
			fmt.Scan(&subChoice)
			if subChoice >= 1 && subChoice <= len(listItems.Protection) {
				acheterObjet(listItems.Protection[subChoice-1])
			}

		case 4:
			accessInventory()

		case 5:
			craftItem(atelier)

		case 0:
			fmt.Println("Merci d'avoir visité la boutique et l’atelier. À bientôt !")
			return

		default:
			fmt.Println("Choix invalide, veuillez entrer un chiffre entre 0 et 5.")
		}
	}
}
