var ghost_tip_data = {
		/**
	 * 0 : id, smallint, 主键ID
	 * 1 : tip, varchar, 提示信息 
	 */

	Id : 0,
	Tip : 1,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "战斗中，角色每次行动、被击和击杀敌人都会获得魂力\n当魂力满100时，可以召唤主魂侍\n若主魂侍与辅魂侍达成连锁，会同时召唤主、辅魂侍" /*[1]*/],
		[2 /*[0]*/, "每只魂侍在同个关卡内只能被触发一次" /*[1]*/]
	],
};
