package routes

import (
	"encoding/json"
	"net/http"
	"path/mod/db"
	"path/mod/models"
	"github.com/gorilla/mux"
)

func GeTasksHandler(w http.ResponseWriter, r *http.Request){
	var Task []models.Task
	db.DB.Find(&Task)
	json.NewEncoder(w).Encode(&Task)
}

func CreateTaskHansddler(w http.ResponseWriter, r *http.Request){
	var TaskRecieved models.Task
	json.NewDecoder(r.Body).Decode(&TaskRecieved)
	checker:=db.DB.Create(&TaskRecieved)	
	err:= checker.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The Post Method Failed"))
	}
	json.NewEncoder(w).Encode(&TaskRecieved)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	var Task models.Task
	vars := mux.Vars(r)
	db.DB.First(&Task,vars["id"])
	if Task.UserId==0{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The Task was not found"))
		return
	}
	db.DB.Unscoped().Delete(&Task)
	w.Write([]byte("The task has been deleted"))
	w.WriteHeader(http.StatusNoContent)
}
func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	var Task models.Task
	vars := mux.Vars(r)
	db.DB.Find(&Task,vars["id"])

	if Task.UserId==0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("That id does not exists"))
		return
	}else{
	json.NewEncoder(w).Encode(&Task)
	}
}