package userstore


import "go.mongodb.org/mongo-driver/mongo"

type MongoStore struct{
	db *mongo.Database
}

func NewMongoStore(db *mongo.Database) *MongoStore{
	return &MongoStore{
		db: db,
	}
}