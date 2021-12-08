package main

import (
	"fmt"
	"log"

	"github.com/GibranHL0/devblog/configuration"
	"github.com/GibranHL0/devblog/connection"
)

func init() {
	config := configuration.Get(".env")
	log.Println("Configuration generated: ", config.Database)

	db := connection.StartMongo(*config)
	log.Println("Connected to: ", db.Mongo.Name())
}

func main() {

	fmt.Println("Hey there 🦑 ...")

	// output := markdown.ToHTML([]byte(jfile.Article), nil, nil)

	// fmt.Println(string(output))
}
