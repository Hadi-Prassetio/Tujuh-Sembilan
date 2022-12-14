package routes

import (
	"test_backend/handlers"
	"test_backend/pkg/mysql"
	"test_backend/repositories"

	"github.com/gorilla/mux"
)

func AccountRoutes(r *mux.Router) {
	accountRepository := repositories.RepositoryAccount(mysql.DB)
	h := handlers.HandlerAccount(accountRepository)

	r.HandleFunc("/account", h.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts", h.FindAccounts).Methods("GET")
}
