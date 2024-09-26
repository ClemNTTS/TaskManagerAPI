package taskmanagerAPI

import (
	"encoding/json"
	"log"
	"net/http"
)

//Handle request realated with the task list
func TasksRequests(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method{
	//Load all tasks	
	case "GET":
		json.NewEncoder(w).Encode(LoadTasks())

	//Add a new Task and save it	
	case "POST":
		var tache Task
		if err := json.NewDecoder(r.Body).Decode(&tache); err != nil{
			log.Fatal(err)
		}
		AddTask(tache)
		json.NewEncoder(w).Encode(LoadTasks())
		
	default:
		http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)	
	}
}

//Handle requests for the task specified with an id in request
func SingleTaskRequests(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method{
	//Load a single Task	
	case "GET":
		json.NewEncoder(w).Encode(LoadTasks())

	//Add a new Task and save it	
	case "POST":
		var tache Task
		if err := json.NewDecoder(r.Body).Decode(&tache); err != nil{
			log.Fatal(err)
		}
		AddTask(tache)
		json.NewEncoder(w).Encode(LoadTasks())
		
	default:
		http.Error(w, "Méthode non supportée", http.StatusMethodNotAllowed)	
	}
}
