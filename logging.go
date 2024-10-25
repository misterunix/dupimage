package main

import (
	"log"
	"os"
)

func openLogFile(path string) (*os.File, error) {

	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func createLog() {
	lfn := "./dupimage.log"

	// delete the old logfile for now
	_, err := os.Stat(lfn)
	if err != nil {
		os.Remove(lfn)
	}

	file, err := openLogFile("./dupimage.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}
