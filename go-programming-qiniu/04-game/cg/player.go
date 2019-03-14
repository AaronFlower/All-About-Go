package cg

import "fmt"

// Player defines a player info in the system.
type Player struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
	Exp   int    `json:"exp"`
	Room  int    `json:"room"`

	mq chan *Message // 等待收取信息
}

// NewPlayer returns the a player instance.
// 便于演示也聊天系统，我们为每个玩家都启了一个独立的 goroutine, 监听所有发送给他们的聊天信息。
// 一旦收到就打印到控制台上。
func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, " received message: ", msg.Content)
		}
	}(player)

	return player
}
