package servicestore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"server/common"
	"server/models"
)

func (m *MongoStore) CreateService(ctx context.Context, service *models.DBService) error {

	coll := m.db.Collection(service.CollectionName())
	one, err := coll.InsertOne(ctx, service)
	if one != nil {
		id := (one.InsertedID).(primitive.ObjectID)
		service.ServiceId = id
		fmt.Println("Service")
	}
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
