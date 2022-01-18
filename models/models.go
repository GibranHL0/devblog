package models

import (
	"time"
)

type Article struct {
	Title       string    `bson:"title"`
	Author      string    `bson:"author"`
	Date        time.Time `bson:"date"`
	Tags        []string  `bson:"tags"`
	Content     string    `bson:"article"`
	Description string    `bson:"description"`
	Picture     string    `bson:"picture"`
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
	Url       string
}

type Subscriber struct {
	Email   string    `bson:"email"`
	SignOn  time.Time `bson:"sign_on"`
	SignOff time.Time `bson:"sign_off"`
	Enable  bool      `bson:"enable"`
}
