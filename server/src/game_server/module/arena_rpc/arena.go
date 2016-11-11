package arena_rpc

import (
	"game_server/api/protocol/arena_api"
	"game_server/dat/arena_award_box_dat"
	"game_server/global"
	"game_server/mdb"
	"math"
	"math/rand"
)

func getRandomRank(rank int32) int32 {
	var ceilRate, floorRate float64
	//根据策划规则，当参数 rank 为 2，3时，可能计算得到 非法排名0，需要特殊处理一下
	if rank <= 3 {
		return rand.Int31n(rank-1) + 1
	}
	if rank <= 30 {
		floorRate, ceilRate = 0.4, 0.5
	} else if rank <= 250 {
		floorRate, ceilRate = 0.2, 0.3
	} else {
		floorRate, ceilRate = 0.1, 0.15
	}
	lowRank := int32(math.Floor(float64(rank) * floorRate))
	hightRank := int32(math.Ceil(float64(rank) * ceilRate))
	return rand.Int31n(1+hightRank-lowRank) + lowRank
}

func getPlayerRankWithRank(db *mdb.Database, rank int32) (ranks []arena_api.Enter_Out_Ranks) {
	addRank := func(rank int32) {
		if row := db.Lookup.GlobalArenaRank(rank); row != nil {
			playerInfo := global.GetPlayerInfo(row.Pid)
			ranks = append(ranks, arena_api.Enter_Out_Ranks{
				Pid:      playerInfo.PlayerId,
				Nick:     playerInfo.PlayerNick,
				RoleId:   playerInfo.RoleId,
				Rank:     row.Rank,
				Level:    playerInfo.RoleLevel,
				FightNum: playerInfo.FightNum,
			})
		}
	}

	if rank <= 8 {
		for i := int32(1); i <= 8; i++ {
			addRank(i)
		}
		return ranks
	}

	//获取自己的信息
	addRank(rank)
	//获取比自己高一名的玩家
	addRank(rank - 1)

	//根据排名间隔获取5名玩家数据
	step := arena_award_box_dat.GetArenaGapByRank(rank)
	for i := int32(1); i <= 5; i++ {
		addRank(rank - 1 - i*step)
	}
	//根据排名间隔计算出来的最高排名获取玩家可挑战的最高排名
	addRank(getRandomRank(rank - 2 - 5*step))

	return
}
