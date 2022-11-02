package heroes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type HeroRepository struct {
	DBconn *mongo.Database
}

const collectionNameHero string = "hero"

func (h HeroRepository) GetAllHero(ctx context.Context) ([]Hero, error) {
	heroes := []Hero{}

	cur, err := h.DBconn.Collection(collectionNameHero).Find(ctx, bson.D{{}})
	cur.All(ctx, &heroes)

	if err != nil {
		return nil, err
	}

	return heroes, nil
}
