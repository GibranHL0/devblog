package repository

import "github.com/GibranHL0/devblog/connection"

func BuildMongo(db *connection.Database) mongorepository {
	return mongorepository{
		Db: db,
	}
}