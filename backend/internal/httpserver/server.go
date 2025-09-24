package httpserver

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.uber.org/fx"

	"github.com/oryaacov/personal-blog/internal/common"
	"github.com/oryaacov/personal-blog/internal/controllers"
)

type HTTPServer struct {
	Router *gin.Engine
	deps   HTTPServerDeps
}

type HTTPServerDeps struct {
	fx.In
	Config common.Config
	controllers.ArticleController
	controllers.ThumbnailsController
}

func (s *HTTPServer) setRoutes() {
	s.Router.GET("/api/v1/thumbnails", s.deps.ThumbnailsController.GetAll)
	s.Router.GET("/api/v1/thumbnails/:category", s.deps.ThumbnailsController.GetByCategory)
	s.Router.GET("/api/v1/articles/:id", s.deps.ArticleController.GetArticleById)
}

func (s *HTTPServer) setMiddlewares() {
	s.Router.Use(cors.Middleware(cors.Config{
		Origins:         s.deps.Config.WebServer.AllowedOrigins,
		Methods:         s.deps.Config.WebServer.AllowedMethods,
		RequestHeaders:  s.deps.Config.WebServer.AllowedHeaders,
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	s.Router.Use(gin.Recovery())
	s.Router.Use(VerifyHeader())
}

func NewHttpServer(deps HTTPServerDeps) HTTPServer {
	server := HTTPServer{Router: gin.New(), deps: deps}
	server.setMiddlewares()
	server.setRoutes()

	return server
}

func ListenAndServer(server HTTPServer) {
	server.Router.Run(fmt.Sprintf("%s:%d", server.deps.Config.WebServer.BaseURL, server.deps.Config.WebServer.Port))
}
