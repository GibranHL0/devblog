package repository

import (
	"context"

	"github.com/GibranHL0/devblog/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	Db *connection.Database
}

func (mr *MongoRepository) GetArticle(id primitive.ObjectID) (
	*mongo.SingleResult, error) {

	filter := bson.M{"_id": id}
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	result := mr.Db.Collection.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (mr *MongoRepository) GetHomeView(skip int64, limit int64) (
	*mongo.Cursor, error) {

	filter := bson.M{}
	// sorting := bson.D{{"date", 1}}
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	articlesView := mr.Db.Collection.Database().Collection("HomeArticles")

	findOptions := options.FindOptions{
		Skip: &skip,
		Limit: &limit,
	}

	cursor, err := articlesView.Find(
		ctx,
		filter,
		&findOptions,
	)

	if err != nil {
		return nil, err
	}

	return cursor, nil
}
