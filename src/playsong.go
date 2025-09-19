package game

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	streamer     beep.StreamSeekCloser
	soundControl *os.File
	format       beep.Format
)

func playSoundAsync(pathSong string) {
	var err error
	soundControl, err = os.Open(pathSong)
	if err != nil {
		fmt.Println(Red+"Erreur ouverture du fichier audio :"+Reset, err)
		return
	}

	streamer, format, err = mp3.Decode(soundControl)
	if err != nil {
		fmt.Println(Red+"Erreur décodage mp3 :"+Reset, err)
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
