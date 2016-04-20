package main

import (
	"os"

	"github.com/cloudnativego/wof-eventprocessing/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	server := service.NewServer()
	server.Run(":" + port)
}
