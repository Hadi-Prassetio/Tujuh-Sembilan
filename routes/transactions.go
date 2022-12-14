package routes

import (
	"test_backend/handlers"
	"test_backend/pkg/mysql"
	"test_backend/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transactions-account", h.FindAccountTransactions).Methods("GET")
}
