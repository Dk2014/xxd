var heart_draw_data = {
		/**
	 * 0 : id, smallint, 主键ID
	 * 1 : draw_type, tinyint, 抽奖类型（1-大转盘；2-刮刮卡）
	 * 2 : daily_num, tinyint, 每日可抽奖次数
	 * 3 : cost_heart, tinyint, 每次抽奖消耗爱心数 
	 */

	Id : 0,
	Draw_type : 1,
	Daily_num : 2,
	Cost_heart : 3,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 1 /*[1]*/, 10 /*[2]*/, 2 /*[3]*/]
	],
};
