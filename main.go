package main

import (
	"encoding/json"
	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

func main() {
	host := ""
	u := url.URL{Scheme: "ws", Host: host, Path: "/ws/receive"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	log.Printf("已经连接到%s", host)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败, err=%s", err)
			return
		}
		log.Printf("received message: %s", message)

		resp := make(map[string]string)
		_ = json.Unmarshal(message, &resp)
		err = beeep.Notify("短信", resp["message"], "icon.jpeg")
		if err != nil {
			log.Printf("发送系统通知失败, err=" + err.Error())
		}
	}
}
