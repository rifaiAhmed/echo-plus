package main

import (
	"fmt"
	"time"

	"test-echo/config"
	"test-echo/db"
	mid "test-echo/middleware"
	"test-echo/routes"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Tidak ada env")
		//logs.InfoLogger.Println("- Info Main : No .env file found ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	db.Init()
	e := routes.Init()
	e.HTTPErrorHandler = mid.ErrorHandler

	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(":9191") }(lock)

	time.Sleep(1 * time.Millisecond)
	config.MakeLogEntry(nil).Warning("application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		config.MakeLogEntry(nil).Panic("failed to start application")
	}

	//e.Logger.Fatal(e.Start(":9191"))
}
