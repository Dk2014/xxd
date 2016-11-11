var event_first_recharge_daily_data = {
		/**
	 * 0 : require_day, smallint, 索引天数
	 * 1 : ingot, smallint, 奖励元宝
	 * 2 : coins, int, 奖励铜钱
	 * 3 : heart, smallint, 奖励爱心
	 * 4 : item1_id, smallint, 物品1
	 * 5 : item1_num, smallint, 物品1数量
	 * 6 : item2_id, smallint, 物品2
	 * 7 : item2_num, smallint, 物品2数量
	 * 8 : item3_id, smallint, 物品3
	 * 9 : item3_num, smallint, 物品3数量
	 * 10 : item4_id, smallint, 物品4
	 * 11 : item4_num, smallint, 物品4数量
	 * 12 : item5_id, smallint, 物品5
	 * 13 : item5_num, smallint, 物品5数量 
	 */

	Require_day : 0,
	Ingot : 1,
	Coins : 2,
	Heart : 3,
	Item1_id : 4,
	Item1_num : 5,
	Item2_id : 6,
	Item2_num : 7,
	Item3_id : 8,
	Item3_num : 9,
	Item4_id : 10,
	Item4_num : 11,
	Item5_id : 12,
	Item5_num : 13,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 0 /*[3]*/, 335 /*[4]*/, 3 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/],
		[2 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 0 /*[3]*/, 425 /*[4]*/, 3 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/],
		[3 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 0 /*[3]*/, 274 /*[4]*/, 3 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/]
	],
};
