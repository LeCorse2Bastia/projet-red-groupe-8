package game

import "fmt"

type Perso struct {
	Name   string
	Classe string
	Level  int
	MaxHP  int
	HP     int
	Damage int
	Stuff  string
}

var choixPerso int
var Money int

var perso1 = Perso{"Marc", "Brigade Anti-Émeute", 1, 200, 200, 10, "Matraque / 1 Donut"}
var perso2 = Perso{"Cyril", "Opérateur Précis", 1, 90, 90, 30, "Matraque / 1 Donut"}
var perso3 = Perso{"Bastien", "Infiltré", 1, 130, 130, 20, "Matraque / 1 Donut"}

func Initcharacter() {
	HistPerso1 := "Histoire: Toujours en première ligne, il avance lentement, bouclier levé, prêt à encaisser. Silencieux et massif, il ne frappe pas fort, mais il ne tombe jamais. Sa présence seule suffit souvent à calmer les plus téméraires.\n"
	HistPerso2 := "Histoire: Ancien militaire surnommé le chirurgien, il frappe rarement, mais toujours au bon endroit. Chaque geste est calculé, chaque coup vise à faire tomber net. Calme, froid, redoutablement efficace.\n"
	HistPerso3 := "Histoire: Cinq ans sous couverture ont fait de lui un prédateur discret. Il frappe vite, fort, sans prévenir. Ni vraiment flic, ni criminel, il connaît les règles de la rue et surtout comment les briser.\n"

	fmt.Println("Choisissez votre personnage :")
	fmt.Println("1 -", perso1.Name, ":", perso1.Classe, "(", perso1.HP, "PV )")
	fmt.Println(HistPerso1)
	fmt.Println("2 -", perso2.Name, ":", perso2.Classe, "(", perso2.HP, "PV )")
	fmt.Println(HistPerso2)
	fmt.Println("3 -", perso3.Name, ":", perso3.Classe, "(", perso3.HP, "PV )")
	fmt.Println(HistPerso3)

	fmt.Print("Appuyez sur 1, 2 ou 3 pour choisir votre personnage : ")
	fmt.Scan(&choixPerso)

	switch choixPerso {
	case 1:
		fmt.Println("Vous avez choisi :", perso1.Name)
	case 2:
		fmt.Println("Vous avez choisi :", perso2.Name)
	case 3:
		fmt.Println("Vous avez choisi :", perso3.Name)
	default:
		fmt.Println("Choix invalide.")
	}
}
