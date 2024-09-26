package taskmanagerAPI

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Launch() {
	//Set port
	port := "8080"
	if len(os.Args) == 2 {
		if len(os.Args[1]) != 4{
			log.Fatal("Bad port entered in arguments.\nExample : go run main.go 3000")
		}else{
			port = os.Args[1]
		}
	}

	mux := http.NewServeMux()

	http.HandleFunc("/tasks",TasksRequests)
	http.HandleFunc("/tasks/",SingleTaskRequests)

	server := &http.Server{
		Addr: port,
		Handler: mux,
		ReadHeaderTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 120*time.Second,
		MaxHeaderBytes: 1<<20,
	}

	fmt.Printf("Listening on http://localhost:%v\n",port)
	log.Fatal(server.ListenAndServe())

}