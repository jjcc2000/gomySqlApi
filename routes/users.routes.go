package routes

import (
	"encoding/json"
	"net/http"
	"path/mod/db"
	"path/mod/models"

	"github.com/gorilla/mux"
)
func GetUsersHandlers(w http.ResponseWriter, r *http.Request){
	//Creates a Interface
	var users []models.User
	//Find() is a method on the gorm.DB struct that retrieves records from the database based on certain conditions
	db.DB.Find(&users)
	//Encode it trough JSON
	json.NewEncoder(w).Encode(&users)
	
}
func GetUserHandler(w http.ResponseWriter, r *http.Request){
	//Variable to Point to the Database
	var dab models.User
	//Identificas el Id de la Ruta 
	vars := mux.Vars(r)
	//Especified the Database and what you are looking for
	db.DB.First(&dab,vars["id"])
	if dab.ID ==0{
		//Send the code Not Foud trough the header 
		w.WriteHeader(http.StatusNotFound)
		//Response to the frontend 
		w.Write([]byte("The id you are looking for does not Exist"))
		json.NewEncoder(w).Encode(&dab)
	}else{
		json.NewEncoder(w).Encode(&dab)
	}

}
func PostUsersHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUsed:=db.DB.Create(&user)
	err:=createdUsed.Error 

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)//Code 404
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
func DeletedUserHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Delete the User"))
}