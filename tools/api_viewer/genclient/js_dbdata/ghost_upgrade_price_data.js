var ghost_upgrade_price_data = {
		/**
	 * 0 : quality, tinyint, 品质
	 * 1 : cost, int, 碎片单价 
	 */

	Quality : 0,
	Cost : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[0 /*[0]*/, 40 /*[1]*/],
		[1 /*[0]*/, 50 /*[1]*/],
		[2 /*[0]*/, 60 /*[1]*/]
	],
};
