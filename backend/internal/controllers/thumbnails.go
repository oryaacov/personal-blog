package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/oryaacov/personal-blog/internal/thumbnails"
	"github.com/oryaacov/personal-blog/pkg/log"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ThumbnailsController struct {
	deps ThumbnailsContorllerDeps
}

type ThumbnailsContorllerDeps struct {
	fx.In
	ThumbnailsHandler thumbnails.ThumbnailsHandler
}

func NewThumbnailsController(deps ThumbnailsContorllerDeps) ThumbnailsController {
	return ThumbnailsController{deps: deps}
}

func (a *ThumbnailsController) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	thumbnails, err := a.deps.ThumbnailsHandler.GetAll(ctx)
	if err != nil {
		log.Log().Error("failed to get thumbnails", zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, thumbnails)
}
