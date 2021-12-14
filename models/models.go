package models

import "time"

type Article struct {
	Title  string    `bson:"title"`
	Author string    `bson:"author"`
	Date   time.Time `bson:"date"`
	Tags   []string  `bson:"tags"`
	Body   []byte    `bson:"article"`
}

type HomeView struct {
	Picture     string    `bson:"picture"`
	Title       string    `bson:"title"`
	Author      string    `bson:"author"`
	Date        time.Time `bson:"date"`
	Description string    `bson:"description"`
	Tags        []string  `bson:"tags"`
}
