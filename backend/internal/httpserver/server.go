package httpserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"go.uber.org/fx"

	article "github.com/oryaacov/personal-blog/internal/articles"
	"github.com/oryaacov/personal-blog/internal/common"
)

type HTTPServer struct {
	Router *gin.Engine
	deps   HTTPServerDeps
}

type HTTPServerDeps struct {
	fx.In
	Config         *common.Config
	ArticleHandler *article.ArticleHandler
}

func InitHttpServer(deps HTTPServerDeps) {
	router := gin.New()
	router.Use(cors.Middleware(cors.Config{
		Origins:         deps.Config.WebServer.AllowedOrigins,
		Methods:         deps.Config.WebServer.AllowedMethods,
		RequestHeaders:  deps.Config.WebServer.AllowedHeaders,
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.Use(gin.Recovery())
	router.Use(VerifyHeader())

	router.StaticFS(fmt.Sprintf("/%s",
		deps.Config.WebServer.BaseURL),
		http.Dir(deps.Config.WebServer.StaticFilesLocation))

	router.GET("/api/v1/articles", deps.ArticleHandler.Test)

	router.Run(fmt.Sprintf(":8080"))
}
