package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"usermanagement.com/authentication"
	"usermanagement.com/database"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get username from JWT token
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Parse and validate JWT token
	token, err := jwt.ParseWithClaims(tokenString, &authentication.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return authentication.JwtKey, nil
	})
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Check if token is valid
	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Extract username from token claims
	claims, ok := token.Claims.(*authentication.Claims)
	if !ok {
		http.Error(w, "Failed to extract token claims", http.StatusInternalServerError)
		return
	}

	// Decode request body
	var updatedUser database.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user in the database
	if err := database.Db.Model(&database.User{}).Where("username = ?", claims.Username).Updates(updatedUser).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}
