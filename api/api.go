package api

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/andrepinto/sherlock/config"
	"fmt"
	"encoding/json"
	"github.com/andrepinto/sherlock/system"
)

type Api struct {
	Port int
	Configuration *config.Configuration
	router *mux.Router
}

type StatusResponse struct {
	Status string `json:"status"`
}

type AllResponse struct {
	Service *config.Service `json:"service,omitempty"`
	Dependencies *[]config.Dependency  `json:"dependencies,omitempty"`
}

type SystemResponse struct {
	System *system.SystemInfo `json:"system,omitempty"`
	Memory *system.MemoryInfo `json:"memory,omitempty"`
}

func NewApi(config *config.Configuration) (*Api, error){
	return &Api{
		Port:config.Api.Port,
		Configuration: config,
		router: mux.NewRouter(),
	}, nil
}

func(api *Api) serviceHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(api.Configuration.Service)
}

func(api *Api) statusHandler(w http.ResponseWriter, r *http.Request) {
	status := StatusResponse{"ok"}
	json.NewEncoder(w).Encode(status)
}

func(api *Api) allHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&AllResponse{
		Service: &api.Configuration.Service,
		Dependencies: &api.Configuration.Dependencies,
	})
}

func(api *Api) systemHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		&SystemResponse{
			System: system.GetSystemInfo(),
			Memory: system.GetMemoryInfo(),
		})
}

func(api *Api) Run() {

	api.router.HandleFunc("/service", api.serviceHandler)
	api.router.HandleFunc("/status", api.statusHandler)
	api.router.HandleFunc("/info", api.allHandler)
	api.router.HandleFunc("/system", api.systemHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",api.Port), api.router))
}
