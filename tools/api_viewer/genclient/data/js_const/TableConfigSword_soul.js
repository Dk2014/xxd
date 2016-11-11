

var TableConfigSword_soul = 
{

	// =====================================================================
	//
	// $config[0]
	//
	// =====================================================================
	
	MAX_DRAW_NUM : 100,
	MAX_INGOT_DRAW_NUM : 20,
	RECOVERY_VALUE : 1,
	RECOVERY_TIME : 300,
	BAG_ON_BODY_NUM : 9,
	BAG_START : 0,
	BAG_END : 35,
	BAG_NUM : 36,
	BOX_A : 0,
	BOX_B : 1,
	BOX_C : 2,
	BOX_D : 3,
	BOX_E : 4,
	BOX_NUM : 5,
	QIAN_LONG_ID : 37,
	RUBBISH_AWARD_COIN : 500,
	

	// =====================================================================
	//
	// $config[2]
	//
	// =====================================================================
	
	BoxConfigs :  [
		[1,1000,"败亡铁剑"],
		[2,2000,"败亡阔剑"],
		[4,3000,"败亡魔剑"],
		[8,4000,"败亡宝剑"],
		[16,8000,"败亡神剑"]
	],
	

	// =====================================================================
	//
	// $config[3]
	//
	// =====================================================================
	

	getOpenedPosNum: function(level) 
	{
		if (level <= 9) 
		{
			return 1;
		}
		else if (level >= 10 && level <= 19) 
		{
			return 2;
		}
		else if (level >= 20 && level <= 29) 
		{
			return 3;
		}
		else if (level >= 30 && level <= 39) 
		{
			return 4;
		}
		else if (level >= 40 && level <= 49) 
		{
			return 5;
		}
		else if (level >= 50 && level <= 59) 
		{
			return 6;
		}
		else if (level >= 60 && level <= 69) 
		{
			return 7;
		}
		else if (level >= 70 && level <= 79) 
		{
			return 8;
		}
		else if (level >= 80) 
		{
			return 9;
		}
		return 1;
	},

	// =====================================================================
	//
	// $config[4]
	//
	// =====================================================================
	

	swordSoulMaxLevel: function(level) 
	{
		if (level <= 20) 
		{
			return 4;
		}
		else if (level >= 21 && level <= 40) 
		{
			return 8;
		}
		else if (level >= 41 && level <= 60) 
		{
			return 12;
		}
		else if (level >= 61 && level <= 80) 
		{
			return 16;
		}
		else if (level >= 81) 
		{
			return 20;
		}
		return 1;
	},

	// =====================================================================
	//
	// $config[5]
	//
	// =====================================================================
	

	getSwordDrawPriceInIngot: function(times) 
	{
		if (times >= 1 && times <= 2) 
		{
			return 50;
		}
		else if (times >= 3 && times <= 5) 
		{
			return 75;
		}
		else if (times >= 6 && times <= 10) 
		{
			return 100;
		}
		else if (times >= 11 && times <= 20) 
		{
			return 200;
		}
		return 0;
	},

	// =====================================================================
	//
	// $config[6]
	//
	// =====================================================================
	

	getSwordDrawPriceInIngotCoin: function(times) 
	{
		if (times >= 1 && times <= 2) 
		{
			return 4000;
		}
		else if (times >= 3 && times <= 5) 
		{
			return 4000;
		}
		else if (times >= 6 && times <= 10) 
		{
			return 4000;
		}
		else if (times >= 11 && times <= 20) 
		{
			return 4000;
		}
		return 0;
	},

};