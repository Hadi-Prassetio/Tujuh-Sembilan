package main

import (
	"fmt"
	"net/http"
	"test_backend/database"
	"test_backend/pkg/mysql"
	"test_backend/routes"

	"github.com/gorilla/mux"
)

func main() {

	mysql.DataBaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var port = "5000"
	fmt.Println("server running localhost:" + port)

	http.ListenAndServe("localhost:"+port, (r))
}
