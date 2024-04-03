package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"usermanagement.com/authentication"
)

func GenerateInvitationCodeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged in
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

	// Return the generated invitation code in the response
	json.NewEncoder(w).Encode(map[string]string{"message": "Invitation code generated successfully"})
}
