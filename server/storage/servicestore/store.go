package servicestore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"server/models"
)

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *MongoStore {
	return &MongoStore{
		db: db,
	}
}

type ServiceStore interface {
	CreateService(ctx context.Context, newService *models.DBService) error
	GetAllService(ctx context.Context) ([]*models.DBService, error)
	UpdateService(ctx context.Context, service *models.UpdateService) error
}
