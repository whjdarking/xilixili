package main

import (
	"Refactor_xilixili/config"
	"Refactor_xilixili/server"
)

func main() {
	config.Init()

	r := server.NewRouter()
	r.Run(":3000")
}
