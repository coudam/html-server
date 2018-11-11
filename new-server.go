package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  2048,
	WriteBufferSize: 2048,
	CheckOrigin: func(e *http.Request) bool{ return true ;},
}


func convertMess(str []byte) string {
	i:=0
	t:= ""
	l:= len(str)
	for  i < l && str[i]!= ' ' {
		t+= string(str[i]);
		i++
	}
	
	i++
	key, err := strconv.Atoi(t)
	if err!=nil || i>=l || str[i]!= ':' {
		return "wrong key"
	}
	i+=2

	key%=26
	ans:= ""
	fmt.Println(key)
	for i< l && str[i]!= ' ' {
		fmt.Println(int(str[i]))
		ans+= string(int('a')+ (int(str[i])- int('a')+key)%26)
		i++
	}

	return ans
}

func sendMass(w *websocket.Conn){
	for {
			fmt.Println("Reading message...")
			_, msg, err := w.ReadMessage()
			if err != nil {
				return
			}

			fmt.Println("client said \""+string(msg)+"\"")
			ans:=convertMess(msg)

			ws, err := w.NextWriter(websocket.TextMessage)
			
			if err != nil{
				break 
			}
			ws.Write([]byte(ans))
			ws.Close() 
			fmt.Println("")
		}
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err!= nil{
			fmt.Println(err)
		}

		go sendMass(conn)
	})

	fmt.Println("listening at :8080...")
	http.ListenAndServe(":8080", nil)
}
