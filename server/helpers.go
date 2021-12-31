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

func executeNotFound(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("Content not found!, 404"))
}

func getPageNumber(page string) (pagenumber int64) {
	pagenumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 1
	}

	return pagenumber
}

func calculateSkip(totalarticles int64, pagenumber int64) int64 {
	// Eliminate the issue with negative page numbers
	if pagenumber < 1 { pagenumber = 1 }

	maxpages := int64(math.Ceil(float64(totalarticles) / 9))

	if pagenumber > maxpages {
		pagenumber = maxpages
	}

	skip := (pagenumber - 1) * 9

	return skip
}