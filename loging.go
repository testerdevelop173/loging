package main

import (
	"log"
	"os"
)

func main() {
	log.Println("dfdfdf")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	const logFilePath = "apiLog.log"
	//file, _ := os.Create(logFilePath)
	file, err := os.Create(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("loging in file")
	file.Close()
}

