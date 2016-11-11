package meditation_dat


const( 
	EXP_UNIT_TIME = 15
	KEY_UNIT_TIME = 21600
	MAX_MEDITATION_TIME = 43200
	AUTO_MEDITATE_DELAY = 600
	MEDITATION_STOP = 0
	CLUBHOUSE_MEDITATION_INFO = 1
	CLUBHOUSE_MEDITATION_START = 2
	CLUBHOUSE_MEDITATION_STOP = 3
)







//自动生成请勿修改
//打坐经验
func GetMetationExpByLevel(level int16) int32 {
	if level >= 120 {
		return 8
	}
	if level >= 110 {
		return 7
	}
	if level >= 100 {
		return 6
	}
	if level >= 90 {
		return 5
	}
	if level >= 80 {
		return 4
	}
	if level >= 60 {
		return 3
	}
	if level >= 40 {
		return 2
	}
	if level >= 16 {
		return 1
	}
	return 0
}

