package example

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func example() {

	myhttp := http.NewServeMux()
	myhttp.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		con, _ := upgrader.Upgrade(w, r, nil)
		_, b, er := con.ReadMessage()
		if er != nil {
			panic(er)
		}

		str := string(b)
		println(str)
	})

	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", myhttp)
}
