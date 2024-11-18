package servicestore

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"server/common"
	"server/models"
)

func (m *MongoStore) GetAllService(ctx context.Context) ([]*models.DBService, error) {
	var services []*models.DBService
	coll := m.db.Collection(models.DBService{}.CollectionName())
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, common.ErrDb(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var service models.DBService
		if err := cursor.Decode(&service); err != nil {
			return nil, common.ErrDb(err)
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	return services, nil
}
