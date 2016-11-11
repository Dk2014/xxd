package ghost

import (
	// "core/fail"
	"fmt"
	"game_server/dat/ghost_dat"
	"game_server/dat/item_dat"
	"game_server/dat/mail_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/module/rpc"
	"game_server/tlog"
)

// 添加魂侍, id为Ghost表主键
func addGhost(db *mdb.Database, ghostId int16, itemFlowReason, xdEventType int32) (playerGhostId int64) {
	//for sealedbook  前置 ，需要考虑老玩家情况
	if session, exist := module.Player.GetPlayerOnline(db.PlayerId()); exist {
		state := module.State(session)
		if _, result := state.GetSealedBookRecord().FindRecord(item_dat.STEALDBOOK_TYPE_GHOSTS, ghostId); !result {
			sealedbook := item_dat.GetSealedBookInfo(item_dat.STEALDBOOK_TYPE_GHOSTS, ghostId)
			if sealedbook != nil {
				state.GetSealedBookRecord().AddRecord(item_dat.STEALDBOOK_TYPE_GHOSTS, ghostId, item_dat.STEALDBOOK_HAVING, db)
			}
		}
	}

	ghost := ghost_dat.GetGhost(ghostId)

	// 如果有相同的魂侍
	if isHasGhost(db, ghostId) {

		// 兑换为碎片
		module.Item.AddItem(db, ghost.FragmentId, ghost_dat.GHOST_DECOMPOSE_NUM, itemFlowReason, xdEventType, "")

		return playerGhostId
	}

	playerGhost := &mdb.PlayerGhost{
		Pid:        db.PlayerId(),
		GhostId:    ghostId,
		Star:       ghost.InitStar,
		Level:      1,
		SkillLevel: 1,
		Exp:        0,
		IsNew:      1,
		Pos:        ghost_dat.UNEQUIPPED,
	}

	db.Insert.PlayerGhost(playerGhost)
	playerGhostId = playerGhost.Id
	tlog.PlayerItemFlowLog(db, int16(ghostId), tlog.IT_GHOST, 1, tlog.ADD, itemFlowReason)

	if session, online := module.Player.GetPlayerOnline(db.PlayerId()); online {
		module.Notify.SendHaveNewGhost(session, playerGhostId)
	}

	return playerGhostId
}

func sendMailWithGhosts(db *mdb.Database, ghostIds []int16) {

	ghostAttachments := []*mail_dat.Attachment{}

	for _, ghostId := range ghostIds {
		ghostAttachments = append(ghostAttachments, &mail_dat.Attachment{mail_dat.ATTACHMENT_GHOST, ghostId, 1})
	}

	rpc.RemoteMailSend(db.PlayerId(), mail_dat.MailGhostBagFull{
		Func:        fmt.Sprintf("%d", len(ghostIds)),
		Attachments: ghostAttachments,
	})
}

// 魂侍是否唯一
func isHasGhost(db *mdb.Database, ghostId int16) bool {
	has := false
	db.Select.PlayerGhost(func(row *mdb.PlayerGhostRow) {
		if row.GhostId() == ghostId {
			has = true
			row.Break()
		}
	})
	return has
}

func clearGhostRelation(db *mdb.Database, ghost *mdb.PlayerGhost) {
	if ghost.RelationId > 0 {

		relationGhost := db.Lookup.PlayerGhost(ghost.RelationId)
		if relationGhost != nil && relationGhost.RelationId > 0 {
			relationGhost.RelationId = 0
			db.Update.PlayerGhost(relationGhost)
		}

		ghost.RelationId = 0
		db.Update.PlayerGhost(ghost)
	}
}
