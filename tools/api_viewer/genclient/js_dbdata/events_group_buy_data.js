var events_group_buy_data = {
		/**
	 * 0 : item_id, smallint, 参与团购得物品id
	 * 1 : base_price, smallint, 团购物品低价
	 * 2 : buy_times1, smallint, 购买次数1
	 * 3 : buy_percent1, float, 购买折扣1
	 * 4 : buy_times2, smallint, 购买次数2
	 * 5 : buy_percent2, float, 购买折扣2
	 * 6 : buy_times3, smallint, 购买次数3
	 * 7 : buy_percent3, float, 购买折扣3
	 * 8 : buy_times4, smallint, 购买次数4
	 * 9 : buy_percent4, float, 购买折扣4
	 * 10 : buy_times5, smallint, 购买次数5
	 * 11 : buy_percent5, float, 购买折扣5
	 * 12 : buy_times6, smallint, 购买次数6
	 * 13 : buy_percent6, float, 购买折扣6 
	 */

	Item_id : 0,
	Base_price : 1,
	Buy_times1 : 2,
	Buy_percent1 : 3,
	Buy_times2 : 4,
	Buy_percent2 : 5,
	Buy_times3 : 6,
	Buy_percent3 : 7,
	Buy_times4 : 8,
	Buy_percent4 : 9,
	Buy_times5 : 10,
	Buy_percent5 : 11,
	Buy_times6 : 12,
	Buy_percent6 : 13,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[342 /*[0]*/, 2000 /*[1]*/, 0 /*[2]*/, 0.75 /*[3]*/, 1 /*[4]*/, 0.70 /*[5]*/, 2 /*[6]*/, 0.65 /*[7]*/, 3 /*[8]*/, 0.60 /*[9]*/, 4 /*[10]*/, 0.55 /*[11]*/, 5 /*[12]*/, 0.50 /*[13]*/]
	],
};
