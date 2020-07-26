package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*Register Function to register a user in database.*/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the received data."+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is a required data.", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "You must specify a password of at least 6 characters.", 400)
		return
	}

	_, found, _ := db.CheckIfUserAlreadyExists(t.Email)
	if found {
		http.Error(w, "There is already a registered user with this email.", 400)
		return
	}

	_, status, err := db.InsertUserRegister(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register the user: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "User could not be saved in database: ", 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
