package main

import (
	"fmt"
	"log"

	"github.com/GibranHL0/devblog/app"
	"github.com/GibranHL0/devblog/configuration"
	"github.com/GibranHL0/devblog/connection"
	"github.com/GibranHL0/devblog/server"
	"github.com/GibranHL0/devblog/templates"
)

var blog app.App

func init() {
	config := configuration.Get(".env")
	log.Println("Configuration generated: ", config.Database)

	db := connection.StartMongo(*config)
	log.Println("Connected to: ", db.Mongo.Name())

	templates := templates.Load(*config)
	log.Println("Templates loaded", templates)

	mux := server.StartMux(*config)
	log.Println("Server up! 🚀 ", mux)

	blog = app.App{
		Config: config,
		Db: db,
		Templates: templates,
		Server: mux,
	}

	log.Println("App is up and running! 🔥", blog)
}

func main() {

	fmt.Println("Hey there 🦑 ...")

	// output := markdown.ToHTML([]byte(jfile.Article), nil, nil)

	// fmt.Println(string(output))
}
