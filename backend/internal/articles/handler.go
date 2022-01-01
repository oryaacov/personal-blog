package article

import (
	"context"
	"fmt"

	"github.com/oryaacov/personal-blog/internal/common"
	"github.com/oryaacov/personal-blog/internal/core"
	"github.com/oryaacov/personal-blog/internal/models"
	"github.com/oryaacov/personal-blog/pkg/mongoutils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

const (
	articlesLimit = 20
)

var (
	defaultGetAllOptions = options.Find().SetLimit(articlesLimit)
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

func (h *ArticleHandler) GetAll(ctx context.Context) ([]models.Article, error) {
	cur, err := h.deps.Collection.Find(ctx, mongoutils.EmptyFilter, defaultGetAllOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query articles with: %w", err)
	}

	res := make([]models.Article, 0, articlesLimit)
	for cur.Next(ctx) {
		var article models.Article
		err = cur.Decode(&article)
		if err != nil {
			return nil, fmt.Errorf("failed to decode article %d with %w", cur.ID(), err)
		}

		res = append(res, article)

	}

	return res, nil
}
