package inmemory

import (
	"testing"
)

var storeKey = "active-tabs"
var storeKeyFail = "active"
var storeValue = "getir"

func TestDB_Write(t *testing.T) {
	db := New()
	db.Write(storeKey, storeValue)
	value, err := db.Read(storeKey)
	if err != nil {
		t.Errorf("Status Failed.  Error: %s", err.Error())
	}
	if value != storeValue {
		t.Errorf("Status Failed.  Got: %s, Want: %s", value, storeValue)
	}
}

func TestDB_ReadKey(t *testing.T) {
	db := New()
	db.Write(storeKey, storeValue)
	value, err := db.Read(storeKey)
	if err != nil {
		t.Errorf("Status Failed.  Error: %s", err.Error())
	}
	if value != storeValue {
		t.Errorf("Status Failed.  Got: %s, Want: %s", value, storeValue)
	}
}

func TestDB_ReadKey_Failed(t *testing.T) {
	db := New()
	db.Write(storeKey, storeValue)
	_, err := db.Read(storeKeyFail)
	if err == nil {
		t.Errorf("Status Failed. Expected error but not found.")
	}
}
