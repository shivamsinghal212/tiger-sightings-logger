package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNewTiger(t *testing.T) {
	t.Run("when adding a new tiger", func(t *testing.T) {
		defer resetDB()
		statusMsg, _ := AddNewTiger(postgresDB, "test1", "2001-01-02", 1.9001918, 2.9277827,
			1652819764)
		assert.Equal(t, statusMsg, "New Entry Created")
	})
	t.Run("when adding an exisiting tiger", func(t *testing.T) {
		defer resetDB()
		AddNewTiger(postgresDB, "test1", "2001-01-02", 1.9001918, 2.9277827,
			1652819764)
		statusMsg, tiger := AddNewTiger(postgresDB, "test1", "2001-01-02", 1.9001918, 2.9277827,
			1652819764)
		assert.Equal(t, statusMsg, fmt.Sprintf("%s already exists.", tiger.Name))
	})
	t.Run("when adding incorrect DOB for tiger", func(t *testing.T) {
		defer resetDB()
		statusMsg, _ := AddNewTiger(postgresDB, "test1", "20-01-2002", 1.9001918, 2.9277827,
			1652819764)
		assert.Equal(t, statusMsg, "Incorrect Date Format")
	})
}