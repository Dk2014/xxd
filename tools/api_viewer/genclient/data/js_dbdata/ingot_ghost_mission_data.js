var ingot_ghost_mission_data = {
		/**
	 * 0 : id, bigint, 主键ID
	 * 1 : vip_level, tinyint, 可进入vip等级
	 * 2 : max_egg_num, smallint, 一天开启魂蛋数量
	 * 3 : yellow_ghost_rand, smallint, 金色魂侍概率(万分之)
	 * 4 : purple_ghost_rand, smallint, 紫色魂侍概率(万分之)
	 * 5 : orange_ghost_rand, smallint, 橙色魂侍概率(万分之) 
	 */

	Id : 0,
	Vip_level : 1,
	Max_egg_num : 2,
	Yellow_ghost_rand : 3,
	Purple_ghost_rand : 4,
	Orange_ghost_rand : 5,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 0 /*[1]*/, 5 /*[2]*/, 200 /*[3]*/, 2500 /*[4]*/, 7300 /*[5]*/]
	],
};
