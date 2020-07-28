package routers

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
	"github.com/JavierDominguezGomez/not/models"
)

/*DeleteFollow Delete a specific relation. */
func DeleteFollow(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Follow
	t.UserID = IDUser
	t.UserFollowedID = ID

	status, err := db.DeleteFollow(t)
	if err != nil {
		http.Error(w, "Failed to delete relation from database! "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "Relation couldn't be deleted: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
