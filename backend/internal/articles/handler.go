package article

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oryaacov/personal-blog/internal/core"
	"github.com/oryaacov/personal-blog/internal/models"
	"go.uber.org/fx"
)

type ArticleHandler struct {
	deps ArticleDeps
}

type ArticleDeps struct {
	fx.In
	core.DB
}

func CreateArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h *ArticleHandler) Test(c *gin.Context) {
	var test []models.Article
	err := h.deps.DB.DB.Get(&test, "select * from articles;")
	fmt.Println(err, test)

	c.JSON(http.StatusOK, test)
}
