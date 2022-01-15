package server

import (
	"net/http"
	"text/template"
	"time"

	"github.com/GibranHL0/devblog/connection"
	"github.com/GibranHL0/devblog/models"
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

func newsletterHandler(templates *template.Template, db *connection.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			executeTemplate(templates, "newsletter.html", nil, rw)
			return
		}

		subscriber := models.Subscriber{
			Email:  r.FormValue("email"),
			SignOn: time.Now(),
			Enable: true,
		}

		services.CreateSub(db, subscriber)

		executeTemplate(templates, "newsletter.html", struct{ Success bool }{true}, rw)
	}
}

func homeHandler(templates *template.Template, db *connection.Database, url string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			executeNotFound(rw, templates)
			return
		}

		pagenumber := getPageNumber(r.URL.Query().Get("page"))
		totalarticles, err := services.CountArticles(db)
		if err != nil {
			executeNotFound(rw, templates)
			return
		}
		maxpages := getMaxPages(totalarticles)
		pagenumber = normalizePageNumber(pagenumber, maxpages)
		skip := calculateSkip(pagenumber)
		articles := make([]models.ArticleView, 0)

		err = services.GetArticlesView(db, &articles, skip, 9)
		if err != nil {
			executeNotFound(rw, templates)
			return
		}

		homepage := models.GetHomePage(articles, pagenumber, maxpages, url)

		executeTemplate(templates, "home.html", homepage, rw)
	}
}

func articleHandler(templates *template.Template, db *connection.Database) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		articleId := r.URL.Query().Get("id")
		if articleId == "" {
			executeNotFound(rw, templates)
			return
		}

		article := models.Article{}

		err := services.GetArticle(articleId, db, &article)
		if err != nil {
			executeNotFound(rw, templates)
			return
		}

		executeTemplate(templates, "article.html", article, rw)
	}
}
