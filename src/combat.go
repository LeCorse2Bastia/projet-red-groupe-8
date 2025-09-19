package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Quartier struct {
	Nom         string
	PV          int
	Difficulte  int
	DegatsFixes int
}

var quartiers = map[string]*Quartier{
	"Le Panier":       {"Le Panier", 30, 2, 15},
	"La Marine Bleue": {"La Marine Bleue", 50, 4, 30},
	"Font Vert":       {"Font Vert", 70, 5, 50},
	"La Castellane":   {"La Castellane", 100, 6, 80},
}

func printSlowly(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

func Bavure() {
	playSoundAsync("./../son.mp3")

	asciiArt := `        .---.
      /_____\
      ( '.' )
       \_-_/_ 
    .-"` + "`" + `'V'//-.
   / ,   |// , \
  / /|Ll //Ll|\ \
 / / |__//   | \_\
 \ \/---|[]==| / /
  \/\__/ |   \/\/
   |/_   | Ll_\|
     |' ^"""^'|
     |   |   |
     |   |   |
     |   |   |
     |   |   |
     L___l___J
      |_ | _|
     (___|___)
      ^^^ ^^^ `

	printSlowly(asciiArt, 15*time.Millisecond)
	time.Sleep(5 * time.Second)
	argent += 10
}

func QuartierCombat() {
	perso := getCurrentPerso()
	if perso == nil {
		fmt.Println("Aucun personnage sélectionné.")
		return
	}

	var choix string
	fmt.Println("Choisis un quartier à descendre :")
	fmt.Println("1 - Le Panier")
	fmt.Println("2 - La Marine Bleue")
	fmt.Println("3 - Font Vert")
	fmt.Println("4 - La Castellane")
	fmt.Print("Entrée (1-4) : ")
	fmt.Scan(&choix)

	var quartier *Quartier
	switch choix {
	case "1":
		quartier = quartiers["Le Panier"]
	case "2":
		quartier = quartiers["La Marine Bleue"]
	case "3":
		quartier = quartiers["Font Vert"]
	case "4":
		quartier = quartiers["La Castellane"]
	default:
		fmt.Println("Choix invalide.")
		Pause()
		return
	}

	if quartier.PV <= 0 {
		fmt.Println("Ce quartier a déjà été nettoyé.")
		Pause()
		return
	}

	fmt.Printf("💥 Descente dans %s !\n", quartier.Nom)

	for quartier.PV > 0 && perso.HP > 0 {
		fmt.Println("\n--- Tour de combat ---")
		fmt.Printf("%s - PV: %d\n", quartier.Nom, quartier.PV)
		fmt.Printf("%s - PV: %d/%d | Dégâts: %d | Protection: %d\n", perso.Name, perso.HP, perso.MaxHP, perso.Damage, perso.Protection)

		fmt.Println("1 - Attaquer")
		fmt.Println("2 - Passer le tour")
		fmt.Print("Choix : ")
		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			degatsPerso := rand.Intn(10) + perso.Damage
			quartier.PV -= degatsPerso
			if quartier.PV < 0 {
				quartier.PV = 0
			}
			fmt.Printf("🗡️ Tu infliges %d dégâts à %s.\n", degatsPerso, quartier.Nom)

		case 2:
			fmt.Println("⏩ Tu passes ton tour.")

		default:
			fmt.Println("Choix invalide. Tu perds ton tour.")
		}

		if quartier.PV > 0 {
			degatsQuartier := quartier.DegatsFixes - perso.Protection
			if degatsQuartier < 0 {
				degatsQuartier = 0
			}
			perso.HP -= degatsQuartier
			if perso.HP < 0 {
				perso.HP = 0
			}
			fmt.Printf("💢 %s t’inflige %d dégâts.\n", quartier.Nom, degatsQuartier)
		}
	}

	if quartier.PV <= 0 && perso.HP > 0 {
		fmt.Printf("✅ Tu as nettoyé %s avec succès !\n", quartier.Nom)
		Money += 40
		jetonsAtelier += 10
		fmt.Println("🎟️ +10 jetons d’atelier gagnés pour cette victoire !")
	} else if perso.HP <= 0 {
		playSoundAsync("./../rire.mp3")
		fmt.Println("❌ Tu es tombé au combat...")
	}

	if quartiers["Le Panier"].PV <= 0 &&
		quartiers["La Marine Bleue"].PV <= 0 &&
		quartiers["Font Vert"].PV <= 0 &&
		quartiers["La Castellane"].PV <= 0 {
		fmt.Println("🏆 Félicitations, tous les quartiers ont été nettoyés !")
		Money += 100
	}

	Pause()
}

func ResetQuartiers() {
	quartiers["Le Panier"].PV = 200
	quartiers["La Marine Bleue"].PV = 250
	quartiers["Font Vert"].PV = 300
	quartiers["La Castellane"].PV = 400
}
