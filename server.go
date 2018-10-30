package main

import (
  "fmt"
  "net"
  "os"
)

const (
  CONN_HOST = "lab.posevin.com"
  CONN_PORT = "4001"
  CONN_TYPE = "tcp"
)

func main() {
  l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
  if err != nil {
    fmt.Println("Error listening:", err.Error())
    os.Exit(1)
  }
  defer l.Close()

  fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      os.Exit(1)
    }
    go handleRequest(conn)
  }
}

func handleRequest(conn net.Conn) {
  buf := make([]byte, 1024)
  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  conn.Write([]byte("Message received."))
  conn.Write([]byte(buf))
  fmt.Println("Client say: " + string(buf))
  conn.Close()
}
