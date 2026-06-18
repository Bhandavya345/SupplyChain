package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRole(t *testing.T) {

	role := "Admin"

	assert.Equal(t, "Admin", role)
}
