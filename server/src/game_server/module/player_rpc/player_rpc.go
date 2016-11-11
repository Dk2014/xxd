package player_rpc

import (
	"core/net"
	"core/time"
	"game_server/api/protocol/player_api"
	"game_server/dat/event_dat"
	"game_server/dat/player_dat"
	"game_server/module"
	"game_server/module/player_common"
	"math"
)

func getPlayerRankTable(rankType player_api.RankType) *player_common.PlayerRankTable {
	switch rankType {

	case player_api.RANK_TYPE_FIGHTNUM:
		return playerFightNumRankTable

	case player_api.RANK_TYPE_LEVEL:
		return playerLevelRankTable

	case player_api.RANK_TYPE_MAINROLE_FIGHTNUM:
		return playerMainRoleFightNumRankTable

	case player_api.RANK_TYPE_INGOT:
		return playerIngotRankTable
	}
	return nil
}

func GetRanks(session *net.Session, in *player_api.GetRanks_In) {
	state := module.State(session)
	out := &player_api.GetRanks_Out{}

	pid := state.PlayerId
	pageIndex := in.PageIndex
	rankType := in.Flag

	if rankTypeFlag != in.Flag {
		desplay_num = 0
		rankTypeFlag = player_api.RANK_TYPE_NULL
	}
	// 获取排行列表
	table := getPlayerRankTable(rankType)

	// 自身排行数据
	selfData := table.GetByPid(pid)
	// fail.When(selfData == nil, "Fail cause not find self")
	if selfData == nil {
		session.Send(out)
		return
	}

	// 当页起始名次
	rankIndex := pageIndex*(player_dat.RANK_EVERY_PAGE_NUM-1) + 1
	if rankIndex > selfData.Rank {
		rankIndex++
	}

	// 迭代player_dat.RANK_EVERY_PAGE_NUM + 1个数据
	table.IterRank(rankIndex, func(data *player_common.PlayerRankData) bool {
		outRank := player_api.GetRanks_Out_Ranks{}
		outRank.Pid = data.Pid
		outRank.Nickname = []byte(data.Nick)
		outRank.Rank = data.Rank
		outRank.Values = nil
		outRank.Values = append(outRank.Values, data.Values...)
		out.Ranks = append(out.Ranks, outRank)
		return len(out.Ranks) < player_dat.RANK_EVERY_PAGE_NUM+1 /*check if has next*/
	})

	// 确保自身数据在当前页
	if rankIndex > selfData.Rank || rankIndex+player_dat.RANK_EVERY_PAGE_NUM <= selfData.Rank {
		// selfData位置
		pos := int(math.Min(float64(player_dat.RANK_EVERY_PAGE_NUM-1), float64(len(out.Ranks))))
		if rankIndex > selfData.Rank {
			pos = 0
		}

		// selfData
		outRank := player_api.GetRanks_Out_Ranks{}
		outRank.Pid = selfData.Pid
		outRank.Nickname = []byte(selfData.Nick)
		outRank.Rank = selfData.Rank
		outRank.Values = nil
		outRank.Values = append(outRank.Values, selfData.Values...)

		// 插入
		out.Ranks = append(out.Ranks, player_api.GetRanks_Out_Ranks{})
		copy(out.Ranks[pos+1:], out.Ranks[pos:])
		out.Ranks[pos] = outRank
	}

	// 检查当前页是否显示完整
	out.HasNext = len(out.Ranks) > player_dat.RANK_EVERY_PAGE_NUM

	// 截取一页显示量
	out.Ranks = out.Ranks[:int(math.Min(float64(player_dat.RANK_EVERY_PAGE_NUM), float64(len(out.Ranks))))]

	session.Send(out)
}

//不用每页都显示玩家自己
func GetRanksForIngotRank(session *net.Session, in *player_api.GetRanks_In) {
	isOpenOrEndIngotRankTime(session)
	//	state := module.State(session)
	out := &player_api.GetRanks_Out{}

	if rankTypeFlag != in.Flag {
		desplay_num = 0
		rankTypeFlag = in.Flag
	}

	//	pid := state.PlayerId
	pageIndex := in.PageIndex
	rankType := in.Flag

	// 获取排行列表
	table := getPlayerRankTable(rankType)
	// 当页起始名次
	rankIndex := pageIndex*player_dat.RANK_EVERY_PAGE_NUM + 1
	// 迭代player_dat.RANK_EVERY_PAGE_NUM + 1个数据
	table.IterRank(rankIndex, func(data *player_common.PlayerRankData) bool {
		outRank := player_api.GetRanks_Out_Ranks{}
		outRank.Pid = data.Pid
		outRank.Nickname = []byte(data.Nick)
		outRank.Rank = data.Rank
		outRank.Values = nil
		outRank.Values = append(outRank.Values, data.Values...)
		out.Ranks = append(out.Ranks, outRank)

		if desplay_num > player_dat.RANK_INGOT_DESPLAY_NUM {
			rankTypeFlag = player_api.RANK_TYPE_NULL
			desplay_num = 0
			return false
		}
		desplay_num++
		return len(out.Ranks) < player_dat.RANK_EVERY_PAGE_NUM+1 /*check if has next*/
	})
	// 检查当前页是否显示完整
	out.HasNext = len(out.Ranks) > player_dat.RANK_EVERY_PAGE_NUM

	// 截取一页显示量
	out.Ranks = out.Ranks[:int(math.Min(float64(player_dat.RANK_EVERY_PAGE_NUM), float64(len(out.Ranks))))]

	session.Send(out)
}

//判断当前时间是否是开放或者关闭充值排行榜的时间
func isOpenOrEndIngotRankTime(session *net.Session) {
	eventIngotRankConfig := event_dat.GetEventIngotRankConfig()
	openIngotRankTime := eventIngotRankConfig.StartUnixTime
	endIngotRankTime := eventIngotRankConfig.EndUnixTime
	nowTime := time.GetNowTime()
	isOpen := false
	if nowTime >= openIngotRankTime && nowTime <= endIngotRankTime {
		isOpen = true
	}
	session.Send(&player_api.OpenIngotRank_Out{
		IsOpenIngotRank: isOpen,
	})
	closeIngotRankPanelTime := eventIngotRankConfig.EndUnixTime + 24*60*60
	isClose := false
	if nowTime >= closeIngotRankPanelTime {
		isClose = true
	}
	session.Send(&player_api.CloseIngotRankPanel_Out{
		IsCloseIngotRankPanel: isClose,
	})
}
