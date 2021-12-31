package models

import (
	"time"
)

type Article struct {
	Title   string    `bson:"title"`
	Author  string    `bson:"author"`
	Date    time.Time `bson:"date"`
	Tags    []string  `bson:"tags"`
	Content string    `bson:"article"`
}

type ArticleView struct {
	Id          string    `bson:"_id"`
	Picture     string    `bson:"picture"`
	Title       string    `bson:"title"`
	Author      string    `bson:"author"`
	Date        time.Time `bson:"date"`
	Description string    `bson:"description"`
	Tags        []string  `bson:"tags"`
}

type HomePage struct {
	Articles  []ArticleView
	HasBefore bool
	HasNext   bool
	Current   int64
	Older     int64
	Newer     int64
	Pages     []int64
}
