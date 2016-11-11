var totem_data = {
		/**
	 * 0 : id, smallint, 主键
	 * 1 : name, varchar, 阵印名称
	 * 2 : sign, varchar, 资源标识
	 * 3 : quality, tinyint, 阵印品质
	 * 4 : music_sign, varchar, 音乐资源标识 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Quality : 3,
	Music_sign : 4,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "青·石印" /*[1]*/, "QingShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[2 /*[0]*/, "兰·石印" /*[1]*/, "LanShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[3 /*[0]*/, "莹·石印" /*[1]*/, "YingShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[4 /*[0]*/, "赤·石印" /*[1]*/, "ChiShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[5 /*[0]*/, "缪·石印" /*[1]*/, "MiuShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[6 /*[0]*/, "白·石印" /*[1]*/, "BaiShiYin" /*[2]*/, 0 /*[3]*/, "ShiYin" /*[4]*/],
		[7 /*[0]*/, "风面·剑印" /*[1]*/, "FengJianYin" /*[2]*/, 1 /*[3]*/, "JianYin" /*[4]*/],
		[8 /*[0]*/, "林面·剑印" /*[1]*/, "LinJianYin" /*[2]*/, 1 /*[3]*/, "JianYin" /*[4]*/],
		[9 /*[0]*/, "雷面·剑印" /*[1]*/, "LeiJianYin" /*[2]*/, 1 /*[3]*/, "JianYin" /*[4]*/],
		[10 /*[0]*/, "山面·剑印" /*[1]*/, "ShanJianYin" /*[2]*/, 1 /*[3]*/, "JianYin" /*[4]*/],
		[11 /*[0]*/, "火面·剑印" /*[1]*/, "HuoJianYin" /*[2]*/, 1 /*[3]*/, "JianYin" /*[4]*/],
		[12 /*[0]*/, "兰龙·王印" /*[1]*/, "LanLongWangYin" /*[2]*/, 2 /*[3]*/, "LanLongWangYin" /*[4]*/],
		[13 /*[0]*/, "白虎·王印" /*[1]*/, "BaiHuWangYin" /*[2]*/, 2 /*[3]*/, "BaiHuWangYin" /*[4]*/],
		[14 /*[0]*/, "景莲·王印" /*[1]*/, "JingLianWangYin" /*[2]*/, 2 /*[3]*/, "JingLianWangYin" /*[4]*/],
		[15 /*[0]*/, "黑鹰·王印" /*[1]*/, "HeiYingWangYin" /*[2]*/, 2 /*[3]*/, "JianYin" /*[4]*/]
	],
};
