package arena_award_box_dat


const( 
	MAX_DAILY_NUM = 8
	AWARD_ANYINGGUOSHI_WIN_HIGH_RANK = 5
	AWARD_ANYINGGUOSHI_WIN_LOW_RANK = 3
	AWARD_ANYINGGUOSHI_LOSE_LOW_RANK = 1
	LOSE_CD_TIME_SECONDS = 600
	ARENA_ATTACK_SUCC = 1
	ARENA_ATTACK_FAILED = 2
	ARENA_ATTACKED_SUCC = 3
	ARENA_ATTACKED_FAILED = 4
	CD_TIME_ONE_MIN_COST_INGOT = 1
	MAX_RESERVED_RECORD = 20
	GETBACK_AWARD_BOX_DAY1_COST_INGOT = 20
	GETBACK_AWARD_BOX_DAY2_COST_INGOT = 40
	GETBACK_AWARD_BOX_DAY3_COST_INGOT = 80
)






//自动生成请勿修改
func GetArenaGapByRank(rank int32) (gap int32) {
	if rank <= 30 {
		return 1
	}
	if rank <= 50 {
		return 2
	}
	if rank <= 100 {
		return 5
	}
	if rank <= 250 {
		return 10
	}
	if rank <= 500 {
		return 20
	}
	if rank <= 750 {
		return 50
	}
	if rank <= 1000 {
		return 60
	}
	if rank <= 1500 {
		return 80
	}
	if rank <= 3000 {
		return 150
	}
	if rank <= 5000 {
		return 200
	}
	if rank <= 10000 {
		return 300
	}
	if rank <= 20000 {
		return 800
	}
	if rank <= 30000 {
		return 1500
	}
	if rank <= 50000 {
		return 1800
	}
	if rank <= 100000 {
		return 2500
	}
	if rank <= 200000 {
		return 3000
	}
	return 1
}

