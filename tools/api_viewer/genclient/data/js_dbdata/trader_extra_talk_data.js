var trader_extra_talk_data = {
		/**
	 * 0 : trader_id, smallint, 随机商人ID
	 * 1 : time, tinyint, 点击次数
	 * 2 : talk, varchar, 对话 
	 */

	Trader_id : 0,
	Time : 1,
	Talk : 2,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[2 /*[0]*/, 20 /*[1]*/, "不要再点我了！小天会生气了的！" /*[2]*/],
		[2 /*[0]*/, 1 /*[1]*/, "小天和真真是好朋友。" /*[2]*/],
		[2 /*[0]*/, 5 /*[1]*/, "小天会保护真真探索世界的。" /*[2]*/],
		[2 /*[0]*/, 10 /*[1]*/, "为什么要一直点我！" /*[2]*/]
	],
};
