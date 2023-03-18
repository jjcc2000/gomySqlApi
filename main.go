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
	////Users Routes/////
	r.HandleFunc("/users",routes.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users",routes.GetUsersHandlers).Methods("GET")
	r.HandleFunc("/users/{id}",routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}",routes.DeletedUserHandler).Methods("DELETE")	
	////Task Routes////
	r.HandleFunc("/tasks",routes.GeTasksHandler).Methods("GET")
	r.HandleFunc("/tasks",routes.CreateTaskHansddler).Methods("POST")
	r.HandleFunc("/tasks/{id}",routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}",routes.GetTaskHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
