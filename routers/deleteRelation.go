package routers

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*DeleteRelation Delete a specific relation. */
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := db.DeleteRelation(t)
	if err != nil {
		http.Error(w, "Failed to delete relation from database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Relation couldn't be deleted: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
