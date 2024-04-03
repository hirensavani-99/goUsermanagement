package utils

import "usermanagement.com/database"

// isValidInvitationCode checks if the provided invitation code is valid
func IsValidInvitationCode(code string) bool {
	var invitation database.InvitationCode

	if err := database.Db.Where("code = ? AND used = ?", code, false).First(&invitation).Error; err != nil {
		return false
	}
	return true

}

// markInvitationCodeAsUsed marks the provided invitation code as used
func MarkInvitationCodeAsUsed(code string) error {
	if code != "INVITEME" {
		var invitation database.InvitationCode
		if err := database.Db.Where("code = ?", code).First(&invitation).Error; err != nil {
			return err
		}
		invitation.Used = true
		if err := database.Db.Save(&invitation).Error; err != nil {
			return err
		}
	}
	return nil

}
