package team_dat


const( 
	POS0 = 0
	POS1 = 1
	POS2 = 2
	POS3 = 3
	POS4 = 4
	POS5 = 5
	POS6 = 6
	POS7 = 7
	POS8 = 8
	POS_NO_ROLE = -1
	TEAMSHIP_HEALTH_IND = 0
	TEAMSHIP_ATTACK_IND = 1
	TEAMSHIP_DEFENCE_IND = 2
	TEAMSHIP_MAX_LEVEL = 100
)



func GetMaxInFormRoleNum(level int16) (res int8) {
	if level <= 29 {
		return 3
	}
	if level >= 30 && level <= 49 {
		return 4
	}
	if level >= 50 {
		return 5
	}
	panic("GetMaxInFormRoleNum err")
	return 3
}




