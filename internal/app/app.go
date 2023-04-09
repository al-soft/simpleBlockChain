package app

import (
	"simpleBlockChain/internal/webserver"

	"github.com/joho/godotenv"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := webserver.Run(); err != nil {
		panic(err)
	}
}
