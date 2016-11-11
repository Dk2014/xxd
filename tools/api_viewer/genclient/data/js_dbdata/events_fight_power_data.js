var events_fight_power_data = {
		/**
	 * 0 : id, smallint, 
	 * 1 : lock, smallint, 档位
	 * 2 : fight, int, 战力
	 * 3 : ingot, smallint, 奖励元宝
	 * 4 : item1_id, smallint, 物品1
	 * 5 : item1_num, smallint, 物品1数量
	 * 6 : item2_id, smallint, 物品2
	 * 7 : item2_num, smallint, 物品2数量
	 * 8 : item3_id, smallint, 物品3
	 * 9 : item3_num, smallint, 物品3数量 
	 */

	Id : 0,
	Lock : 1,
	Fight : 2,
	Ingot : 3,
	Item1_id : 4,
	Item1_num : 5,
	Item2_id : 6,
	Item2_num : 7,
	Item3_id : 8,
	Item3_num : 9,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[3 /*[0]*/, 1 /*[1]*/, 20000 /*[2]*/, 100 /*[3]*/, 0 /*[4]*/, 0 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/],
		[6 /*[0]*/, 2 /*[1]*/, 50000 /*[2]*/, 200 /*[3]*/, 0 /*[4]*/, 0 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/],
		[7 /*[0]*/, 3 /*[1]*/, 75000 /*[2]*/, 300 /*[3]*/, 0 /*[4]*/, 0 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/],
		[8 /*[0]*/, 4 /*[1]*/, 100000 /*[2]*/, 400 /*[3]*/, 0 /*[4]*/, 0 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/]
	],
};
