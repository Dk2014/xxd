package fashion

import (
	"core/time"
	"game_server/dat/fashion_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
)

type FashionMod struct{}

func init() {
	module.Fashion = FashionMod{}
}

//增加一件时装
func (mod FashionMod) AddFashion(db *mdb.Database, fashionId int16, validHours int64) *mdb.PlayerFashion {
	var playerFashion *mdb.PlayerFashion
	db.Select.PlayerFashion(func(row *mdb.PlayerFashionRow) {
		if row.FashionId() == fashionId {
			playerFashion = row.GoObject()
			row.Break()
		}
	})
	now := time.GetNowTime()
	if playerFashion == nil { //新获得？
		playerFashion = &mdb.PlayerFashion{
			Pid:       db.PlayerId(),
			FashionId: fashionId,
		}
		if validHours > 0 { //时装有有效期？
			playerFashion.ExpireTime = now + validHours*3600
		}
		db.Insert.PlayerFashion(playerFashion)
	} else {
		if playerFashion.ExpireTime > 0 { //已有的不是永久时装
			if validHours == 0 { //获得永久时装
				playerFashion.ExpireTime = 0
			} else {
				if playerFashion.ExpireTime < now { //已有时装已过期
					playerFashion.ExpireTime = now + validHours*3600
				} else {
					playerFashion.ExpireTime += validHours * 3600
				}
			}
			db.Update.PlayerFashion(playerFashion)
		}
	}
	tlog.PlayerSystemModuleFlowLog(db, tlog.SMT_FASHION)
	return playerFashion
}

//登录时删除过期时装
func (mod FashionMod) LoginUpdateFashion(db *mdb.Database) {
	now := time.GetNowTime()
	playerFashionState := db.Lookup.PlayerFashionState(db.PlayerId())
	var shouldUndress bool = true
	db.Select.PlayerFashion(func(row *mdb.PlayerFashionRow) {
		if row.ExpireTime() > 0 && row.ExpireTime() < now { //时装有时限且过期
			db.Delete.PlayerFashion(row.GoObject())
		} else {
			if row.FashionId() == playerFashionState.DressedFashionId {
				shouldUndress = false
			}
		}
	})
	if shouldUndress {
		playerFashionState.DressedFashionId = 0
		db.Update.PlayerFashionState(playerFashionState)
	}
}

func (mod FashionMod) FashionEnable(db *mdb.Database) bool {
	playerFashionState := db.Lookup.PlayerFashionState(db.PlayerId())
	return playerFashionState.DressedFashionId != fashion_dat.EMPTY_FASHION
}
