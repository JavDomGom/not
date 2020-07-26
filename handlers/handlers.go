package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/JavierDominguezGomez/not/middleware"
	"github.com/JavierDominguezGomez/not/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Handlers set port, handler, listen and serve the HTTP server.*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.ValidateJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.CheckDB(middleware.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/msg", middleware.CheckDB(middleware.ValidateJWT(routers.RecordMsg))).Methods("POST")
	router.HandleFunc("/readMsg", middleware.CheckDB(middleware.ValidateJWT(routers.ReadMsg))).Methods("GET")
	router.HandleFunc("/deleteMsg", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteMsg))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middleware.CheckDB(middleware.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.CheckDB(middleware.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/newRelation", middleware.CheckDB(middleware.ValidateJWT(routers.NewRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middleware.CheckDB(middleware.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/checkRelation", middleware.CheckDB(middleware.ValidateJWT(routers.CheckRelation))).Methods("GET")

	router.HandleFunc("/getUsers", middleware.CheckDB(middleware.ValidateJWT(routers.GetUsers))).Methods("GET")
	router.HandleFunc("/getFollowersMsg", middleware.CheckDB(middleware.ValidateJWT(routers.GetFollowersMsg))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
