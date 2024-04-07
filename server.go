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
    
    defer conn.Close()

    go func(c net.Conn){  
      log.Printf("Connected: %v\n", c.LocalAddr().String())
      for {
        var message []byte
        _, err := c.Read(message)
        if err != nil {
          break
        }
        log.Printf("%v says: %v", c.LocalAddr().String(), string(message))
      }
    } (conn)
  }
  return nil
}


