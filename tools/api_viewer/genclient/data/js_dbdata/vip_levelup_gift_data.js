var vip_levelup_gift_data = {
		/**
	 * 0 : vip_level, smallint, 要求vip等级
	 * 1 : item_id, smallint, 物品ID
	 * 2 : item_num, smallint, 物品数量 
	 */

	Vip_level : 0,
	Item_id : 1,
	Item_num : 2,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[8 /*[0]*/, 350 /*[1]*/, 1 /*[2]*/]
	],
};
