package user

import (
	"log"
	"server/storage/userstore"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) string{
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("Error hashing password")
	}
	return string(encryptedPassword)
}
func UserToStoreUser(u User) userstore.User { 
	return userstore.User{ 
		UserId: u.UserId,
		UserName: u.UserName,
		Email: u.Email,
		Password:
		u.Password,
		Credits: u.Credits,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func CreateNewUser(newUser *User) (error, User){
	return nil, User{
		UserId: primitive.NewObjectID(),
		UserName: newUser.UserName,
		Password: hashPassword(newUser.Password),
		Email: newUser.Email,
		Credits: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
