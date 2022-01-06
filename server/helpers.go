package server

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/GibranHL0/devblog/errorhandler"
	"github.com/GibranHL0/devblog/models"
)

func executeTemplate(
	templates *template.Template, 
	name string, 
	data models.IModel,
	rw http.ResponseWriter){
	
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		info := fmt.Sprintf("%s template was not executed", name)
		errorhandler.ReportError(err, info)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Something bad happened!"))
	}
}

func executeNotFound(rw http.ResponseWriter, templates *template.Template) {
	rw.WriteHeader(http.StatusNotFound)
	executeTemplate(templates, "404.html", nil, rw)
}

func getPageNumber(page string) (pagenumber int64) {
	pagenumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 1
	}

	return pagenumber
}

func calculateSkip(pagenumber int64) int64 {
	skip := (pagenumber - 1) * 9

	return skip
}

func normalizePageNumber(pagenumber int64, maxpages int64) int64 {
	// Eliminate the issue with negative page numbers
	if pagenumber < 1 { pagenumber = 1 }

	if pagenumber > maxpages {
		pagenumber = maxpages
	}

	return pagenumber
}

func getMaxPages(totalarticles int64) int64 {
	return int64(math.Ceil(float64(totalarticles) / 9))
}