package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Not enough arguments specified. Exiting")
    os.Exit(1)
  }


  switch mode := os.Args[1]; mode {
    case "server":
      fmt.Println("Starting a server...")
      server_mode()
    case "client":
      fmt.Println("Starting as client...")
      client_mode()
    default: 
      fmt.Println("Unknown mode specified! Exiting.")
  }
}

func server_mode() {
  if err := start_server(); err != nil {
    log.Println(err)
  }
}
func client_mode() {
  if err := start_client(); err != nil {
    log.Println(err)
  }
}
