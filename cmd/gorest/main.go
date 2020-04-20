package main

import (
	"github.com/urunimi/ddd-go/internal/app/server"
)

//noinspection GoUnhandledErrorResult
func main() {
	server.NewServer().Start()
}
