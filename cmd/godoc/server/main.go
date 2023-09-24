package main

import "golang-mono/godoc-test/server"

const (
	ADDR = "127.0.0.1:9000"
)

func main() {
	server := server.CreateServer()
	server.Run(ADDR)
}
