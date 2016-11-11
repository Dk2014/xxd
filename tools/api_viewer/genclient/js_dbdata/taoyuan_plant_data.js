var taoyuan_plant_data = {
		/**
	 * 0 : id, smallint, 
	 * 1 : item_id, smallint, 对应物品Id
	 * 2 : need_level, smallint, 需求桃源之心等级
	 * 3 : bring_nums, smallint, 产出数量
	 * 4 : cost_item_id, smallint, 消耗物品ID
	 * 5 : cost_item_nums, smallint, 需求种子数量
	 * 6 : cost_physical, smallint, 需求体力
	 * 7 : bless_weight, smallint, 愿望权重
	 * 8 : cost_time, smallint, 种植所需时间
	 * 9 : cost_time_black, smallint, 黑土地种植所需时间 
	 */

	Id : 0,
	Item_id : 1,
	Need_level : 2,
	Bring_nums : 3,
	Cost_item_id : 4,
	Cost_item_nums : 5,
	Cost_physical : 6,
	Bless_weight : 7,
	Cost_time : 8,
	Cost_time_black : 9,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 9 /*[1]*/, 1 /*[2]*/, 1 /*[3]*/, 1 /*[4]*/, 1 /*[5]*/, 1 /*[6]*/, 2 /*[7]*/, 150 /*[8]*/, 120 /*[9]*/],
		[2 /*[0]*/, 10 /*[1]*/, 2 /*[2]*/, 1 /*[3]*/, 2 /*[4]*/, 1 /*[5]*/, 1 /*[6]*/, 5 /*[7]*/, 375 /*[8]*/, 300 /*[9]*/],
		[3 /*[0]*/, 11 /*[1]*/, 3 /*[2]*/, 1 /*[3]*/, 3 /*[4]*/, 1 /*[5]*/, 1 /*[6]*/, 10 /*[7]*/, 750 /*[8]*/, 600 /*[9]*/],
		[4 /*[0]*/, 12 /*[1]*/, 4 /*[2]*/, 1 /*[3]*/, 4 /*[4]*/, 1 /*[5]*/, 1 /*[6]*/, 15 /*[7]*/, 1125 /*[8]*/, 900 /*[9]*/],
		[5 /*[0]*/, 13 /*[1]*/, 5 /*[2]*/, 1 /*[3]*/, 5 /*[4]*/, 1 /*[5]*/, 1 /*[6]*/, 20 /*[7]*/, 1500 /*[8]*/, 1200 /*[9]*/]
	],
};
