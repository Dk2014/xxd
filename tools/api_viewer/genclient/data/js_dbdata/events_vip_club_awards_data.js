var events_vip_club_awards_data = {
		/**
	 * 0 : require_vip_count, smallint, 需要VIP人数
	 * 1 : ingot, smallint, 奖励元宝
	 * 2 : coins, int, 奖励铜钱
	 * 3 : item1_id, smallint, 物品1
	 * 4 : item1_num, smallint, 物品1数量
	 * 5 : item2_id, smallint, 物品2
	 * 6 : item2_num, smallint, 物品2数量
	 * 7 : item3_id, smallint, 物品3
	 * 8 : item3_num, smallint, 物品3数量
	 * 9 : item4_id, smallint, 物品4
	 * 10 : item4_num, smallint, 物品4数量
	 * 11 : item5_id, smallint, 物品5
	 * 12 : item5_num, smallint, 物品5数量 
	 */

	Require_vip_count : 0,
	Ingot : 1,
	Coins : 2,
	Item1_id : 3,
	Item1_num : 4,
	Item2_id : 5,
	Item2_num : 6,
	Item3_id : 7,
	Item3_num : 8,
	Item4_id : 9,
	Item4_num : 10,
	Item5_id : 11,
	Item5_num : 12,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[20 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 263 /*[3]*/, 5 /*[4]*/, 355 /*[5]*/, 5 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/],
		[50 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 263 /*[3]*/, 10 /*[4]*/, 355 /*[5]*/, 10 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/],
		[100 /*[0]*/, 0 /*[1]*/, 0 /*[2]*/, 263 /*[3]*/, 20 /*[4]*/, 355 /*[5]*/, 20 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/]
	],
};
