package gift_code

import (
	"core/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync/atomic"
)

var (
	g_db         *sql.DB
	g_MaxVersion int64
)

func Init(connectString string) error {
	if g_db != nil {
		panic("dup init")
	}
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	initMaxVersion(db)
	g_db = db
	return nil
}

func initMaxVersion(db *sql.DB) {
	row, err := db.Query("select MAX(`version`) from `gift_code`")
	if err != nil {
		log.Errorf("can not get current max version. err [%v]", err)
		panic(err)
	}
	for row.Next() {
		row.Scan(&g_MaxVersion)
	}
	row.Close()
}

func GetVersion() int64 {
	return atomic.AddInt64(&g_MaxVersion, 1)
}

type GiftCode struct {
	Code            string   `json:"code"`
	Type            CodeType `json:"type"`
	ItemId          int16    `json:"item_id"`
	ServerId        int      `json:"item_id"`
	EffectTimestamp int64    `json:"effect_timestamp"`
	ExpireTimestamp int64    `json:"expire_timestamp"`
	//EffectTimestamp int64 `json:"_"`
	//ExpireTimestamp int64 `json:"_"`
	Version int64  `json:"version"`
	Disable int8   `json:"disable"`
	Content string `json:"content"`
	Config  string `json:"config"`
}

type ResultData struct {
	Gift     []GiftCode `json:"gift_code"`
	Totality int64      `json:"totality"`
}

type ComputeTotality struct {
	Version  int64
	Totality int64
}

func (_ GiftCode) InsertSQL() string {
	return "INSERT IGNORE INTO `gift_code` " +
		"( `code`, `type`, `server_id`, `effect_timestamp`, `expire_timestamp`, `item_id`,`version`,`content`, `config`) " +
		" values (?, ?, ?, ?, ?, ?, ?, ?, ?) "
}

func InsertCode(code GiftCode) (bool, error) {
	sql := code.InsertSQL()

	result, err := g_db.Exec(sql, code.Code, code.Type, code.ServerId, code.EffectTimestamp, code.ExpireTimestamp, code.ItemId, code.Version, code.Content, code.Config)
	if err != nil {
		log.Errorf("InsertCode Exec error [%v] arg [%v]\n", err, code)
		return false, err
	}
	effectRows, err := result.RowsAffected()
	if err != nil {
		log.Errorf("InsertCode RowsAffected erro [%v] arg [%v] effectRows [%v]\n", err, code, effectRows)
		return false, err
	}
	return effectRows > 0, nil
}

func CancelCode(version int64, serverId int) error {
	var sql string
	if serverId > 0 {
		sql = fmt.Sprintf("update gift_code set `disable`=1 where `version`='%d' and `server_id`='%d'", version, serverId)
	} else {
		sql = fmt.Sprintf("update gift_code set `disable`=1 where `version`='%d' ", version)
	}
	_, err := g_db.Exec(sql)
	if err != nil {
		log.Errorf("CancelCode  erro [%v] sql [%v] effectRows [%v]\n", err, sql)
		return err
	}
	return nil
}

//
func QueryCode(version int64, serverId int, offset int, limit int) (*ResultData, error) {
	var sql string
	var sql_count string
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 20
	}
	if serverId > 0 {
		sql_count = fmt.Sprintf("select count(*) `total` from gift_code  where `version`=%d and `server_id`=%d", version, serverId)
	} else {
		sql_count = fmt.Sprintf("select count(*) `total` from gift_code  where `version`=%d", version)
	}
	rows_count, err := g_db.Query(sql_count)
	if err != nil {
		log.Errorf("QueryCode  error [%v] sql [%v]\n", err, sql_count)
		return nil, err
	}
	var total int64
	for rows_count.Next() {
		rows_count.Scan(&total)
	}
	if serverId > 0 {
		sql = fmt.Sprintf("select `code`, `type`, `item_id`, `effect_timestamp`, `expire_timestamp`, `version`, `disable`, `content`, `config` from gift_code  where `version`='%d' and `server_id`= %d limit %d offset %d ", version, serverId, limit, offset)
	} else {
		sql = fmt.Sprintf("select `code`, `type`, `item_id`, `effect_timestamp`, `expire_timestamp`, `version`, `disable`, `content`, `config` from gift_code  where `version`='%d' limit %d offset %d ", version, limit, offset)
	}
	rows, err := g_db.Query(sql)
	if err != nil {
		log.Errorf("QueryCode  error [%v] sql [%v]\n", err, sql)
		return nil, err
	}
	var results ResultData
	for rows.Next() {
		var code GiftCode
		rows.Scan(&code.Code, &code.Type, &code.ItemId, &code.EffectTimestamp, &code.ExpireTimestamp, &code.Version, &code.Disable, &code.Content, &code.Config)
		results.Gift = append(results.Gift, code)
	}
	results.Totality = total
	return &results, nil
}
