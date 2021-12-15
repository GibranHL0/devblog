package services

import (
	"context"

	"github.com/GibranHL0/devblog/models"
	"github.com/GibranHL0/devblog/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetArticle(
	id string, 
	mr repository.MongoRepository, 
	article *models.Article) error {
	
	// Create the ObjectId for the Mongo DB
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Obtain the result from the database
	result, err := mr.GetArticle(objectId)
	if err != nil {
		return err
	}

	// Finally, Decode the result into the struct
	return result.Decode(&article)
}

func GetHomeView(
	mr repository.MongoRepository,
	views []*models.HomeView,
	skip int64,
	limit int64) error {
	
	// Create a context with the timing in the configuration
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	// Obtains the cursor with all the results
	cursor, err := mr.GetHomeView(skip, limit)
	if err != nil {
		return err
	}

	// Decode all the info
	return cursor.All(ctx, views)
}