package notify

import (
	"core/net"
	"game_server/api/protocol/notify_api"
	"game_server/dat/item_dat"
	"game_server/mdb"
	"game_server/module"
	"time"
)

func init() {
	module.Notify = NotifyMod{}
}

type NotifyMod struct {
}

func (mod NotifyMod) SendPlayerKeyChanged(session *net.Session, key int32, max_order int8) {
	session.Send(&notify_api.PlayerKeyChanged_Out{
		Key:      int32(module.Item.GetItemNum(module.State(session).Database, item_dat.ITEM_MISSION_KEY_ID)),
		MaxOrder: max_order,
	})
}

func (mod NotifyMod) SendMissionLevelLockChanged(session *net.Session, max_lock int32, award_lock int32) {
	session.Send(&notify_api.MissionLevelLockChanged_Out{
		//Lock:    lock,
		MaxLock:   max_lock,
		AwardLock: award_lock,
	})
}

func (mod NotifyMod) SendRoleExpChanged(session *net.Session, roleId int8, addExp int64, exp int64, level int16) {
	session.Send(&notify_api.RoleExpChange_Out{
		RoleId: roleId,
		AddExp: addExp,
		Exp:    exp,
		Level:  level,
	})
}

func (mod NotifyMod) SendPhysicalChange(session *net.Session, value int16) {
	session.Send(&notify_api.PhysicalChange_Out{
		Value: value,
	})
}

func (mod NotifyMod) SendMoneyChange(session *net.Session, moneytype int8, value int64) {
	session.Send(&notify_api.MoneyChange_Out{
		Moneytype: moneytype,
		Value:     value,
		Timestamp: time.Now().UnixNano(),
	})
}

func (mod NotifyMod) SendSkillAdd(session *net.Session, roleId int8, skillId int16) {
	session.Send(&notify_api.SkillAdd_Out{
		RoleId:  roleId,
		SkillId: skillId,
	})
}

func (mod NotifyMod) SendItemChange(session *net.Session, items *notify_api.ItemChange_Out) {
	session.Send(items)
}

func (mod NotifyMod) RoleBattleStatusChange(session *net.Session) {
	state := module.State(session)
	rsp := &notify_api.RoleBattleStatusChange_Out{
		Roles: make([]notify_api.RoleBattleStatusChange_Out_Roles, len(state.MissionLevelState.AttackerInfo.Fighters)),
	}

	i := 0
	for roleId, fighter := range state.MissionLevelState.AttackerInfo.Fighters {
		rsp.Roles[i] = notify_api.RoleBattleStatusChange_Out_Roles{
			RoleId: int8(roleId),
			Health: int32(fighter.Health),
			Buffs:  []notify_api.RoleBattleStatusChange_Out_Roles_Buffs{},
		}
		i++
	}

	for _, role := range rsp.Roles {
		if _, ok := state.MissionLevelState.AttackerInfo.Buffs[int(role.RoleId)]; !ok {
			continue
		}

		for _, buff := range state.MissionLevelState.AttackerInfo.Buffs[int(role.RoleId)] {
			role.Buffs = append(role.Buffs, notify_api.RoleBattleStatusChange_Out_Roles_Buffs{
				Buffer: notify_api.BufferInfo{
					Mode:        notify_api.BuffMode(buff.Mode),
					Keep:        int8(buff.Keep),
					Value:       int32(buff.Value),
					SkillId:     int16(buff.Skill),
					MaxOverride: int8(buff.MaxOverride),
					OverrideNum: int8(buff.OverrideNum),
				},
			})
		}
	}

	session.Send(rsp)
}

func (mod NotifyMod) NewMail(session *net.Session) {
	out := &notify_api.NewMail_Out{}
	session.Send(out)
}

func (mod NotifyMod) SendHeartChange(session *net.Session, value int16) {
	session.Send(&notify_api.HeartChange_Out{
		Value: value,
	})
}

func (mod NotifyMod) SendQuestChange(session *net.Session, questId int16, state int8) {
	session.Send(&notify_api.QuestChange_Out{
		QuestId: questId,
		State:   state,
	})
}

func (mod NotifyMod) SendTownLockChange(session *net.Session, lock int32) {
	session.Send(&notify_api.TownLockChange_Out{
		Lock: lock,
	})
}

func (mod NotifyMod) SendFuncKeyChange(session *net.Session, funcKey int16) {
	session.Send(&notify_api.FuncKeyChange_Out{FuncKey: funcKey})
}

// 通知重建装备重铸
func (mod NotifyMod) SendItemRecastStateRebuild(session *net.Session) {
	state := module.State(session)
	db := state.Database

	// 发送物品信息
	module.Item.SendAllItems(session)

	// 装备重铸状态
	playerItemRecastState := db.Lookup.PlayerItemRecastState(state.PlayerId)
	out := &notify_api.ItemRecastStateRebuild_Out{
		Id:           playerItemRecastState.PlayerItemId,
		SelectedAttr: notify_api.Attribute(playerItemRecastState.SelectedAttr),
		Attrs: []notify_api.ItemRecastStateRebuild_Out_Attrs{
			{
				Attr:  notify_api.Attribute(playerItemRecastState.Attr1Type),
				Value: playerItemRecastState.Attr1Value,
			},
			{
				Attr:  notify_api.Attribute(playerItemRecastState.Attr2Type),
				Value: playerItemRecastState.Attr2Value,
			},
			{
				Attr:  notify_api.Attribute(playerItemRecastState.Attr3Type),
				Value: playerItemRecastState.Attr3Value,
			},
		},
	}

	session.Send(out)
}

//VIP等级变更
func (mod NotifyMod) VIPLevelChange(session *net.Session, level int16) {
	session.Send(&notify_api.VipLevelChange_Out{
		Level: level,
	})
}

// 通知增加新伙伴
func (mod NotifyMod) SendNewBuddy(session *net.Session, roleId int8, level int16) {
	session.Send(&notify_api.NotifyNewBuddy_Out{
		RoleId:    roleId,
		RoleLevel: level,
	})
}

// 通知难度关卡功能权值变化
func (mod NotifyMod) SendHardLevelLockChange(session *net.Session, lock int32) {
	session.Send(&notify_api.HardLevelLockChanged_Out{
		Lock: lock,
	})
}

// 通知剑心拔剑次数变更
func (mod NotifyMod) SendSwordSoulDrawNumChange(session *net.Session, num int16, cdTime int64) {
	session.Send(&notify_api.SendSwordSoulDrawNumChange_Out{
		Num:    num,
		CdTime: cdTime,
	})
}

//通知获得新魂侍
func (mod NotifyMod) SendHaveNewGhost(session *net.Session, ghostId int64) {
	session.Send(&notify_api.SendHaveNewGhost_Out{
		PlayerGhostId: ghostId,
	})
}

//通知下次爱心恢复时间
func (mod NotifyMod) SendHeartRecoverTime(session *net.Session, timestamp int64) {
	session.Send(&notify_api.SendHeartRecoverTime_Out{
		Timestamp: timestamp,
	})
}

//通知下次体力恢复时间
func (mod NotifyMod) SendPhysicalRecoverTime(session *net.Session, timestamp int64) {
	session.Send(&notify_api.SendPhysicalRecoverTime_Out{
		Timestamp: timestamp,
	})
}

//通知时装变更
func (mod NotifyMod) SendFashionChange(session *net.Session, playerFashion *mdb.PlayerFashion) {
	session.Send(&notify_api.SendFashionChange_Out{
		FashionId:  playerFashion.FashionId,
		ExpireTime: playerFashion.ExpireTime,
	})
}
