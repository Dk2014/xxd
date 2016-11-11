package town

import (
	"core/fail"

	"game_server/api/protocol/town_api"
	"game_server/dat/town_dat"
	"game_server/mdb"
	"game_server/module"
	"game_server/tlog"
	"game_server/xdlog"
)

func talkedNpcList(state *module.SessionState, townId int16, out *town_api.TalkedNpcList_Out) {
	fail.When(state.TownId != townId, "player not in target town")
	//获取玩家任务状态
	state.Database.Select.PlayerNpcTalkRecord(func(row *mdb.PlayerNpcTalkRecordRow) {
		if row.TownId() == townId {
			out.NpcList = append(out.NpcList, town_api.TalkedNpcList_Out_NpcList{
				NpcId:   row.NpcId(),
				QuestId: row.QuestId(),
			})
		}
	})
}

func npcTalkAward(state *module.SessionState, npcId int32) {
	//获取玩家任务状态 玩家进入游戏会自动完成任务1，并自动接任务2
	playerQuest := state.Database.Lookup.PlayerQuest(state.PlayerId)

	//获取NPC对话奖励配置
	npcTalk := town_dat.GetNPCTalkAward(npcId)
	fail.When(npcTalk.TownId != state.TownId, "玩家与NPC不在同一个城镇")

	//获取玩家对话记录
	var talkRecord *mdb.PlayerNpcTalkRecord
	state.Database.Select.PlayerNpcTalkRecord(func(row *mdb.PlayerNpcTalkRecordRow) {
		if row.NpcId() == npcId {
			talkRecord = row.GoObject()
			row.Break()
		}
	})

	if talkRecord == nil {
		//对话记录为空，优先考虑首次对话奖励
		talkRecord = &mdb.PlayerNpcTalkRecord{
			Pid:     state.PlayerId,
			TownId:  npcTalk.TownId,
			NpcId:   npcId,
			QuestId: town_dat.FIRST_TALK_QUEST_ID,
		}

		//优先查找首次对话奖励
		award, awardTalk := npcTalk.Awards[town_dat.FIRST_TALK_QUEST_ID]

		//如果NPC没有首次对话奖励则考虑任务奖励
		if !awardTalk {
			award, awardTalk = npcTalk.Awards[playerQuest.QuestId]
			fail.When(!awardTalk, "NPC没有配置对话奖励")
			talkRecord.QuestId = playerQuest.QuestId
		}

		state.Database.Insert.PlayerNpcTalkRecord(talkRecord)
		module.Item.AddItem(state.Database, award.ItemId, award.ItemNum, tlog.IFR_NPC_FIRST_TALK, xdlog.ET_NPC_TALK, "")
	} else {
		//存在对话记录，检查玩家任务状态
		fail.When(talkRecord.QuestId >= playerQuest.QuestId, "对话奖励已领取")

		//查找任务对话奖励
		award, awardTalk := npcTalk.Awards[playerQuest.QuestId]
		fail.When(!awardTalk, "NPC没有配置对话奖励")

		talkRecord.QuestId = playerQuest.QuestId
		state.Database.Update.PlayerNpcTalkRecord(talkRecord)
		module.Item.AddItem(state.Database, award.ItemId, award.ItemNum, tlog.IFR_NPC_FIRST_TALK, xdlog.ET_NPC_TALK, "")
	}
}
