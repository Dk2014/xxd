var clique_building_data = {
		/**
	 * 0 : id, int, 标示ID
	 * 1 : name, varchar, 建筑名称
	 * 2 : biaoshi, varchar, 建筑标识
	 * 3 : desc, text, 建筑描述
	 * 4 : order, smallint, 优先权排序 
	 */

	Id : 0,
	Name : 1,
	Biaoshi : 2,
	Desc : 3,
	Order : 4,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[5 /*[0]*/, "总舵" /*[1]*/, "ZongDuo" /*[2]*/, "" /*[3]*/, 0 /*[4]*/],
		[1 /*[0]*/, "宗祠" /*[1]*/, "ZongCi" /*[2]*/, "" /*[3]*/, 1 /*[4]*/],
		[7 /*[0]*/, "战备仓库" /*[1]*/, "ZhanBeiCangKu" /*[2]*/, "" /*[3]*/, 2 /*[4]*/],
		[6 /*[0]*/, "钱庄" /*[1]*/, "QianZhuang" /*[2]*/, "" /*[3]*/, 3 /*[4]*/],
		[2 /*[0]*/, "回春堂" /*[1]*/, "HuiChunTang" /*[2]*/, "" /*[3]*/, 4 /*[4]*/],
		[3 /*[0]*/, "神兵堂" /*[1]*/, "ShenBingTang" /*[2]*/, "" /*[3]*/, 5 /*[4]*/],
		[4 /*[0]*/, "金刚堂" /*[1]*/, "JinGangTang" /*[2]*/, "" /*[3]*/, 6 /*[4]*/]
	],
};
