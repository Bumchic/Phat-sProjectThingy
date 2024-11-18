package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const EntityName = "services"

func (DBService) CollectionName() string {
	return EntityName
}

type DBService struct {
	ServiceId   primitive.ObjectID `bson:"_id"`
	ServiceName string             `bson:"service_name"`
	Price       float64            `bson:"price"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type CreateService struct {
	ServiceId   primitive.ObjectID `bson:"_id"`
	ServiceName string             `bson:"service_name" json:"service_name"`
	Price       float64            `bson:"price" json:"price"`
	CreatedAt   time.Time          `bson:"created_at"`
}

type UpdateService struct {
	ServiceId   primitive.ObjectID `bson:"_id" json:"_id"`
	ServiceName string             `bson:"service_name" json:"service_name"`
	Price       float64            `bson:"price" json:"price"`
}
