package user

import (
	"encoding/json"
	"net/http"
	"server/storage/userstore"
)

type Handler struct{
	store userstore.MongoStore
}


func NewHandler(store userstore.MongoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request){
	var user User 
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil { 
		http.Error(w, "Invalid input: "+ err.Error(), http.StatusBadRequest)
		return 
	}
	err, newUser := CreateNewUser(&user)
	if err != nil{
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
	}
	//convert  from user to storeUser
	userStoreUser := UserToStoreUser(newUser)

	if err := h.store.CreateUser(&userStoreUser); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&newUser)
	if err != nil {
		http.Error(w, "Failed to encode response: "+ err.Error(), http.StatusInternalServerError) 
	}
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request){

}