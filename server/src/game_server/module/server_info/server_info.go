package server_info

import (
	"core/fail"
	"core/mysql"
	. "game_server/config"
)

var version []byte

func init() {
	RefreshServerVersion()
}

func RefreshServerVersion() {
	db, err1 := mysql.Connect(GetDBConfig())
	fail.When(err1 != nil, err1)

	defer db.Close()
	res, err := db.ExecuteFetch([]byte(`select * from server_info`), -1)
	fail.When(err != nil, err)

	iVersion := res.Map("version")

	for _, row := range res.Rows {
		version = row.Bin(iVersion)
	}
}
