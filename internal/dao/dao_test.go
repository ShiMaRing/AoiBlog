package dao

import "testing"

func TestCreateTables(t *testing.T) {
	err := CreateTables()
	if err != nil {
		t.Error(err)
	}
}
