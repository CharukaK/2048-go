package logger

import (
	"log"
	"os"
)

const LogEnvName = "LOGS_ENABLED"

func PrintInfo(msg string) {
	if _, ok := os.LookupEnv(LogEnvName); ok {
		log.SetPrefix("[Info] ")
		log.Println(msg)
	}
}

func PrintError(msg string) {
	if _, ok := os.LookupEnv(LogEnvName); ok {
		log.SetPrefix("[Error] ")
		log.Println(msg)
	}
}

func init() {
	if _, ok := os.LookupEnv(LogEnvName); !ok {
		return
	}
	file, err := os.OpenFile("game.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Fatal("Error opening log file")
	}

	log.SetOutput(file)
}
