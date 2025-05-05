package customlog

import (
	"log"
	"os"
	"os/user"
)

func CreateLogDirectories() {
	currentUser, _ := user.Current()
	requestLogDirectory := currentUser.HomeDir + "/requestlogs"
	projectLogDirectory := currentUser.HomeDir + "/studentprojectlogs"
	err := os.MkdirAll(requestLogDirectory, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create the log directory: %+v", err)
	}

	err = os.MkdirAll(projectLogDirectory, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create the log directory: %+v", err)
	}
}

func GetLogFile(pathName string) *os.File {
	logFile, err := os.OpenFile(pathName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open request log file: %v", err)
	}
	return logFile
}
