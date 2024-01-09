// backend/database/database_test.go

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	db, err := InitDB()
	defer db.Close()

	// Check if there's no error during database initialization
	assert.NoError(t, err)
	assert.NotNil(t, DB)
}
