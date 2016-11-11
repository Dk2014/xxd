

var TableConfigTeam = 
{

	// =====================================================================
	//
	// $config[0]
	//
	// =====================================================================
	
	POS0 : 0,
	POS1 : 1,
	POS2 : 2,
	POS3 : 3,
	POS4 : 4,
	POS5 : 5,
	POS6 : 6,
	POS7 : 7,
	POS8 : 8,
	POS_NO_ROLE : -1,
	TEAMSHIP_HEALTH_IND : 0,
	TEAMSHIP_ATTACK_IND : 1,
	TEAMSHIP_DEFENCE_IND : 2,
	TEAMSHIP_MAX_LEVEL : 100,
	

	// =====================================================================
	//
	// $config[1]
	//
	// =====================================================================
	

	getMaxInFormRoleNum: function(level) 
	{
		if (level <= 29) 
		{
			return 3;
		}
		else if (level >= 30 && level <= 49) 
		{
			return 4;
		}
		else if (level >= 50) 
		{
			return 5;
		}
		return 3;
	},

};