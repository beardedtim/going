package http

import (
	"fmt"
	"net/http"

	env "mkc-p/modi/env"
	mclog "mkc-p/modi/log"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Close()
	Start()
}

type Server struct {
	IServer
	Raw http.Server
	Log mclog.Logger
}

func (server Server) Start() {
	if err := server.Raw.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			server.Log.Warn("Server closed under request")
		} else {
			server.Log.Info(err.Error())
			server.Log.Error("Server closed unexpect")
		}
	}
}

func (server Server) Stop() error {
	return server.Raw.Close()
}

func CreateServer() Server {
	log := mclog.CreateLogger("SERVER")
	handler := gin.Default()

	handler.GET("/healthcheck", Healthcheck)
	handler.GET("/authors", FindAllAuthors)
	handler.POST("/authors", CreateAuthor)

	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	rawServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.GetEnv("PORT", "8080")),
		Handler: handler,
	}

	server := Server{
		Raw: *rawServer,
		Log: log,
	}

	return server
}
