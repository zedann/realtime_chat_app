package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zedann/realtime_chat_app/server/db"
	"github.com/zedann/realtime_chat_app/server/internal/user"
	"github.com/zedann/realtime_chat_app/server/router"
)

func main() {

	godotenv.Load()

	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatal("could not connect to database : ", err)
	}

	userRepo := user.NewRepository(dbConn.GetDB())

	userSvc := user.NewService(userRepo)

	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)

	router.Start(os.Getenv("LISTEN_ADDR"))

	fmt.Printf("Server running %s\n", os.Getenv("LISTEN_ADDR"))
}
