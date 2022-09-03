package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	UserReposity := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(UserReposity)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user-profile",middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}