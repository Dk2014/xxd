var buy_resource_times_config_data = {
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
		
	],
};
