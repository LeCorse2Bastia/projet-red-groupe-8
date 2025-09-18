package game

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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
	choixPerso   int
	Money        int
	streamer     beep.StreamSeekCloser
	soundControl *os.File
	format       beep.Format
)

func playSoundAsync2() {
	var err error
	soundControl, err = os.Open("attente.mp3")
	if err != nil {
		fmt.Println("Erreur ouverture du fichier audio :", err)
		return
	}

	streamer, format, err = mp3.Decode(soundControl)
	if err != nil {
		fmt.Println("Erreur décodage mp3 :", err)
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Millisecond*50))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		streamer.Close()
		soundControl.Close()
	})))
}

func stopSound() {
	if streamer != nil {
		streamer.Close()
	}
	if soundControl != nil {
		soundControl.Close()
	}
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
	playSoundAsync2()

	HistPerso1 := "Histoire: Toujours en première ligne, il avance lentement, bouclier levé, prêt à encaisser..."
	HistPerso2 := "Histoire: Ancien militaire surnommé le chirurgien, il frappe rarement, mais toujours au bon endroit..."
	HistPerso3 := "Histoire: Cinq ans sous couverture ont fait de lui un prédateur discret..."

	fmt.Println("Choisissez votre personnage :")
	fmt.Println("1 -", perso1.Name, ":", perso1.Classe, "(", perso1.HP, "PV )")
	fmt.Println(HistPerso1)
	fmt.Println("2 -", perso2.Name, ":", perso2.Classe, "(", perso2.HP, "PV )")
	fmt.Println(HistPerso2)
	fmt.Println("3 -", perso3.Name, ":", perso3.Classe, "(", perso3.HP, "PV )")
	fmt.Println(HistPerso3)

	fmt.Print("Appuyez sur 1, 2 ou 3 pour choisir votre personnage : ")
	fmt.Scan(&choixPerso)
	stopSound()

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
