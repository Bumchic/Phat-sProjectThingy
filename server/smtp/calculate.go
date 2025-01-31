package smtp

import (
	"fmt"
	"sync"
	"time"

	"server/models"
	"server/storage/billstore"
	"server/storage/userstore"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateNewBill(total float32, userId primitive.ObjectID) *models.Bill{
	return &models.Bill{
		BillId: primitive.NewObjectID(),
		BillName: "Payment",
		UserId: userId,
		Price: float64(total),
		Paid: true,
		Expired: time.Now().Add(time.Hour * 72),
		CreatedAt: time.Now(),
	}
}

func CalculateUserBill(store *MongoStore) {
	userDB := userstore.NewMongoStore(store.db)
	billDB := billstore.NewMongoStore(store.db)
	var wg sync.WaitGroup
	users, err := userDB.FindAll()
	if err != nil {
		fmt.Errorf("Error getting users" + err.Error())
	}
	for _, user := range users {
		var total float32
		if len(user.ServiceIds) == 0{
			continue
		}

		wg.Add(1)
		go func(){
			defer wg.Done()
			//Calculate user service totals
			for _, service :=  range user.ServiceIds{
				total += service.Price
			}
			newBill := CreateNewBill(total, user.UserId)
			err := userDB.UpdateUserBill(user.UserName, *newBill)
			if err != nil{
				fmt.Errorf("Error updating user %s bill", newBill.UserId)
			}
			err = billDB.CreateBill(*newBill)
			if err != nil{
				fmt.Errorf("Error creating bill for %s", newBill.UserId)
			}
			//Send mail to the user
			err = SendMail([]string{user.Email}, total)
			if err != nil{
				fmt.Errorf("Error sending mail for %s", newBill.UserId)
			}
		}()
	}
	wg.Wait()
}
