package main

import (
	"github.com/go-ready/core"
	"github.com/go-ready/core/pkg/health"
	"github.com/go-ready/core/pkg/server"
	"go.uber.org/fx"
	"net/http"
)

func main() {
	fx.New(
		fx.Provide(health.NewHealth),
		fx.Provide(core.NewCoreHandler),
		fx.Provide(server.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
