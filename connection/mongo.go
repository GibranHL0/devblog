package connection

import (
	"context"
	"time"

	"github.com/GibranHL0/devblog/configuration"
	"github.com/GibranHL0/devblog/errorhandler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Collection *mongo.Collection
	Timing time.Duration
}

func checkConnection(client *mongo.Client, ctx context.Context) error {
	return client.Ping(ctx, nil)
}

func StartMongo(config configuration.Configuration) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), config.WaitingTime)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(config.Uri).
		SetMaxPoolSize(config.PoolSize)

	client, err := mongo.Connect(ctx, clientOptions)
	errorhandler.CheckFatal(err)

	err = checkConnection(client, ctx)
	errorhandler.CheckFatal(err)

	return &Database{
		Collection: client.Database(config.Database).
			Collection(config.Collection),
		Timing: config.WaitingTime,
	}
}
