package main

import (
	"Dp-218_Go/configs"
	"Dp-218_Go/server"
)

func main() {
	cfg := configs.Get()
	err := server.Run(cfg)
	if err != nil {
		return 
	}
}

