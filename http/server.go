package http

import (
	"fmt"
	"net/http"

	env "mck-p/modi/env"
	mclog "mck-p/modi/log"
	views "mck-p/modi/view"

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

func setupHandler(handler *gin.Engine) *gin.Engine {
	/*
		Set up Global Headers
	*/
	handler.Use(SetGlobalHeaders())

	/*
		Set up View Layer Configuration
	*/
	handler.HTMLRender = views.LoadTemplates("view")

	/*
		Handle Static Assets

		TODO: Replace Static Assets with S3/CDN
	*/
	handler.Static("/assets", "./artifacts/assets")

	/*
		Assign HTTP Handlers
	*/
	handler.GET("/", views.Home)

	handler.GET("/healthcheck", Healthcheck)

	handler.GET("/authors", FindAllAuthors)
	handler.GET("/authors/:id", FindAuthorById)
	handler.POST("/authors", CreateAuthor)

	return handler
}

func CreateServer() Server {
	log := mclog.CreateLogger("SERVER")
	handler := setupHandler(gin.Default())

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
