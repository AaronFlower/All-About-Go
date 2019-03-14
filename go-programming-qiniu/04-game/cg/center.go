package cg

import (
	"encoding/json"
	"errors"
	"sync"

	mipc "github.com/AaronFlower/All-About-Go/go-programming-qiniu/04-game/ipc"
)

// CenterServer 实现了几个示范用的指令：添加用户，删除用户，列出用户和向用户广播消息。

// Message defines message struct
type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

// Room defines a room struct
type Room struct {
}

// CenterServer manages other servers in the system.
type CenterServer struct {
	servers map[string]mipc.Server
	players []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

// NewCenterServer returns a CenterServer instance.
func NewCenterServer() *CenterServer {
	servers := make(map[string]mipc.Server)
	players := make([]*Player, 0)

	return &CenterServer{
		servers: servers,
		players: players,
	}
}

// Name returns the sever name
func (server *CenterServer) Name() string {
	return "CenterServer"
}

func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	// Should to check repeated players
	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i, v := range server.players {
		if v.Name == params {
			server.players = append(server.players[:i], server.players[i+1:]...)
			return nil
		}
	}
	return errors.New("player not found")
}

func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()

	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("no player online")
	}
	return
}

func (server *CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("no player online")
	}
	return err
}

// Handle handle the request from client and send response.
func (server *CenterServer) Handle(method, params string) *mipc.Response {
	switch method {
	case "addPlayer":
		err := server.addPlayer(params)
		if err != nil {
			return &mipc.Response{Code: err.Error()}
		}
	case "removePlayer":
		err := server.removePlayer(params)
		if err != nil {
			return &mipc.Response{Code: err.Error()}
		}
	case "listPlayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &mipc.Response{Code: err.Error()}
		}
		return &mipc.Response{Code: "200", Body: players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &mipc.Response{Code: err.Error()}
		}
	default:
		return &mipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &mipc.Response{Code: "200"}
}
