package handlers

import (
	"net/http"
	"strconv"

	"usermanagement.com/database"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request parameters
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve user from the database
	var user database.User
	if err := database.Db.First(&user, userID).Error; err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Check if user exists
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete the user
	if err := database.Db.Delete(&user).Error; err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
