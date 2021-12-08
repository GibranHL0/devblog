package errorhandler

import "log"

func CheckFatal(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func RaiseFatal(info string) {
	log.Fatal(info)
}