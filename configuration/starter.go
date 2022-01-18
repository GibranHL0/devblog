package configuration

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/GibranHL0/devblog/errorhandler"
)

type Configuration struct {
	Uri            string
	Database       string
	Collection     string
	WaitingTime    time.Duration
	PoolSize       uint64
	FileServerPath string
	TemplatesPath  string
	Port           string
	Url            string
}

func getEnv(env string) string {
	var result string

	if value, ok := os.LookupEnv(env); ok {
		result = value
	} else {
		errorhandler.RaiseFatal(fmt.Sprintf("Env variable not found: %s", env))
	}

	return result
}

func Get() *Configuration {
	size, err := strconv.ParseUint(getEnv("POOLSIZE"), 10, 32)
	errorhandler.CheckFatal(err)

	timing, err := strconv.ParseInt(getEnv("WAITING"), 10, 32)
	errorhandler.CheckFatal(err)

	uri := getEnv("MONGO_URI")
	db := getEnv("DB")
	collection := getEnv("COLLECTION")
	fileServerPath := getEnv("FILESERVER")
	templatesPath := getEnv("TEMPLATESPATH")
	port := getEnv("PORT")
	url := getEnv("URL")

	return &Configuration{
		Uri:            uri,
		Database:       db,
		Collection:     collection,
		WaitingTime:    time.Duration(timing * int64(time.Second)),
		PoolSize:       size,
		FileServerPath: fileServerPath,
		TemplatesPath:  templatesPath,
		Port:           port,
		Url:            url,
	}
}
