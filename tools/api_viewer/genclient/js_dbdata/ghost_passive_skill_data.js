var ghost_passive_skill_data = {
		/**
	 * 0 : id, bigint, 主键ID
	 * 1 : name, varchar, 被动技名称
	 * 2 : sign, varchar, 图标标识
	 * 3 : desc, varchar, 描述
	 * 4 : star, tinyint, 星级 
	 */

	Id : 0,
	Name : 1,
	Sign : 2,
	Desc : 3,
	Star : 4,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "护盾1级" /*[1]*/, "HunShiShouHu" /*[2]*/, "当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的25%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次" /*[3]*/, 2 /*[4]*/],
		[3 /*[0]*/, "护盾2级" /*[1]*/, "HunShiShouHuErJi" /*[2]*/, "当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的50%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次" /*[3]*/, 3 /*[4]*/],
		[5 /*[0]*/, "护盾3级" /*[1]*/, "HunShiShouHuSanJi" /*[2]*/, "当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的75%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次" /*[3]*/, 4 /*[4]*/],
		[6 /*[0]*/, "护盾4级" /*[1]*/, "HunShiShouHuSiJi" /*[2]*/, "当生命少于30%时，自动施放吸收伤害的魂侍护盾，持续3回合。魂侍护盾吸收的伤害等于魂侍生命值的100%，与其他魂侍合计。魂侍护盾每次战斗只能触发1次" /*[3]*/, 5 /*[4]*/]
	],
};