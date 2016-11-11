var mission_level_drama_data = {
		/**
	 * 0 : id, int, 数据ID
	 * 1 : name, varchar, 剧情名称
	 * 2 : mission_level_id, int, 任务区域关卡ID
	 * 3 : quest_id, smallint, 关联任务ID
	 * 4 : quest_state, tinyint, 任务状态
	 * 5 : area_x, smallint, 区域X坐标
	 * 6 : area_y, smallint, 区域Y坐标
	 * 7 : area_width, smallint, 区域
	 * 8 : area_height, smallint, 区域 
	 */

	Id : 0,
	Name : 1,
	Mission_level_id : 2,
	Quest_id : 3,
	Quest_state : 4,
	Area_x : 5,
	Area_y : 6,
	Area_width : 7,
	Area_height : 8,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		
	],
};
