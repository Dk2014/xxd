var totem_call_cost_config_data = {
		/**
	 * 0 : times, smallint, 召唤次数
	 * 1 : cost, int, 单价 
	 */

	Times : 0,
	Cost : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 100 /*[1]*/],
		[2 /*[0]*/, 100 /*[1]*/],
		[3 /*[0]*/, 150 /*[1]*/],
		[4 /*[0]*/, 150 /*[1]*/],
		[5 /*[0]*/, 150 /*[1]*/],
		[6 /*[0]*/, 200 /*[1]*/],
		[7 /*[0]*/, 200 /*[1]*/],
		[8 /*[0]*/, 200 /*[1]*/],
		[9 /*[0]*/, 200 /*[1]*/],
		[10 /*[0]*/, 200 /*[1]*/],
		[11 /*[0]*/, 300 /*[1]*/]
	],
};
