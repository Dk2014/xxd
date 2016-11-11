var fate_box_data = {
		/**
	 * 0 : id, int, id
	 * 1 : name, varchar, 命锁名称
	 * 2 : sign, varchar, 标识符
	 * 3 : level, smallint, 要求等级
	 * 4 : require_lock, int, 命锁宝箱权值
	 * 5 : award_lock, int, 奖励命锁宝箱权值
	 * 6 : fixed_prop, smallint, 固定道具 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Level : 3,
	Require_lock : 4,
	Award_lock : 5,
	Fixed_prop : 6,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "星辉宝箱" /*[1]*/, "StarBox" /*[2]*/, 10 /*[3]*/, 0 /*[4]*/, 1000 /*[5]*/, 823 /*[6]*/],
		[2 /*[0]*/, "月影宝箱" /*[1]*/, "MoonBox" /*[2]*/, 40 /*[3]*/, 1000 /*[4]*/, 2000 /*[5]*/, 824 /*[6]*/],
		[3 /*[0]*/, "日耀宝箱" /*[1]*/, "SunBox" /*[2]*/, 80 /*[3]*/, 2000 /*[4]*/, 3000 /*[5]*/, 825 /*[6]*/],
		[4 /*[0]*/, "混元宝箱" /*[1]*/, "HunyuanBox" /*[2]*/, 120 /*[3]*/, 3000 /*[4]*/, 4000 /*[5]*/, 826 /*[6]*/]
	],
};
