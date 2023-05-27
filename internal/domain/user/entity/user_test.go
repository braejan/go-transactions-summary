package entity_test

import (
	"testing"

	"github.com/braejan/go-transactions-summary/internal/domain/user/entity"
	"github.com/stretchr/testify/assert"
)

// Test_NewUser tests the NewUser function.
func Test_NewUser(t *testing.T) {
	// Create a new user instance.
	user := entity.NewUser(1, "Juana María", "juana.maria@amazingemail.com")
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "Juana María", user.Name)
	assert.Equal(t, "juana.maria@amazingemail.com", user.Email)
}
