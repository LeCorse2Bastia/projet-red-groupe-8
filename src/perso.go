package game

import (
	"fmt"
	"time"
)

type Perso struct {
	Name       string
	Classe     string
	Level      int
	MaxHP      int
	HP         int
	Damage     int
	Protection int
	Stuff      string
}

var (
	choixPerso int
	Money      int
)

func PrintWelcome() {
	msg1 := `|       \ |      \|        \|  \  |  \|  \   |  \|        \|  \  |  \|  \  |  \|        \       /      \ |  \  |  \|       \       
| $$$$$$$\ \$$$$$$| $$$$$$$$| $$\ | $$| $$   | $$| $$$$$$$$| $$\ | $$| $$  | $$| $$$$$$$$      |  $$$$$$\| $$  | $$| $$$$$$$\      
| $$__/ $$  | $$  | $$__    | $$$\| $$| $$   | $$| $$__    | $$$\| $$| $$  | $$| $$__          | $$___\$$| $$  | $$| $$__| $$      
| $$    $$  | $$  | $$  \   | $$$$\ $$ \$$\ /  $$| $$  \   | $$$$\ $$| $$  | $$| $$  \          \$$    \ | $$  | $$| $$    $$      
| $$$$$$$\  | $$  | $$$$$   | $$\$$ $$  \$$\  $$ | $$$$$   | $$\$$ $$| $$  | $$| $$$$$          _\$$$$$$\| $$  | $$| $$$$$$$\      
| $$__/ $$ _| $$_ | $$_____ | $$ \$$$$   \$$ $$  | $$_____ | $$ \$$$$| $$__/ $$| $$_____       |  \__| $$| $$__/ $$| $$  | $$      
| $$    $$|   $$ \| $$     \| $$  \$$$    \$$$   | $$     \| $$  \$$$ \$$    $$| $$     \       \$$    $$ \$$    $$| $$  | $$      
 \$$$$$$$  \$$$$$$ \$$$$$$$$ \$$   \$$     \$     \$$$$$$$$ \$$   \$$  \$$$$$$  \$$$$$$$$        \$$$$$$   \$$$$$$  \$$   \$$      `

	msg2 := ` /      \  /      \ |  \  |  \ /      \        /      \  /      \ |  \  |  \|        \|       \  /      \ |  \                     
|  $$$$$$\|  $$$$$$\| $$  | $$|  $$$$$$\      |  $$$$$$\|  $$$$$$\| $$\ | $$ \$$$$$$$$| $$$$$$$\|  $$$$$$\| $$                     
| $$___\$$| $$  | $$| $$  | $$| $$___\$$      | $$   \$$| $$  | $$| $$$\| $$   | $$   | $$__| $$| $$  | $$| $$                     
 \$$    \ | $$  | $$| $$  | $$ \$$    \       | $$      | $$  | $$| $$$$\ $$   | $$   | $$    $$| $$  | $$| $$                     
 _\$$$$$$\| $$  | $$| $$  | $$ _\$$$$$$\      | $$   __ | $$  | $$| $$\$$ $$   | $$   | $$$$$$$\| $$  | $$| $$                     
|  \__| $$| $$__/ $$| $$__/ $$|  \__| $$      | $$__/  \| $$__/ $$| $$ \$$$$   | $$   | $$  | $$| $$__/ $$| $$_____                
 \$$    $$ \$$    $$ \$$    $$ \$$    $$       \$$    $$ \$$    $$| $$  \$$$   | $$   | $$  | $$ \$$    $$| $$     \               
  \$$$$$$   \$$$$$$   \$$$$$$   \$$$$$$         \$$$$$$   \$$$$$$  \$$   \$$    \$$    \$$   \$$  \$$$$$$  \$$$$$$$$               `
	time.Sleep(600 * time.Millisecond)
	printSlowly(msg1, 3*time.Millisecond)
	printSlowly(msg2, 3*time.Millisecond)
}

var perso1 = Perso{"Marc", "Brigade Anti-Émeute", 1, 100, 100, 10, 0, "Matraque / 1 Donut"}
var perso2 = Perso{"Cyril", "Opérateur Précis", 1, 45, 45, 30, 0, "Matraque / 1 Donut"}
var perso3 = Perso{"Bastien", "Infiltré", 1, 65, 65, 20, 0, "Matraque / 1 Donut"}

func getCurrentPerso() *Perso {
	switch choixPerso {
	case 1:
		return &perso1
	case 2:
		return &perso2
	case 3:
		return &perso3
	default:
		return nil
	}
}

func Initcharacter() {
	PrintWelcome()

	playSoundAsync("./../attente.mp3")

	HistPerso1 := "Toujours en première ligne, Marc avance lentement, bouclier levé, prêt à encaisser les coups pour protéger ses frères d’armes.\n Son armure le rend presque inébranlable, et sa force tranquille impose le respect.\n Il n’est pas le plus rapide, mais quand il frappe, ça compte. Chaque mission est pour lui un devoir, pas un combat.\n"
	HistPerso2 := "Ancien militaire d’élite, surnommé “le Chirurgien”, Cyril ne frappe jamais par hasard.\n Il observe, attend, et place une attaque décisive qui change le cours du combat.\n Froid et méthodique, il maîtrise l'art de neutraliser sans gaspiller d'énergie.\n Un seul tir, une seule chance, aucun échec permis.\n"
	HistPerso3 := "Cinq années passées sous couverture ont fait de Bastien une ombre parmi les hommes.\n Il se glisse partout, entend tout, et frappe là où l’ennemi se croit en sécurité.\n Sa discrétion est son arme la plus redoutable, bien plus que sa matraque.\n Bastien agit sans être vu, mais ses résultats parlent pour lui.\n"

	for {
		fmt.Println("\nChoisissez votre personnage :")
		fmt.Println("1 -", perso1.Name, ":", perso1.Classe, "(", perso1.HP, "PV )")
		printSlowly(HistPerso1, 15*time.Millisecond)
		fmt.Println("2 -", perso2.Name, ":", perso2.Classe, "(", perso2.HP, "PV )")
		printSlowly(HistPerso2, 15*time.Millisecond)
		fmt.Println("3 -", perso3.Name, ":", perso3.Classe, "(", perso3.HP, "PV )")
		printSlowly(HistPerso3, 15*time.Millisecond)

		fmt.Print(Cyan + "Appuyez sur 1, 2 ou 3 pour choisir votre personnage : " + Reset)
		fmt.Scan(&choixPerso)

		if choixPerso >= 1 && choixPerso <= 3 {
			stopSound()

			switch choixPerso {
			case 1:
				fmt.Println(Green+"Vous avez choisi :", perso1.Name, Reset)
			case 2:
				fmt.Println(Green+"Vous avez choisi :", perso2.Name, Reset)
			case 3:
				fmt.Println(Green+"Vous avez choisi :", perso3.Name, Reset)
			}
			break
		} else {
			fmt.Println(Red + "Choix invalide. Veuillez recommencer." + Reset)
		}
	}
}
