package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"usermanagement.com/database"
)

func main() {
	// Initialize database
	database.InitDB()
	defer database.Db.Close()

	// Initialize router
	router := mux.NewRouter()

	// Start server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
