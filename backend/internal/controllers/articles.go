package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	article "github.com/oryaacov/personal-blog/internal/articles"
	"github.com/oryaacov/personal-blog/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
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

func (a *ArticleController) GetArticleById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	logger := log.Log().With(zap.String("id", id))
	logger.Debug("getting article by id")

	articles, err := a.deps.ArticleHandler.GetByID(ctx, c.Param("id"))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Info("not found")
			c.Status(404)
			return
		}

		logger.Error("failed to get article by id", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, articles)
}
