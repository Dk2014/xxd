var fame_system_data = {
		/**
	 * 0 : id, smallint, 功能ID
	 * 1 : name, varchar, 系统名称
	 * 2 : sign, varchar, 唯一标示
	 * 3 : max_fame, int, 最大产生声望 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Max_fame : 3,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "比武场" /*[1]*/, "ARENA" /*[2]*/, 100000 /*[3]*/],
		[2 /*[0]*/, "多人关卡" /*[1]*/, "MULT_LEVEL" /*[2]*/, 15000 /*[3]*/],
		[4 /*[0]*/, "帮派" /*[1]*/, "CLIQUE" /*[2]*/, 1000000 /*[3]*/]
	],
};
