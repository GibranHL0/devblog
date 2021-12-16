package server

import (
	"net/http"
	"text/template"

	"github.com/GibranHL0/devblog/connection"
	"github.com/GibranHL0/devblog/models"
	"github.com/GibranHL0/devblog/repository"
	"github.com/GibranHL0/devblog/services"
)

func aboutHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "about.html", nil, rw)
	}
}

func contactHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "contact.html", nil, rw)
	}
}

func newsletterHandler(templates *template.Template) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		executeTemplate(templates, "newsletter.html", nil, rw)
	}
}

func articleHandler(templates *template.Template, db *connection.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		articleId := r.URL.Query().Get("id")
		if articleId == "" {
			executeNotFound(rw)
			return
		}

		article := models.Article{}
		mr := repository.MongoRepository{Db: db}

		err := services.GetArticle(articleId, mr, &article)
		if err != nil {
			executeNotFound(rw)
			return
		}

		executeTemplate(templates, "article.html", article, rw)
	}
}
