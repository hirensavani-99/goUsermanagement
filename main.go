package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"usermanagement.com/database"
	"usermanagement.com/handlers"
)

func main() {

	database.InitDB()
	defer database.Db.Close()

	// Initialize router
	router := mux.NewRouter()

	// Define API endpoints
	router.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/generate-invite-code", handlers.GenerateInvitationCodeHandler).Methods("POST")
	router.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
	router.HandleFunc("/user/delete", handlers.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/user/update", handlers.UpdateUserHandler).Methods("PUT")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("There we are !")
		fmt.Fprintf(w, "Hello, welcome to the greeting endpoint!")
	}).Methods("GET")

	// Start server
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
