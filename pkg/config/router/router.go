package router

import (
	"github.com/gorilla/mux"
	"myMongoTest/pkg/config/controller"
)

func Router() *mux.Router {
	// using mux as a router provided
	r := mux.NewRouter()

	// creating some addresses via link, to access db
	r.HandleFunc("/api/tickets", controller.GetAllTickets).Methods("GET")
	r.HandleFunc("/api/match/create", controller.CreateOneMatch).Methods("POST")
	r.HandleFunc("/api/match/{id}", controller.UpdateOneMatch).Methods("PUT")
	r.HandleFunc("/api/match/{id}", controller.DeleteOneMatch).Methods("DELETE")

	r.HandleFunc("/api/match", controller.DeleteAllRecords).Methods("DELETE")

	return r
}
