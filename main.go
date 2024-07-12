package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	// "golang.org/x/net/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Server struct {
	conn map[*websocket.Conn]bool
}

type Message struct {
	Data interface{} `json:"message"`
	Type string      `json:"type"`
}

// handler for websocket-request
// func (s *Server) handleWs(ws *websocket.Conn) {
func (s *Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Printf("New User Connected %s ", conn.RemoteAddr())
	s.conn[conn] = true
	fmt.Println("total users", len(s.conn))

	// s.readLoop(conn)
	// go s.readLoop(conn)

	// Open the file
	// file, err := os.Open("dummy.png")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer file.Close()

	// // Read the file into a byte slice
	// data, err := io.ReadAll(file)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// // Create a new Message struct with the file data
	// msg := &Message{
	// 	Type: "file",
	// 	Data: data,
	// }

	// *sync.Mutex.Lock()
	// err = conn.WriteJSON(msg)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	s.readLoop(conn)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	for {
		my, msg, err := ws.ReadMessage()
		// ws.cookie
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		fmt.Println(my)

		var m Message
		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Println("Error unmarshalling message:", err)
			return
		}

		log.Printf("Received message: %s   %v+\n", m.Type, m.Data)

		s.broadcast(&m)
	}
}

func (s *Server) broadcast(b *Message) {
	for client := range s.conn {
		fmt.Print("hello")
		go func(ws *websocket.Conn, msg *Message) {
			if err := ws.WriteJSON(msg); err != nil {
				fmt.Println(err)
				ws.SetCloseHandler(func(code int, text string) error {
					fmt.Printf("User Disconnected %s\n", ws.RemoteAddr())
					delete(s.conn, ws)
					fmt.Println("total users", len(s.conn))
					return nil
				})
			}
		}(client, b)
	}
}

func main() {
	cache = make(map[string]*cacheFile)
	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	serve.ListenAndServe()
}
