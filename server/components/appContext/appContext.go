package appContext

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AppContext có các thành phần như client MongoDB và logger
type AppContext struct {
	DBClient *mongo.Client
	DB *mongo.Database
	Logger   *log.Logger
}

var once sync.Once
var appCtx *AppContext

// khởi tạo AppContext một lần duy nhất (singleton)
func InitAppContext(uri string) *AppContext {
	once.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}

		db := client.Database("database")

		appCtx = &AppContext{
			DBClient: client,
			DB: db,
			Logger:   log.Default(),
		}
	})
	return appCtx
}

// đóng kết nối cơ sở dữ liệu khi ứng dụng kết thúc
func (ctx *AppContext) CloseDB() {
	if err := ctx.DBClient.Disconnect(context.TODO()); err != nil {
		ctx.Logger.Printf("Error disconnecting from MongoDB: %v", err)
	} else {
		ctx.Logger.Println("Disconnected from MongoDB successfully.")
	}
}
