var rainbow_buy_cost_config_data = {
		/**
	 * 0 : times, int, 购买次数
	 * 1 : cost, int, 购买所需元宝 
	 */

	Times : 0,
	Cost : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 80 /*[1]*/],
		[2 /*[0]*/, 100 /*[1]*/],
		[3 /*[0]*/, 100 /*[1]*/],
		[4 /*[0]*/, 200 /*[1]*/],
		[5 /*[0]*/, 200 /*[1]*/],
		[6 /*[0]*/, 500 /*[1]*/],
		[7 /*[0]*/, 500 /*[1]*/],
		[8 /*[0]*/, 500 /*[1]*/]
	],
};