package main

import "go-echo-real-project/internal/server"

func main() {
	s := server.Server{}
	s.AutoInject()
	s.Listen()
}
