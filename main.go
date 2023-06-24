package main

import (
	"encoding/json"
	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

func main() {
	for {
		_ = connect()
		time.Sleep(time.Second * 5)
	}
}

func connect() error {
	host := "101.42.239.41:10042"
	u := url.URL{Scheme: "ws", Host: host, Path: "/ws/receive"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer c.Close()

	log.Printf("已经连接到%s", host)

	for {
		// read message
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败, err=%s", err)
			return err
		}
		log.Printf("received message: %s", message)

		resp := make(map[string]string)
		_ = json.Unmarshal(message, &resp)
		err = beeep.Notify("†✰⇝->♡†", resp["message"], "icon.jpeg")
		if err != nil {
			log.Printf("发送系统通知失败, err=" + err.Error())
		}
	}
}
