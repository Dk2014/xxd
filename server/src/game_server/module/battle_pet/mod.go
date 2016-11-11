package battle_pet

import (
	//"core/fail"
	"game_server/api/protocol/notify_api"
	"game_server/dat/battle_pet_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func init() {
	module.BattlePet = ModBattlePet{}
}

type ModBattlePet struct {
}

func (mod ModBattlePet) OpenFunc(db *mdb.Database) {

	for grid := int8(1); grid <= 5; grid++ {
		db.Insert.PlayerBattlePetGrid(&mdb.PlayerBattlePetGrid{
			Pid:    db.PlayerId(),
			GridId: grid,
			//Level:       1,
			//Exp:         0,
			BattlePetId: 0,
		})
	}
	db.Insert.PlayerBattlePetState(&mdb.PlayerBattlePetState{
		Pid: db.PlayerId(),
	})

	// 功能开启奖励草精
	module.BattlePet.AddPetWihoutNotify(db, battle_pet_dat.BATTLE_PET_CAOJING, tlog.IFR_OPEN_FUNC, xdlog.ET_OPEN_FUNC)
}

//根据格子ID获取当前宠物信息
func (mod ModBattlePet) GetBattlePetByGrid(db *mdb.Database, grid int8) (playerBattlePet *mdb.PlayerBattlePet) {
	var petId int32
	db.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.GridId() == grid {
			petId = row.BattlePetId()
			row.Break()
		}
	})
	db.Select.PlayerBattlePet(func(row *mdb.PlayerBattlePetRow) {
		if row.BattlePetId() == petId {
			playerBattlePet = row.GoObject()
			row.Break()
		}
	})

	return playerBattlePet
}

//获取上阵灵宠信息
func (mod ModBattlePet) GetAvailableBattlePet(db *mdb.Database) map[int8]int32 {
	petInfo := make(map[int8]int32)

	var petId int32
	db.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		petId = row.BattlePetId()
		if petId > 0 {
			petInfo[row.GridId()] = petId
		}
	})
	return petInfo
}

//pet id 是怪物 id
func (mod ModBattlePet) AddPet(db *mdb.Database, petId, itemFlowReason, xdEventType int32) {
	addPet(db, petId, itemFlowReason, xdEventType)
	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		session.Send(&notify_api.SendHaveNewPet_Out{
			PetId: petId,
		})
	}
}

func (mod ModBattlePet) AddPetWihoutNotify(db *mdb.Database, petId, itemFlowReason, xdEventType int32) {
	addPet(db, petId, itemFlowReason, xdEventType)
}

func (mod ModBattlePet) GetPlayerGridById(db *mdb.Database, gridId int8) (grid *mdb.PlayerBattlePetGrid) {
	db.Select.PlayerBattlePetGrid(func(row *mdb.PlayerBattlePetGridRow) {
		if row.GridId() == gridId {
			grid = row.GoObject()
			row.Break()
		}
	})
	return
}
