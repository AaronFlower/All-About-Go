package mp

import (
	"fmt"
	"time"
)

// MP3Player defines a mp3 player
type MP3Player struct {
	stat     int
	progress int
}

// Play plays the mp3 type music
func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music ", source)

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\n Finished playing.", source)
}
