package main

import (
	"log"
	"net"
)

const (
  HOST = "127.0.0.1"
  PORT = "8080"
  TYPE = "tcp"
)


func start_server() error {
  l, err := net.Listen(TYPE, HOST+":"+PORT)
  
  if err != nil {
    log.Printf("failed starting the listener: %v\n", err)
    return err
  }
  
  defer l.Close()
  
  for {
    conn, err := l.Accept()
    if err != nil {
      log.Printf("failed accepting a connection: %v\n", err)
      return err
    }
    
    go func(c net.Conn){  
      for {
        message := make([]byte, 1000)
        c.Read(message)
        log.Println(c.LocalAddr().String() + " says: " + string(message))
      }
    } (conn)
  }
  return nil
}


