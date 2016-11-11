package database

import (
	"testing"
)

func TestInit(t *testing.T) {
	cnnStr := "supei:123456@tcp(127.0.0.1:3306)/payment"
	Init(cnnStr)
	if err := g_db.Ping(); err != nil {
		t.Error(err)
	}
}
