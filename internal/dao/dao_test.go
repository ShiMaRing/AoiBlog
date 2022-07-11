package dao

import "testing"

func TestCreateTables(t *testing.T) {
	t.Fatalf("dont do this again!")
	err := CreateTables()
	if err != nil {
		t.Error(err)
	}
}
