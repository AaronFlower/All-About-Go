package mp

import "fmt"

// Play plays a music by source name and type
func Play(name, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		fallthrough
	case "mp3":
		p = &MP3Player{}
	case "wav":
		fallthrough
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("Unsupported music type", mtype)
		return
	}

	p.Play(name)
}
