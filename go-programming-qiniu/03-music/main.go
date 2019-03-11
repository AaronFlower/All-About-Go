package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mlib "github.com/AaronFlower/All-About-Go/go-programming-qiniu/03-music/manager"
	mp "github.com/AaronFlower/All-About-Go/go-programming-qiniu/03-music/player"
)

var lib *mlib.MusicManager
var id = 1

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.GetByName(tokens[1])
	if e == nil {
		fmt.Println("The music ", tokens[1], "does not exist.")
		return
	}

	mp.Play(e.Name, e.Type)
}

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.GetByIndex(i)
			fmt.Println(i+1, " : ", e.Name, e.Artist, e.Type)
		}
	case "add":
		if len(tokens) == 5 {
			id++
			lib.Add(&mlib.MusicEntry{
				ID:     strconv.Itoa(id),
				Name:   tokens[2],
				Artist: tokens[3],
				Type:   tokens[4],
			})
		} else {
			fmt.Println("USAGE: lib add <name> <artist> <type>")
		}
	case "remove":
		if len(tokens) == 3 {
			i, err := strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println("USAGE: lib remove <index>")
			}
			lib.RemoveByIndex(i)
		} else {
			fmt.Println("USAGE: lib remove <index>")
		}
	default:
		fmt.Println("Unknown command: ", tokens[1])
	}
}

func main() {
	fmt.Println(`
		Enter following commands to control the player:
		lib list 			-- View the existing music lib
		lib add <name><artist><type> -- Add a music to the music lib
		lib remove <index> 	-- Remove the specified music from the lib
		play <name> 		-- play the specified music
	`)

	lib = mlib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command -> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unknown command: ", tokens[0])
		}
	}
}
