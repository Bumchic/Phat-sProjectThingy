package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"server/controllers"
	"server/storage/servicestore"
)

func ServicesRoutes(db *mongo.Database) http.Handler {

	store := servicestore.NewMongoStore(db)
	c := controllers.NewServiceHandler(store)

	router := http.NewServeMux()

	router.HandleFunc("POST  /service", c.CreateServiceHandler)
	router.HandleFunc("GET /service", c.GetAllServicesHandler)
	router.HandleFunc("PUT /service", c.UpdateServiceHandler)

	return router

}
