var hoilday_trader_data = {
		/**
	 * 0 : id, int, 标识
	 * 1 : item_id, smallint, 活动兑换礼包ID
	 * 2 : coin, bigint, 礼包铜钱价格
	 * 3 : ingot, bigint, 礼包元宝价格 
	 */

	Id : 0,
	Item_id : 1,
	Coin : 2,
	Ingot : 3,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 853 /*[1]*/, 0 /*[2]*/, 188 /*[3]*/],
		[2 /*[0]*/, 854 /*[1]*/, 0 /*[2]*/, 1688 /*[3]*/]
	],
};
