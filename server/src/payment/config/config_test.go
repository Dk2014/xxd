package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	if err := Load("./config.tpl.json"); err != nil {
		t.Error(err)
	}
	cnnStr := "supei:123456@tcp(127.0.0.1:3306)/payment"
	if CFG.Database.ToConnectStr() != cnnStr {
		t.Errorf("except %s; get %s", cnnStr, CFG.Database.ToConnectStr())
	}

}
