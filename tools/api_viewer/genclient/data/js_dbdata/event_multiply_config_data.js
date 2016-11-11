var event_multiply_config_data = {
		/**
	 * 0 : condition, int, 加成的事件id
	 * 1 : times, float, 加成的倍数 
	 */

	Condition : 0,
	Times : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 2.00 /*[1]*/],
		[2 /*[0]*/, 2.00 /*[1]*/],
		[3 /*[0]*/, 1.05 /*[1]*/],
		[4 /*[0]*/, 1.05 /*[1]*/],
		[5 /*[0]*/, 1.10 /*[1]*/],
		[6 /*[0]*/, 1.10 /*[1]*/]
	],
};
