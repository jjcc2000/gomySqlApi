package routes

import (
	"encoding/json"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World "))
}
func SecondPage(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Should be in the SECOND PAGE")
}