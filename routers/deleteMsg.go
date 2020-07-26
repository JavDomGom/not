package routers

import (
	"net/http"

	"github.com/JavierDominguezGomez/not/db"
)

/*DeleteMsg Delete a specific message. */
func DeleteMsg(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the id parameter.", http.StatusBadRequest)
		return
	}

	err := db.DeleteMsg(ID, IDUser)
	if err != nil {
		http.Error(w, "An error occurred while trying to delete the message. "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
