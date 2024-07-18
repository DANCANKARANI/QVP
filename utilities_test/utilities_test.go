// Validate a correctly formatted email address
package utilities_test

import (
	"testing"
	"github.com/DANCANKARANI/QVP/utilities"
	"github.com/stretchr/testify/assert"
)
	
	func TestValidateEmail_ValidEmail(t *testing.T) {
		email := "test@example.com"
		result, err := utilities.ValidateEmail(email)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, email, *result)
	}
	
	// Handle an empty email string
	func TestValidateEmail_EmptyEmail(t *testing.T) {
		email := ""
		result, err := utilities.ValidateEmail(email)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "your email  is not valid")
	}
	