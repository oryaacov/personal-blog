package main

import (
	article "github.com/oryaacov/personal-blog/internal/articles"
	"github.com/oryaacov/personal-blog/internal/common"
	"github.com/oryaacov/personal-blog/internal/controllers"
	"github.com/oryaacov/personal-blog/internal/core"
	"github.com/oryaacov/personal-blog/internal/httpserver"
	"github.com/oryaacov/personal-blog/pkg/service"
	"go.uber.org/fx"
)

const (
	serviceName = "personal-blog"
)

func main() {
	fx.New(
		fx.Provide(
			func() service.Name { return service.ServiceName(serviceName) },
			common.BuildConfig,
			core.InitDB,
			article.NewArticleHandler,
			controllers.NewArticleController,
			httpserver.NewHttpServer,
		),
		fx.Invoke(
			httpserver.ListenAndServer,
		),
	)
}
