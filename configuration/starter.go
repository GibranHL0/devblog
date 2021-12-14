package configuration

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/GibranHL0/devblog/errorhandler"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Uri            string
	Database       string
	Collection     string
	WaitingTime    time.Duration
	PoolSize       uint64
	FileServerPath string
	TemplatesPath  string
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

func Get(envFilePath string) *Configuration {
	err := godotenv.Load(envFilePath)
	errorhandler.CheckFatal(err)

	size, err := strconv.ParseUint(getEnv("POOLSIZE"), 10, 32)
	errorhandler.CheckFatal(err)

	timing, err := strconv.ParseInt(getEnv("WAITING"), 10, 32)
	errorhandler.CheckFatal(err)

	uri := getEnv("MONGO_URI")
	db := getEnv("DB")
	collection := getEnv("COLLECTION")
	fileServerPath := getEnv("FILESERVER")
	templatesPath := getEnv("TEMPLATESPATH")

	return &Configuration{
		Uri:            uri,
		Database:       db,
		Collection:     collection,
		WaitingTime:    time.Duration(timing * int64(time.Second)),
		PoolSize:       size,
		FileServerPath: fileServerPath,
		TemplatesPath:  templatesPath,
	}
}
