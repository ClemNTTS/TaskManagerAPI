package taskmanagerAPI

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//Handle request realated with the task list
func TasksRequests(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method{
	//Load all tasks	
	case "GET":
		json.NewEncoder(w).Encode(SortTasks(LoadTasks()))
		log.Println("GET : Load all tasks")

	//Add a new Task and save it	
	case "POST":
		var tache Task
		if err := json.NewDecoder(r.Body).Decode(&tache); err != nil{
			http.Error(w, "Wrong format for your request, check doc and retry", http.StatusNotAcceptable)
			log.Println(err)
			return
		}
		if !IDExist(tache){
			AddTask(tache)
		}else{
			http.Error(w, "A Task with the same ID already exist", http.StatusNotAcceptable)
			return
		}
		json.NewEncoder(w).Encode(SortTasks(LoadTasks()))
		log.Println("POST : New Task")
		
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)	
	}
}

//Handle requests for the task specified with an id in request
func SingleTaskRequests(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	switch r.Method{
	//Load a single Task	
	case "GET":
		GetATask(w,r)
		log.Println("GET : Load a single Task")

	//Modify information for a task
	case "PUT":
		UpdateATask(w,r)
		log.Println("PUT : Modify a Task")
		
	//Delete a specific task	
	case "DELETE":
		DeleteATask(w,r)
		log.Println("DELETE : Delete a Task")
		
	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)	
	}
}

func GetATask(w http.ResponseWriter, r *http.Request){
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	tasks := LoadTasks()

	for _,t := range tasks.All{
		if t.ID == id{
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func UpdateATask(w http.ResponseWriter, r *http.Request){
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	tasks := LoadTasks()
	var new Task
	if err := json.NewDecoder(r.Body).Decode(&new); err != nil{
		http.Error(w, "Wrong format in the request", http.StatusNotAcceptable)
		return
	}


	for i,t := range tasks.All{
		if t.ID == id && id == new.ID{
			tasks.All[i] = new
			WriteTasks(SortTasks(tasks))
			json.NewEncoder(w).Encode(tasks.All[i])
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func DeleteATask(w http.ResponseWriter, r *http.Request){
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	tasks := LoadTasks()

	for i, t := range tasks.All{
		if t.ID == id{
			tasks.All = append(tasks.All[:i],tasks.All[i+1:]...)
			WriteTasks(SortTasks(tasks))
			json.NewEncoder(w).Encode(SortTasks(tasks))
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func IDExist(task Task) bool{
	tasks := LoadTasks()
	for _, t := range tasks.All{
		if t.ID == task.ID{
			return true
		}
	}

	return false
}