package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/GibranHL0/devblog/connection"
	"github.com/GibranHL0/devblog/errorhandler"
	"github.com/GibranHL0/devblog/models"
	"github.com/GibranHL0/devblog/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetArticle(
	id string, 
	db *connection.Database, 
	article *models.Article) error {
	
	// Create the ObjectId for the Mongo DB
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mr := repository.BuildMongo(db)

	// Obtain the result from the database
	result, err := mr.GetArticle(objectId)
	if err != nil {
		info := fmt.Sprintf("%s article caused the error", id)
		errorhandler.ReportError(err, info)

		return err
	}

	// Finally, Decode the result into the struct
	return result.Decode(&article)
}

func GetArticlesView(
	db *connection.Database,
	articles *[]models.ArticleView,
	skip int64,
	limit int64) error {

	mr := repository.BuildMongo(db)
	
	// Create a context with the timing in the configuration
	ctx, cancel := context.WithTimeout(context.Background(), mr.Db.Timing)
	defer cancel()

	// Obtains the cursor with all the results
	cursor, err := mr.GetArticlesView(skip, limit)
	if err != nil {
		info := fmt.Sprintf(
			"Home View couldn't load by %v cursor issue", cursor.ID(),
		)
		errorhandler.ReportError(err, info)
		return err
	}

	// Decode all the info
	return cursor.All(ctx, articles)
}

func CountArticles(db *connection.Database) (int64, error) {
	mr := repository.BuildMongo(db)

	articles, err := mr.CountArticles()
	if err != nil {
		errorhandler.ReportError(err, "services: Couldn't count home articles")
		return 0, err
	}

	return articles, nil
}

func CreateSub(db *connection.Database ,sub models.Subscriber) (error) {
	mr := repository.BuildMongo(db)

	result, err := mr.CreateSub(sub)
	if err != nil {
		return err
	}

	if result.InsertedID == nil {
		err := errors.New("services: Couldn't create sub")
		errorhandler.ReportError(err, sub.Email)
		return err
	}

	return nil
}