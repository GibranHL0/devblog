package repository

import (
	"context"

	"github.com/GibranHL0/devblog/connection"
	"github.com/GibranHL0/devblog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongorepository struct {
	Db *connection.Database
}

func (mr *mongorepository) GetArticle(id primitive.ObjectID) (
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

func (mr *mongorepository) GetArticlesView(skip int64, limit int64) (
	*mongo.Cursor, error) {

	filter := bson.M{}
	// sorting := bson.D{{"date", 1}}
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	articlesView := mr.Db.Collection.Database().Collection("HomeView")

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

func (mr *mongorepository) CountArticles() (articles int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	articles, err = mr.Db.Collection.CountDocuments(ctx, bson.D{}, nil)
	if err != nil {
		return 0, err
	}

	return articles, nil
}

func (mr *mongorepository) CreateSub(sub models.Subscriber) (
	*mongo.InsertOneResult, error) {
	subCollection := mr.Db.Collection.Database().Collection("subscribers")

	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	result, err := subCollection.InsertOne(ctx, sub)
	if err != nil {
		return nil, err
	}

	return result, nil
}