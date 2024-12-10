package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening file:%v", err)
	}
	defer file.Close()
	InfoLog := log.New(file, "Info: ", log.LstdFlags)
	ErrorLog := log.New(file, "ERROR: ", log.Llongfile|log.Lmicroseconds|log.LstdFlags)
	InfoLog.SetFlags(log.Llongfile | log.Lmicroseconds | log.LstdFlags)
	InfoLog.Println("Application Started")
	InfoLog.Println("User Logged In")
	ErrorLog.Println("This is an error log")
}
