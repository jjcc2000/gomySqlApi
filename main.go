package main

import (
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"path/mod/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConection()
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})
	
	r := mux.NewRouter()	

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/SecondPage", routes.SecondPage)
	
	r.HandleFunc("/users",routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users",routes.GetUsersHandlers).Methods("GET")

	r.HandleFunc("/users/{id}",routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}",routes.DeletedUserHandler).Methods("DELETE")	

	http.ListenAndServe(":8080", r)
}
