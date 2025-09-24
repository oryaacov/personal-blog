package thumbnails

import (
	"context"
	"fmt"

	"github.com/oryaacov/personal-blog/internal/common"
	"github.com/oryaacov/personal-blog/internal/core"
	"github.com/oryaacov/personal-blog/internal/models"
	"github.com/oryaacov/personal-blog/pkg/mongoutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

const (
	thumbnailsLimit = 50
)

var (
	getAllSortByPriorityDescOptions *options.FindOptions
)

func init() {
	getAllSortByPriorityDescOptions = options.Find().SetLimit(thumbnailsLimit)
	getAllSortByPriorityDescOptions.SetSort( bson.D{{Key: "priority", Value: -1}})
}

type ThumbnailsHandler struct {
	deps ThumbnailsHandlerDeps
}

type ThumbnailsHandlerDeps struct {
	fx.In
	Collection *mongo.Collection
}

func NewThumbnailsHandler(DB *core.DB, Config common.Config) ThumbnailsHandler {
	return ThumbnailsHandler{
		deps: ThumbnailsHandlerDeps{
			Collection: DB.GetCollection(Config.Database.Collections.Thumbnails),
		},
	}
}

func (h *ThumbnailsHandler) GetAll(ctx context.Context) ([]models.Thumbnail, error) {
	cur, err := h.deps.Collection.Find(ctx, mongoutils.EmptyFilter, getAllSortByPriorityDescOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query thumbnails with: %w", err)
	}

	return h.buildArticles(ctx, cur)
}


func (h *ThumbnailsHandler) GetByCategory(ctx context.Context, category string) ([]models.Thumbnail, error) {
	cur, err := h.deps.Collection.Find(ctx, bson.M{"category": category}, getAllSortByPriorityDescOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query thumbnails with: %w", err)
	}

	return h.buildArticles(ctx, cur)
}


func (h *ThumbnailsHandler) buildArticles(ctx context.Context, cur *mongo.Cursor) ([]models.Thumbnail, error) {
	res := make([]models.Thumbnail, 0, thumbnailsLimit)
	for cur.Next(ctx) {
		var thumbnail models.Thumbnail
		err := cur.Decode(&thumbnail)
		if err != nil {
			return nil, fmt.Errorf("failed to decode thumbnail with: %w", err)
		}

		res = append(res, thumbnail)
	}

	return res, nil
}