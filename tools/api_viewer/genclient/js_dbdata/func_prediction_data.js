var func_prediction_data = {
		/**
	 * 0 : order, smallint, 顺序
	 * 1 : type, tinyint, 类型 0--等级触发 1--新功能权值触发
	 * 2 : condition_value, smallint, 触发条件 等级 或 功能权值
	 * 3 : sign, varchar, 资源标识
	 * 4 : summary, varchar, 下一功能描述
	 * 5 : tips, varchar, tips 
	 */

	Order : 0,
	Type : 1,
	Condition_value : 2,
	Sign : 3,
	Summary : 4,
	Tips : 5,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		
	],
};
