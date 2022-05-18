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

func TestAddNewTigerSighting(t *testing.T) {
	t.Run("when adding new tiger sighting", func(t *testing.T) {
		defer resetDB()
		_, tigerObj := AddNewTiger(postgresDB, "test1", "2001-01-02", 1.9001918, 2.9277827,
			1652819764)
		statusMsg := AddNewTigerSighting(postgresDB, tigerObj.ID, 1.9001213, 2.9277789,
			1652819764)
		assert.Equal(t, "Tiger Sighting Logged", statusMsg)
	})
	t.Run("when adding new tiger sighting to non existent tiger", func(t *testing.T) {
		defer resetDB()
		statusMsg := AddNewTigerSighting(postgresDB, 999, 1.9001918, 2.9277827,
			1652819764)
		assert.Equal(t, "Invalid Tiger ID 999", statusMsg)
	})
}

func TestGetAllTigers(t *testing.T) {
	t.Run("when no tigers are Added", func(t *testing.T) {
		defer resetDB()
		data := GetAllTigers(postgresDB, c)
		assert.Equal(t, 0, len(data))
	})
	t.Run("when tigers are Added, Sorted by last seen desc", func(t *testing.T) {
		defer resetDB()
		AddNewTiger(postgresDB, "First", "2001-01-02", 1.9001918, 2.9277827,
			1652819764)
		AddNewTiger(postgresDB, "Second", "2001-01-02", 1.9001918, 2.9277827,
			1652819765)
		data := GetAllTigers(postgresDB, c)
		assert.Equal(t, 2, len(data))
		assert.Equal(t, "Second", data[0].TigerName)

	})
}
