var events_buy_partner_data = {
		/**
	 * 0 : id, int, 
	 * 1 : patner_id, smallint, 伙伴ID
	 * 2 : buddy_level, smallint, 伙伴等级
	 * 3 : cost, bigint, 价格
	 * 4 : skill_id1, smallint, 招式名称1
	 * 5 : skill_id2, smallint, 招式名称2 
	 */

	Id : 0,
	Patner_id : 1,
	Buddy_level : 2,
	Cost : 3,
	Skill_id1 : 4,
	Skill_id2 : 5,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, 7 /*[1]*/, 40 /*[2]*/, 3000 /*[3]*/, 1300 /*[4]*/, 1301 /*[5]*/]
	],
};
