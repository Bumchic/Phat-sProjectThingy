package servicestore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"server/common"
	"server/models"
)

func (m *MongoStore) UpdateService(ctx context.Context, service *models.UpdateService) error {
	coll := m.db.Collection(models.DBService{}.CollectionName())
	filter := bson.M{"_id": service.ServiceId}
	update := bson.M{"$set": service}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
