package controllers

import (
	"fmt"
	"regexp"
)

func CheckPasswordStrength(password string) (bool, error) {
	// Regular expressions for password strength
	var (
		uppercaseRegex = regexp.MustCompile(`[A-Z]`)
		lowercaseRegex = regexp.MustCompile(`[a-z]`)
		digitRegex     = regexp.MustCompile(`[0-9]`)
		specialRegex   = regexp.MustCompile(`[^a-zA-Z0-9]`)
	)

	// Check if password length is at least 8 characters
	if len(password) < 8 {
		return false, fmt.Errorf("password must be at least 8 characters long")
	}

	// Check if password contains at least one uppercase letter
	if !uppercaseRegex.MatchString(password) {
		return false, fmt.Errorf("password must contain at least one uppercase letter")
	}

	// Check if password contains at least one lowercase letter
	if !lowercaseRegex.MatchString(password) {
		return false, fmt.Errorf("password must contain at least one lowercase letter")
	}

	// Check if password contains at least one digit
	if !digitRegex.MatchString(password) {
		return false, fmt.Errorf("password must contain at least one digit")
	}

	// Check if password contains at least one special character
	if !specialRegex.MatchString(password) {
		return false, fmt.Errorf("password must contain at least one special character")
	}

	// Password meets all criteria
	return true, nil
}
