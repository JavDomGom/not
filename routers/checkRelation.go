package routers

import (
	"encoding/json"
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*CheckRelation Check for relation beetween two users in database. */
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var response models.ResponseCheckRelation

	status, err := db.CheckRelation(t)
	if err != nil || !status {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
