package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

var host *string
var port *int

func main() {
	host = flag.String("h", "127.0.0.1", "输入域名或ip, 默认127.0.0.1")
	port = flag.Int("p", 10041, "输入端口, 默认10041")
	flag.Parse()
	for {
		_ = connect()
		time.Sleep(time.Second * 5)
	}
}

func connect() error {
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("app-recover.")
		}
	}()

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%v", *host, *port), Path: "/ws/receive"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	defer c.Close()

	log.Printf("已经连接到%s", *host)

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
