var sword_soul_type_data = {
		/**
	 * 0 : id, smallint, 类型ID
	 * 1 : name, varchar, 类型名称
	 * 2 : sign, varchar, 程序标示 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "攻击" /*[1]*/, "ATTACK " /*[2]*/],
		[2 /*[0]*/, "防御" /*[1]*/, "DEFENCE " /*[2]*/],
		[3 /*[0]*/, "生命" /*[1]*/, "HEALTH " /*[2]*/],
		[4 /*[0]*/, "速度" /*[1]*/, "SPEED " /*[2]*/],
		[5 /*[0]*/, "内力" /*[1]*/, "CULTIVATION " /*[2]*/],
		[6 /*[0]*/, "命中" /*[1]*/, "HIT_LEVEL " /*[2]*/],
		[7 /*[0]*/, "暴击" /*[1]*/, "CRITICAL_LEVEL" /*[2]*/],
		[8 /*[0]*/, "格挡" /*[1]*/, "BLOCK_LEVEL " /*[2]*/],
		[9 /*[0]*/, "破击" /*[1]*/, "DESTROY_LEVEL " /*[2]*/],
		[10 /*[0]*/, "韧性" /*[1]*/, "TENACITY_LEVEL" /*[2]*/],
		[11 /*[0]*/, "闪避" /*[1]*/, "DODGE_LEVEL" /*[2]*/],
		[12 /*[0]*/, "护甲" /*[1]*/, "SUNDER_MAX_VALUE" /*[2]*/],
		[13 /*[0]*/, "剑心经验" /*[1]*/, "EXP" /*[2]*/]
	],
};
