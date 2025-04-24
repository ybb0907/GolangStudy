package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 升级 HTTP 连接到 WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		log.Printf("客户端已连接: %s", conn.RemoteAddr())

		// 持续监听客户端消息
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("客户端断开: %s (%v)", conn.RemoteAddr(), err)
				return
			}
			log.Printf("收到消息: %s", string(message))
		}
	})

	log.Println("服务端启动，监听 :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
