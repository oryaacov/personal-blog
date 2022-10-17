package article

import (
	"context"
	"fmt"

	"github.com/oryaacov/personal-blog/internal/common"
	"github.com/oryaacov/personal-blog/internal/core"
	"github.com/oryaacov/personal-blog/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

type ArticleHandler struct {
	deps ArticleDeps
}

type ArticleDeps struct {
	fx.In
	Collection *mongo.Collection
}

func NewArticleHandler(DB *core.DB, Config common.Config) ArticleHandler {
	return ArticleHandler{
		deps: ArticleDeps{
			Collection: DB.GetCollection(Config.Database.Collections.Article),
		},
	}
}

func (h *ArticleHandler) GetByID(ctx context.Context, id string) (*models.Article, error) {
	res := h.deps.Collection.FindOne(ctx, bson.M{"_id": id})
	if err := res.Err(); err != nil {
		return nil, fmt.Errorf("failed to query articles with: %w", err)
	}

	var article *models.Article
	if err := res.Decode(&article); err != nil {
		return nil, fmt.Errorf("failed to parse article with: %w", err)
	}

	return article, nil
}
