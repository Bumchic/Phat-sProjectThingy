package userstore

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId    primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"username"`            
	Email     string             `bson:"email"`                
	Password  string             `bson:"password"`             
	Credits   float64            `bson:"credits"`              
	CreatedAt time.Time          `bson:"created_at,omitempty"` 
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

type UserStore interface{
	CreateUser(newUser *User) error
}


func (m *MongoStore) CreateUser(newUser *User) error{
	coll := m.db.Client().Database("database").Collection("users")
	existingUser, err := m.FindUser(newUser.Email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("user already exists")
	}
	_, err = coll.InsertOne(context.TODO(), newUser)
	return err
} 