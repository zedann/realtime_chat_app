package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zedann/realtime_chat_app/server/db"
)

func main() {

	godotenv.Load()

	_, err := db.NewDatabase()

	if err != nil {
		log.Fatal("could not connect to database : ", err)
	}

	fmt.Printf("Server running %s\n", os.Getenv("LISTEN_ADDR"))
}
