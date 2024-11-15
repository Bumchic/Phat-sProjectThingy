package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/components/appContext"
	"server/middleware"
	"server/routes"
	"server/storage/userstore"

	"github.com/joho/godotenv"
)

func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}
	
	appCtx := appContext.InitAppContext(uri)
	defer appCtx.CloseDB()
	
	
	userStore := userstore.NewMongoStore(appCtx.DB)
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", routes.UserRoutes(userStore)))

	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(v1),
	}
	fmt.Println("Server is running on port 8080")

	_ = server.ListenAndServe()
}
