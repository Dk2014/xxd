var sword_soul_quality_data = {
		/**
	 * 0 : id, smallint, 剑心等级ID
	 * 1 : name, varchar, 品质名称
	 * 2 : sign, varchar, 程序标示
	 * 3 : color, varchar, 颜色 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Color : 3,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[0 /*[0]*/, "杂物" /*[1]*/, "NULL" /*[2]*/, "0xc5c3b7" /*[3]*/],
		[1 /*[0]*/, "特殊" /*[1]*/, "SPECIAL" /*[2]*/, "0xfff100" /*[3]*/],
		[2 /*[0]*/, "优良" /*[1]*/, "FINE" /*[2]*/, "0x22ac38" /*[3]*/],
		[3 /*[0]*/, "精良" /*[1]*/, "EXCELLENT" /*[2]*/, "0x00a0e9" /*[3]*/],
		[4 /*[0]*/, "传奇" /*[1]*/, "LEGEND" /*[2]*/, "0xc301c3" /*[3]*/],
		[5 /*[0]*/, "神器" /*[1]*/, "ARTIFACT" /*[2]*/, "0xfff100" /*[3]*/],
		[6 /*[0]*/, "唯一" /*[1]*/, "ONLY" /*[2]*/, "0xf39700" /*[3]*/]
	],
};
