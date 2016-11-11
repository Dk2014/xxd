package sword_soul

import (
	"core/fail"
	"fmt"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/dat/sword_soul_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

func addSwordSoul(state *module.SessionState, swordSoulId int16, itemReason int32) (playerSwordSoulId int64) {
	db := state.Database

	//for sealedbook  前置 ，需要考虑老玩家情况
	if _, result := state.GetSealedBookRecord().FindRecord(item_dat.STEALDBOOK_TYPE_SWORDSOULS, swordSoulId); !result {
		sealedbook := item_dat.GetSealedBookInfo(item_dat.STEALDBOOK_TYPE_SWORDSOULS, swordSoulId)
		if sealedbook != nil {
			state.GetSealedBookRecord().AddRecord(item_dat.STEALDBOOK_TYPE_SWORDSOULS, swordSoulId, item_dat.STEALDBOOK_HAVING, db)
		}
	}

	// 判断是否已经有此唯一剑心
	swordSoulDat := sword_soul_dat.GetSwordSoul(swordSoulId)
	if swordSoulDat.Quality == sword_soul_dat.QUALITY_ONLY {
		db.Select.PlayerSwordSoul(func(row *mdb.PlayerSwordSoulRow) {
			fail.When(row.SwordSoulId() == swordSoulId, "already have this sword soul")
		})
	}

	//检查剑心背包是否已满
	posNum := emptyPosNum(state)
	if posNum <= 0 {
		sendMailWithSwordSouls(db, []int16{swordSoulId})
		return
	}

	playerSwordSoul := &mdb.PlayerSwordSoul{
		Pid:         state.PlayerId,
		SwordSoulId: swordSoulDat.Id,
		Exp:         0,
		Level:       1,
	}

	db.Insert.PlayerSwordSoul(playerSwordSoul)

	tlog.PlayerItemFlowLog(db, swordSoulId, tlog.IT_SWORDSOULFRAGMENT, 1, tlog.ADD, itemReason)
	playerSwordSoulId = playerSwordSoul.Id

	return
}

func sendMailWithSwordSouls(db *mdb.Database, swordSoulIds []int16) {

	swordSoulAttachments := []*mail_dat.Attachment{}

	for _, swordSoulId := range swordSoulIds {
		swordSoulAttachments = append(swordSoulAttachments, &mail_dat.Attachment{mail_dat.ATTACHMENT_SWORD_SOUL, swordSoulId, 1})
	}

	rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailSwordSoulBagFull{
		Func:        fmt.Sprintf("%d", len(swordSoulIds)),
		Attachments: swordSoulAttachments,
	})
}

// 判断背包是否以满，如果没有满返回第一个空位
func firstEmptyPos(state *module.SessionState) (pos int16) {
	db := state.Database

	var posBits int64

	db.Select.PlayerSwordSoul(func(row *mdb.PlayerSwordSoulRow) {
		pos := row.Pos()
		if pos != EQUIPPED {
			posBits |= 1 << uint(pos)
		}
	})

	for i := 0; i < sword_soul_dat.BAG_NUM; i++ {
		// 移位，弹出，判断是否为空位
		if posBits>>uint(i)&1 == 0 {
			return int16(i)
		}
	}

	return BAG_FULL
}

func emptyPosNum(state *module.SessionState) (posNum int16) {
	db := state.Database

	db.Select.PlayerSwordSoul(func(row *mdb.PlayerSwordSoulRow) {
		pos := row.Pos()
		if pos != EQUIPPED {
			posNum += 1
		}
	})

	return sword_soul_dat.BAG_NUM - posNum
}

func delSwordSoul(state *module.SessionState, playerSwordSoulId int64) {
	db := state.Database
	playerSwordSoul := db.Lookup.PlayerSwordSoul(playerSwordSoulId)
	fail.When(playerSwordSoul == nil, "wrong player sword soul id")
	db.Delete.PlayerSwordSoul(playerSwordSoul)
}
