package despair_dat


const( 
	START_EVERY_WEEK_DAY = 1
	START_EVERY_WEEK_HOUR = 5
	MAX_BATTLE_NUM = 5
	MAX_BOSS_LEVEL = 99
	WAR_POINT = 2
	DESPAIR_TYPE_IN_MISSION_LEVEL = 19
	DEC_LEVEL = 1
	XS_BOSS = 1149
	SG_BOSS = 1150
	YM_BOSS = 1148
	XS_BOSS_LEVEL = 841
	SG_BOSS_LEVEL = 842
	YM_BOSS_LEVEL = 843
	XS = 1
	SG = 2
	YM = 3
)






var (
	BOSS_UP_CONFIG = map[int8]int16 {  // { time_hours : add_level } 
 	    1 : 5, 
	    12 : 3, 
	    24 : 2, 
	    48 : 1, 
	} 

 	CAMP_NAME_CONFIG = map[int8]string {  // { camp_type : name } 
 	    1 : "血兽军团",
	    2 : "尸鬼军团",
	    3 : "影魔军团",
	} 

 	BOSS_NAME_CONFIG = map[int8]string {  // { camp_type : name } 
 	    1 : "炼狱魁拔",
	    2 : "不朽尸王",
	    3 : "暗影魔煞",
	} 

 )

