package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)

	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer conn.Close()

	log.Println("已连接服务端，输入文字发送（按 Ctrl+C 退出）")

	// 启动协程读取用户输入
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
				log.Println("发送失败:", err)
				return
			}
		}
	}()

	// 保持主程序运行
	for {
	}
}
