package taskmanagerAPI

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
}

type Tasks struct{
	All []Task `json:"tasks"`
}

func LoadTasks() Tasks {
	result := Tasks{}

	//Open the tasks file
	file, err := os.Open("./data/tasks.json")
	if err != nil{
		return result
	}
	defer file.Close()

	//Reading json file
	result_decoded,err := io.ReadAll(file)
	if err != nil{
		log.Fatal("Error at reading file : " + err.Error())
	}

	//Reading JSON
	err = json.Unmarshal(result_decoded, &result)
	if err != nil{
		log.Fatal("Error at parsing JSON : "+err.Error())
	}

	return result
}

func WriteTasks(liste Tasks){
	encoded, err := json.Marshal(liste)
	if err != nil{
		log.Fatal("Error at parse data : "+err.Error())
	}

	err = os.WriteFile("./data/tasks.json", encoded, 0666)
	if err != nil{
		log.Fatal("Error writing to file : ", err)
	}
}

func AddTask(new Task){
	tasks := LoadTasks()
	tasks.All = append(tasks.All,new)
	WriteTasks(tasks)
}