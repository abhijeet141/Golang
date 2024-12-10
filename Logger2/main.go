package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	checkError(err)
	defer file.Close()
	//lstdflag is equivalent to ldate and ltime
	Infolog := log.New(file, "INFO: ", log.Llongfile|log.Lmicroseconds|log.LstdFlags)
	Errorlog := log.New(file, "Error: ", log.Llongfile|log.Lmicroseconds|log.LstdFlags)
	Infolog.Println("This is log file")
	Errorlog.Println("Error")
}
func checkError(err error) {
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
}
