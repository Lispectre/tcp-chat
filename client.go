package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
  "log"
)

func start_client() error {
  conn, err := net.Dial(TYPE, HOST+":"+PORT)
  if err != nil {
    log.Println("couldn't dial to server")
    return err
  }

  fmt.Println("Connected!")
  
  go func(c net.Conn){
    received := make([]byte, 1000)
    c.Read(received)
    fmt.Println("< " + string(received))
  } (conn)
  
  for reader := bufio.NewReader(os.Stdin);; {
    fmt.Print("> ")
  
    text, err := reader.ReadString('\n')

    if err != nil {
      fmt.Println("input error! message was not sent")
      continue
    }

    if text == "exit" {
      break
    }

    _, err = conn.Write([]byte(text))
    if err != nil {
      fmt.Println("connection error! message was not sent")
    }
  }

  return nil
}
