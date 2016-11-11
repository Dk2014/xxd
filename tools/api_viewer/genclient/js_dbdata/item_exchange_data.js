var item_exchange_data = {
		/**
	 * 0 : id, bigint, 主键ID
	 * 1 : target_item_id, smallint, 目标物品id
	 * 2 : item_id, smallint, 物品id
	 * 3 : item_num, smallint, 物品数量
	 * 4 : target_item_num, smallint, 目标物品数量 
	 */

	Id : 0,
	Target_item_id : 1,
	Item_id : 2,
	Item_num : 3,
	Target_item_num : 4,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		
	],
};
