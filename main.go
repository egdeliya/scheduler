package main

import (
	"fmt"
	"log"
	"net/http"
	"scheduler/api"
	"scheduler/service"
	"scheduler/startup"

	"github.com/gorilla/mux"
)

func main() {
	cfg := startup.LoadConfig()
	resourceAssigner := service.NewResourceAssigner()
	resourceApi := api.NewResourceApi(resourceAssigner)

	r := mux.NewRouter()
	r.HandleFunc("/resources", resourceApi.ResourcesHandler).Methods(http.MethodPost)
	fmt.Println("server started on port ", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
