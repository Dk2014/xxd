var tower_level_data = {
		/**
	 * 0 : id, smallint, 
	 * 1 : floor, smallint, 楼层 
	 */

	Id : 0,
	Floor : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 1 /*[1]*/]
	],
};
