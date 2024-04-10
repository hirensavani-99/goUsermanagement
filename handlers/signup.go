package handlers

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"usermanagement.com/database"
	"usermanagement.com/utils"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate invitation code
	invitationCode := r.Header.Get("Invitation-Code")
	if invitationCode != "INVITEME" {
		if !utils.IsValidInvitationCode(invitationCode) {
			http.Error(w, "Invalid or already used invitation code", http.StatusBadRequest)
			return
		}
	}

	// Hash password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Create user in database
	if err := database.Db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mark invitation code as used
	if err := utils.MarkInvitationCodeAsUsed(invitationCode); err != nil {
		http.Error(w, "Error marking invitation code as used", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
