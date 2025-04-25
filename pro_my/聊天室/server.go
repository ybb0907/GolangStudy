package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true }, // 允许跨域
}

var clients = make(map[*websocket.Conn]bool) // 存储所有连接的客户端

func closeAllClients() {
	for conn := range clients {
		if clients[conn] {
			conn.Close()
			delete(clients, conn)
		}
	}
}

func sendMessageToAll(conn *websocket.Conn, message string) {
	for client := range clients {
		if client != conn && clients[client] {
			err := client.WriteMessage(websocket.TextMessage, []byte("收到来自于"+conn.RemoteAddr().String()+"的消息:"+string(message)))
			if err != nil {
				log.Printf("发送消息失败: %v", err)
				client.Close()
				clients[client] = false
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 升级 HTTP 连接到 WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}
		clients[conn] = true
		log.Printf("客户端已连接: %s", conn.RemoteAddr())
		defer conn.Close()
		// 持续监听客户端消息
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("客户端断开: %s (%v)", conn.RemoteAddr(), err)
				clients[conn] = false
				conn.Close()
				delete(clients, conn)
				sendMessageToAll(conn, "客户端"+conn.RemoteAddr().String()+"已断开")
				break
			}
			log.Printf("收到来自 %s 的消息: %s", conn.RemoteAddr(), string(message))

			// 将消息广播给所有客户端
			sendMessageToAll(conn, string(message))
		}
	})

	log.Println("服务端启动，监听 :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
