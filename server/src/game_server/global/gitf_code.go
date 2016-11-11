package global

import (
	"core/fail"
	"core/mysql"
	"core/time"
	"encoding/json"
	"fmt"
	"game_server/config"
	"game_server/mdb"
	"sync"
)

const (
	GIFT_CODE_MONOP_TYPE int8 = 0 //独占类型兑换码 本服务器哪A玩家使用了之后其他玩家不能使用
	GIFT_CODE_SHARE_TYPE int8 = 1
)

var (
	g_GiftCodeMutex sync.Mutex
	g_GiftCode      map[string]GiftCodeInfo
)

type GiftCodeConfig struct {
	ItemType int8  `json:"item_type"`
	ItemId   int16 `json:"item_id"`
	Num      int64 `json:"num"`
}

type GiftCodeInfo struct {
	ItemId          int16             `json:"item_id"`
	Type            int8              `json:"type"`
	EffectTimestamp int64             `json:"effect_timestamp"`
	ExpireTimestamp int64             `json:"expire_timestamp"`
	Version         int64             `json:"version"`
	Content         string            `json:"content"`
	Config          []*GiftCodeConfig `json:"config"`
}

//现在 LoadGiftCodeFromDB 需要依赖 内存数据库 mdb.Database 中的数据（需要在mdb加载数据之后运行）
//另外一种做法是直接读mysql
func LoadGiftCodeFromDB(gdb *mdb.Database, serverId int, mysqlInfo map[string]interface{}) {
	db, err := mysql.Connect(mysqlInfo)
	fail.When(err != nil, err)
	defer db.Close()
	now := time.GetNowTime()
	sql := fmt.Sprintf("select `code`, `item_id`, `type`, `effect_timestamp`, `expire_timestamp`, `version`, `content`, `config`from `gift_code` where `disable` = 0 and expire_timestamp > %d and (server_id = 0 or server_id = %d)", now, serverId)
	res, err := db.ExecuteFetch([]byte(sql), -1)
	fail.When(err != nil, fmt.Sprintf("err [%v] ; sql [%s]", sql))

	iCode := res.Map("code")
	iItemId := res.Map("item_id")
	iType := res.Map("type")
	iEffectTimestamp := res.Map("effect_timestamp")
	iExpireTimestamp := res.Map("expire_timestamp")
	iVersion := res.Map("version")
	iContent := res.Map("content")
	iConfig := res.Map("config")

	newGiftCode := map[string]GiftCodeInfo{}
	for _, row := range res.Rows {
		info := GiftCodeInfo{
			ItemId:          row.Int16(iItemId),
			Type:            row.Int8(iType),
			EffectTimestamp: row.Int64(iEffectTimestamp),
			ExpireTimestamp: row.Int64(iExpireTimestamp),
			Version:         row.Int64(iVersion),
			Content:         row.Str(iContent),
		}
		json.Unmarshal(row.Bin(iConfig), &info.Config)
		newGiftCode[row.Str(iCode)] = info
	}
	gdb.Select.GlobalGiftCardRecord(func(row *mdb.GlobalGiftCardRecordRow) {
		delete(newGiftCode, row.Code())
	})
	g_GiftCodeMutex.Lock()
	g_GiftCode = newGiftCode
	g_GiftCodeMutex.Unlock()
}

func LoadGiftCode(db *mdb.Database, configPath string) {
	newGiftCode := map[string]GiftCodeInfo{}
	_ = config.LoadJSONConfig(configPath, &newGiftCode)

	//TODO 过滤过期的
	//now := time.GetNowTime()

	//去掉 global_gift_exchange_log 里面的条目
	db.Select.GlobalGiftCardRecord(func(row *mdb.GlobalGiftCardRecordRow) {
		delete(newGiftCode, row.Code())
	})

	g_GiftCodeMutex.Lock()
	g_GiftCode = newGiftCode
	g_GiftCodeMutex.Unlock()
}

func TakeGift(code string) (GiftCodeInfo, bool) {
	g_GiftCodeMutex.Lock()
	defer g_GiftCodeMutex.Unlock()
	info, ok := g_GiftCode[code]
	if ok {
		if info.Type == GIFT_CODE_MONOP_TYPE {
			delete(g_GiftCode, code)
		}
		return info, true
	}
	return GiftCodeInfo{}, false
}
