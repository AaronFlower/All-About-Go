package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AaronFlower/All-About-Go/go-programming-qiniu/04-game/cg"
	mipc "github.com/AaronFlower/All-About-Go/go-programming-qiniu/04-game/ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := mipc.NewIPCServer(&cg.CenterServer{})
	client := mipc.NewIPCClient(server)
	centerClient = &cg.CenterClient{IPCClient: client}
	return nil
}

func help(args []string) int {
	fmt.Println(`
		Commands:
			login <username> <level> <exp>
			logout <username>
			send <message>
			listplayer
			quit(q)
			help(h)
	`)
	return 0
}

func quit(args []string) int {
	return 1
}

func logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username> <level> <exp>")
		return 0
	}

	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Inavlid Parameter: <level> should be an integer.")
	}

	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Inavlid Parameter: <exp> should be an integer.")
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp

	err = centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("Failed adding player", err)
	}
	return 0
}

func listPlayer(args []string) int {
	ps, err := centerClient.ListPlayer("")

	if err != nil {
		fmt.Println("Failed. ", err)
	} else {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}

func send(args []string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)

	if err != nil {
		fmt.Println("Failed!", err)
	}
	return 0
}

func getCommandHandlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"help":       help,
		"h":          help,
		"quit":       quit,
		"q":          quit,
		"login":      login,
		"logout":     logout,
		"listplayer": listPlayer,
		"send":       send,
	}
}

func main() {
	fmt.Println("Causal Game Server Solution")

	startCenterService()

	help(nil)

	handlers := getCommandHandlers()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Command >")

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			} else {
				fmt.Print("Command > ")
			}
		} else {
			fmt.Println("Unkonw command.", tokens[0])
			fmt.Print("Command > ")
		}
	}
}
