var clique_boat_data = {
		/**
	 * 0 : id, smallint, 主键
	 * 1 : name, varchar, 镖船名称
	 * 2 : sign, varchar, 资源标识
	 * 3 : rate, tinyint, 概率
	 * 4 : escort_time, smallint, 运送时间（单位分钟）
	 * 5 : select_cost_ingot, bigint, 直接选择话费元宝（0则不可直接选择）
	 * 6 : award_coins, bigint, 奖励铜钱
	 * 7 : award_fame, int, 奖励声望
	 * 8 : award_clique_contrib, bigint, 奖励贡献
	 * 9 : hijack_lose_coins, bigint, 抢劫损失铜钱
	 * 10 : hijack_lose_fame, int, 抢劫损失声望
	 * 11 : hijack_lose_clique_contrib, bigint, 抢劫损失贡献
	 * 12 : color, varchar, 颜色 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Rate : 3,
	Escort_time : 4,
	Select_cost_ingot : 5,
	Award_coins : 6,
	Award_fame : 7,
	Award_clique_contrib : 8,
	Hijack_lose_coins : 9,
	Hijack_lose_fame : 10,
	Hijack_lose_clique_contrib : 11,
	Color : 12,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "小货船" /*[1]*/, "XiaoHuoChuan" /*[2]*/, 30 /*[3]*/, 20 /*[4]*/, 0 /*[5]*/, 5000 /*[6]*/, 100 /*[7]*/, 100 /*[8]*/, 2500 /*[9]*/, 50 /*[10]*/, 100 /*[11]*/, "#ffffff" /*[12]*/],
		[2 /*[0]*/, "货船" /*[1]*/, "HuoChuan" /*[2]*/, 30 /*[3]*/, 20 /*[4]*/, 0 /*[5]*/, 10000 /*[6]*/, 200 /*[7]*/, 200 /*[8]*/, 5000 /*[9]*/, 100 /*[10]*/, 200 /*[11]*/, "#56ca00" /*[12]*/],
		[3 /*[0]*/, "商船" /*[1]*/, "ShangChuan" /*[2]*/, 30 /*[3]*/, 20 /*[4]*/, 0 /*[5]*/, 15000 /*[6]*/, 300 /*[7]*/, 300 /*[8]*/, 7500 /*[9]*/, 150 /*[10]*/, 300 /*[11]*/, "#027eca" /*[12]*/],
		[4 /*[0]*/, "大商船" /*[1]*/, "DaShangChuan" /*[2]*/, 10 /*[3]*/, 20 /*[4]*/, 0 /*[5]*/, 25000 /*[6]*/, 400 /*[7]*/, 400 /*[8]*/, 12500 /*[9]*/, 200 /*[10]*/, 400 /*[11]*/, "#ff00f6" /*[12]*/],
		[5 /*[0]*/, "宝船" /*[1]*/, "BaoChuan" /*[2]*/, 0 /*[3]*/, 20 /*[4]*/, 100 /*[5]*/, 90000 /*[6]*/, 3000 /*[7]*/, 3000 /*[8]*/, 30000 /*[9]*/, 1000 /*[10]*/, 3000 /*[11]*/, "#ffd200" /*[12]*/]
	],
};
