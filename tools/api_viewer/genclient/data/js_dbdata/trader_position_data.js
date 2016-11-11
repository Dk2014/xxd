var trader_position_data = {
		/**
	 * 0 : trader_id, smallint, 随机商人ID
	 * 1 : town_id, smallint, 城镇ID
	 * 2 : x, int, x轴坐标
	 * 3 : y, int, y轴坐标
	 * 4 : direction, varchar, 朝向 
	 */

	Trader_id : 0,
	Town_id : 1,
	X : 2,
	Y : 3,
	Direction : 4,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		
	],
};
