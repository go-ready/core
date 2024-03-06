package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ready/core"
	"go.uber.org/fx"
	"log"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle, coreHandler *core.CoreHandler) *http.Server {
	router := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			coreHandler.RegisterHandlers(router)
			fmt.Println("Starting HTTP server at", srv.Addr)
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
