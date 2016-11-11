var mission_level_recovery_meng_yao_data = {
		/**
	 * 0 : id, int, 关卡梦妖标示
	 * 1 : mission_level_id, int, 关卡id
	 * 2 : my_x, int, 关卡梦妖x坐标
	 * 3 : my_y, int, 关卡梦妖y坐标
	 * 4 : my_effect, tinyint, 关卡梦妖效果，1-恢复灵宠使用次数 2-恢复全体生命 3-恢复伙伴绝招次数 
	 * 5 : my_dir, varchar, 关卡梦妖朝向 r=>右, rb=>右下方, b=>下, lb=>左下方, l=>左, lt=>左上方, t=>上, rt=>右上方
	 * 6 : talk, text, 梦妖对话内容 
	 */

	Id : 0,
	Mission_level_id : 1,
	My_x : 2,
	My_y : 3,
	My_effect : 4,
	My_dir : 5,
	Talk : 6,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[6 /*[0]*/, 26 /*[1]*/, 1334 /*[2]*/, 1158 /*[3]*/, 2 /*[4]*/, "rb" /*[5]*/, "你似乎需要治疗！" /*[6]*/],
		[7 /*[0]*/, 29 /*[1]*/, 850 /*[2]*/, 1201 /*[3]*/, 3 /*[4]*/, "rb" /*[5]*/, "神龙赐予你力量！" /*[6]*/]
	],
};
