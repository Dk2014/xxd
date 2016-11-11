package module

import (
	"game_server/mdb"
)

const (
	PLAYER_OPERATE_DRAW_WHEEL      = (1 << 0) // 爱心抽奖
	PLAYER_OPERATE_CHEST_COIN      = (1 << 1) // 青铜宝箱
	PLAYER_OPERATE_CHEST_INGOT     = (1 << 2) // 神龙宝箱
	PLAYER_OPERATE_BECOME_VIP      = (1 << 3) // 充值变为VIP
	PLAYER_OPERATE_SWORD_DRAW_1    = (1 << 4) // 第一次拔剑
	PLAYER_OPERATE_SWORD_DRAW_2    = (1 << 5) // 第二次拔剑
	PLAYER_OPERATE_CHEST_PET       = (1 << 6) // 灵宠宝箱首次
	PLAYER_OPERATE_CHEST_PET_TEN   = (1 << 7) // 灵宠宝箱首次十连抽
	PLAYER_OPERATE_CHEST_COIN_TEN  = (1 << 8) // 青龙宝箱首次十连抽
	PLAYER_OPERATE_CHEST_INGOT_TEN = (1 << 9) // 神龙宝箱首次十连抽
)

func SetPlayerFirstOperated(db *mdb.Database, operateValue int64) {
	playerIsOperated := db.Lookup.PlayerIsOperated(db.PlayerId())

	if (playerIsOperated.OperateValue & operateValue) != operateValue {
		playerIsOperated.OperateValue += operateValue
		db.Update.PlayerIsOperated(playerIsOperated)
	}
}

func CheckPlayerFirstOperate(db *mdb.Database, operateValue int64) bool {
	playerIsOperated := db.Lookup.PlayerIsOperated(db.PlayerId())
	return (playerIsOperated.OperateValue & operateValue) != operateValue
}
