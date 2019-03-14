package cg

import (
	"encoding/json"
	"errors"

	mipc "github.com/AaronFlower/All-About-Go/go-programming-qiniu/04-game/ipc"
)

// CenterClient defines a CenterClient struct
type CenterClient struct {
	*mipc.IPCClient
}

// AddPlayer add a player
func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)

	if err != nil {
		return err
	}

	resp, err := client.Call("addPlayer", string(b))

	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

// RemovePlayer removes a player
func (client *CenterClient) RemovePlayer(name string) error {
	ret, _ := client.Call("removePlayer", name)

	if ret.Code == "200" {
		return nil
	}
	return errors.New(ret.Code)
}

// ListPlayer lists all the players in the server
func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, err := client.Call("listPlayer", params)

	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

// Broadcast broadcasts message to all players in the server.
func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message}

	b, err := json.Marshal(m)

	if err != nil {
		return err
	}

	resp, _ := client.Call("broadcast", string(b))

	if resp.Code == "200" {
		return nil
	}

	return errors.New(resp.Code)
}
