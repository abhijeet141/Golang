package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)
	fileName, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	Infolog := log.New(fileName, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("Hello Log file")
	fmt.Println(&buf)
	log.Println("INFO: Hello")
	Infolog.Println("Error")
}
