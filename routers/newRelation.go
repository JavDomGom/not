package routers

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*NewRelation register a new relation between two users, */
func NewRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.InsertRealtion(t)
	if err != nil {
		http.Error(w, "Failed to save relation to database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Relation couldn't be saved: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
