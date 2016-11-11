package giftcode

import (
	"config_hub/args"
	"config_hub/httpsrv"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"regexp"
)

var mysql_db *sql.DB

func Load() {
	conf := args.Conf.GiftCode

	if conf == nil {
		return
	}

	mysql_db, err := sql.Open("mysql", conf.SqlUser+":"+conf.SqlPass+"@tcp("+conf.SqlAddr+":"+conf.SqlPort+")/"+conf.SqlName)
	if err != nil {
		panic(err)
	}

	httpsrv.RegistAction("/giftcode/", &giftCardAct{
		DB: mysql_db,
	})
}

func Unload() {
	if mysql_db != nil {
		mysql_db.Close()
	}
}

// serve giftcard/srvid_timestamp.txt
type giftCardAct struct {
	DB      *sql.DB
	pattern *regexp.Regexp
}

func (this *giftCardAct) getPattern() *regexp.Regexp {
	if this.pattern == nil {
		// this will match srvID_TIMESTAMP.txt, and get id and timestamp optionally as submatches
		this.pattern = regexp.MustCompile("srv([0-9]+)(\\_([0-9]+))?\\.txt")
	}
	return this.pattern
}

func (this *giftCardAct) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// match out the server id
	match := this.getPattern().FindStringSubmatch(req.URL.Path)
	if len(match) < 2 {
		http.Error(resp, "url didn't match the format", http.StatusNotFound)
		return
	}
	srv_id := match[1]
	match_time := false
	timestamp := ""

	// check if provided the timestamp for querying
	if len(match) >= 4 && len(match[2]) > 0 {
		match_time = true
		timestamp = match[3]
	}

	// query database
	sql_str := "SELECT `code`, `type`, `effect_timestamp`, `expire_timestamp`, `item_id`, `version` FROM `gift_code` WHERE `server_id`=? "
	var row *sql.Rows
	var err error
	if match_time {
		sql_str += "and `expire_timestamp`>=?;"
		row, err = this.DB.Query(sql_str, srv_id, timestamp)
	} else {
		sql_str += ";"
		row, err = this.DB.Query(sql_str, srv_id)
	}
	log.Println(sql_str)
	if err != nil {
		log.Println(err)
		http.Error(resp, "db query statement failed", http.StatusInternalServerError)
		return
	}
	defer row.Close()

	// this is the json object in golang
	code_list := map[string]*GiftCodeInfo{}

	// fill it
	for row.Next() {
		code_info := &GiftCodeInfo{}
		var code string
		err := row.Scan(&code, &code_info.Type, &code_info.EffectTimestamp, &code_info.ExpireTimestamp, &code_info.ItemId, &code_info.Version)
		if err != nil {
			log.Println(err)
			http.Error(resp, "db query result failed", http.StatusInternalServerError)
			return
		}
		code_list[code] = code_info
	}

	// get the response string
	json_str, err := json.Marshal(code_list)
	if err != nil {
		log.Println(err)
		http.Error(resp, "json generating failed", http.StatusInternalServerError)
		return
	}

	// write out
	resp.Write(json_str)
}

type GiftCodeInfo struct {
	ItemId          int16  `json:"item_id"`
	Type            int8   `json:"type"`
	EffectTimestamp int64  `json:"effect_timestamp"`
	ExpireTimestamp int64  `json:"expire_timestamp"`
	Version         int64  `json:"version"`
	Content         string `json:"content"`
}
