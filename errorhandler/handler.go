package errorhandler

import (
	"log"
)

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RaiseFatal(info string) {
	log.Fatal(info)
}

func ReportError(err error, info string) {
	log.Println(info, err)
}
