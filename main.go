package main

import (
	"flag"
	"fmt"

	"mtlstester/client"
	"mtlstester/config"
	"mtlstester/server"
)

func main() {
	cfg := config.Parse()

	if !cfg.IsRunModeValid() {
		fmt.Println("\nUsage: go run main.go -run (server|client)")
		flag.PrintDefaults()
		return
	}

	if cfg.IsServer() {
		server.Run(cfg)
	} else if cfg.IsClient() {
		client.Run(cfg)
	} else {
		flag.PrintDefaults()
	}

}
