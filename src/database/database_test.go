package database

import "testing"

func TestConnection(t *testing.T) {
	db := GetInstance()

	err := db.Ping()
	if err != nil {
		t.Fatalf(`connectDataBase() dont connected to database`)
	}
}
