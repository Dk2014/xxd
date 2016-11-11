var ghost_baptize_data = {
		/**
	 * 0 : id, tinyint, 魂侍洗炼ID
	 * 1 : star, tinyint, 魂侍星级
	 * 2 : max_add_growth, tinyint, 最大增加成长值 
	 */

	Id : 0,
	Star : 1,
	Max_add_growth : 2,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 1 /*[1]*/, 5 /*[2]*/],
		[2 /*[0]*/, 2 /*[1]*/, 10 /*[2]*/],
		[3 /*[0]*/, 3 /*[1]*/, 15 /*[2]*/],
		[4 /*[0]*/, 4 /*[1]*/, 20 /*[2]*/],
		[5 /*[0]*/, 5 /*[1]*/, 25 /*[2]*/]
	],
};
