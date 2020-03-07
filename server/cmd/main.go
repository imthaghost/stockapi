package main

import (
	"github.com/imthaghost/stockapi/server"
)

func main() {

	s := server.NewServer()
	s.Start(":8000")
}
