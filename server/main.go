package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var forumUser = make(map[int]string)
forumUser[1] = "MrBin"


type Client struct {
	user *forumUser,
	conn *websocket.Conn,
}

var clients = make(map[int]*Client)


// Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home Page")
// }

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "WebSocket Endpoint")
	//allow any connections into my web socket regeadless of what the origin of that connection is

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Client Successfully Connected...")

	go reader(ws)
	writer(ws)
}

//tackle the problem of not listening permanently for incomming messages

// reader takes websocket connection
func reader(conn *websocket.Conn) {
	//kick-off a for loop
	for {
		//return the message type p-bytes or error
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(p))

		// I want to be able to echo back that message to the client
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func writer(conn *websocket.Conn) {

	conn.WriteMessage(websocket.TextMessage, []byte("Hello from server"))

}

func setUpRoutes() {
	// http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	// fmt.Println(("Go WebSockets"))
	setUpRoutes()
	fmt.Println(("Port 8080 is running"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
