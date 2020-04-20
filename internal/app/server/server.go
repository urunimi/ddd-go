package server

import (
	"github.com/urunimi/ddd-go/internal/app/api"
	"github.com/urunimi/gorest/core"
)

// NewServer Return new server instance
func NewServer() core.Server {
	agApp := api.CreateAPIApp()
	server := core.NewServer(agApp, &(miscApp{}))
	_ = server.Init()
	return server
}
