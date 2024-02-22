package validator

import (
	"fmt"
	"time"
)

const (
	// Define Limit and Offset for the API
	MaxLimit  = 100
	MaxOffset = 100
)

// Function to validate UserId
func ValidateId(userId int64) error {
	// Assuming the UserId should be a positive integer
	if userId <= 0 {
		return fmt.Errorf("UserId must be a positive integer")
	}
	return nil
}

// ValidateDuration checks if the given string can be parsed as a duration
func ValidateDuration(durationStr string) error {
	_, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	return nil
}

// ValidateDuration checks if the given string can be parsed as a duration
// Placeholder for later validation
// TODO: Implement this function when needed
func ValidateTokenType(tokenType string) error {
	fmt.Printf("Token Type: %s\n", tokenType)
	if tokenType != "access" {
		return fmt.Errorf("invalid token type")
	}
	return nil
}

// Function to validate offset parameters
func ValidateLimit(limit int32) error {
	// Define a reasonable range for limit (e.g.,  1-100)
	if limit <  1 || limit >  MaxLimit {
		return fmt.Errorf("limit must be between  1 and  %d", MaxLimit)
	}

	return nil
}

// Function to validate limit parameters
func ValidateOffset(offset int32) error {
	// Offset should not be negative
	if offset <  0 || MaxOffset >  100{
		return fmt.Errorf("offset must be non-negative and less than %d", MaxOffset)
	}

	return nil
}