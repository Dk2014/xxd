var events_richman_club_awards_data = {
		/**
	 * 0 : require_vip_level, smallint, 所需的vip等级
	 * 1 : require_vip_count, smallint, 所需的vip相应人数
	 * 2 : award_vip_level1, smallint, 能领奖的vip等级1
	 * 3 : award_vip_item1_id, smallint, 能领奖的vip的奖励物品1 ID
	 * 4 : award_vip_item1_num, smallint, 能领奖的vip的奖励物品1数量 默认为1
	 * 5 : award_vip_level2, smallint, 能领奖的vip等级2
	 * 6 : award_vip_item2_id, smallint, 能领奖的vip的奖励物品2 ID
	 * 7 : award_vip_item2_num, smallint, 能领奖的vip的奖励物品2数量 默认为1
	 * 8 : award_vip_level3, smallint, 能领奖的vip等级3
	 * 9 : award_vip_item3_id, smallint, 能领奖的vip的奖励物品3 ID
	 * 10 : award_vip_item3_num, smallint, 能领奖的vip的奖励物品3数量 默认为1
	 * 11 : award_vip_level4, smallint, 能领奖的vip等级4
	 * 12 : award_vip_item4_id, smallint, 能领奖的vip的奖励物品4 ID
	 * 13 : award_vip_item4_num, smallint, 能领奖的vip的奖励物品4数量 默认为1
	 * 14 : award_vip_level5, smallint, 能领奖的vip等级5
	 * 15 : award_vip_item5_id, smallint, 能领奖的vip的奖励物品5 ID
	 * 16 : award_vip_item5_num, smallint, 能领奖的vip的奖励物品5数量 默认为1 
	 */

	Require_vip_level : 0,
	Require_vip_count : 1,
	Award_vip_level1 : 2,
	Award_vip_item1_id : 3,
	Award_vip_item1_num : 4,
	Award_vip_level2 : 5,
	Award_vip_item2_id : 6,
	Award_vip_item2_num : 7,
	Award_vip_level3 : 8,
	Award_vip_item3_id : 9,
	Award_vip_item3_num : 10,
	Award_vip_level4 : 11,
	Award_vip_item4_id : 12,
	Award_vip_item4_num : 13,
	Award_vip_level5 : 14,
	Award_vip_item5_id : 15,
	Award_vip_item5_num : 16,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[6 /*[0]*/, 10 /*[1]*/, 1 /*[2]*/, 404 /*[3]*/, 1 /*[4]*/, 0 /*[5]*/, 0 /*[6]*/, 0 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/, 0 /*[14]*/, 0 /*[15]*/, 0 /*[16]*/],
		[7 /*[0]*/, 5 /*[1]*/, 1 /*[2]*/, 405 /*[3]*/, 1 /*[4]*/, 2 /*[5]*/, 406 /*[6]*/, 1 /*[7]*/, 0 /*[8]*/, 0 /*[9]*/, 0 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/, 0 /*[14]*/, 0 /*[15]*/, 0 /*[16]*/],
		[8 /*[0]*/, 3 /*[1]*/, 1 /*[2]*/, 407 /*[3]*/, 1 /*[4]*/, 2 /*[5]*/, 408 /*[6]*/, 1 /*[7]*/, 3 /*[8]*/, 409 /*[9]*/, 1 /*[10]*/, 0 /*[11]*/, 0 /*[12]*/, 0 /*[13]*/, 0 /*[14]*/, 0 /*[15]*/, 0 /*[16]*/]
	],
};
