package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	article "github.com/oryaacov/personal-blog/internal/articles"
	"github.com/oryaacov/personal-blog/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ArticleController struct {
	deps ArticleContorllerDeps
}

type ArticleContorllerDeps struct {
	fx.In
	ArticleHandler article.ArticleHandler
}

func NewArticleController(deps ArticleContorllerDeps) ArticleController {
	return ArticleController{deps: deps}
}

func (a *ArticleController) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	articles, err := a.deps.ArticleHandler.GetAll(ctx)
	if err != nil {
		logger.Log().Error("failed to get all articles", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, articles)
}
