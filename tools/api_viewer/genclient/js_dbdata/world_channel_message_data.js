var world_channel_message_data = {
		/**
	 * 0 : id, smallint, 消息模版ID
	 * 1 : parameters, varchar, 参数
	 * 2 : content, varchar, 内容 
	 */

	Id : 0,
	Parameters : 1,
	Content : 2,

	Each: function(cb) {
		for(var i = 0; i < this._list.length; i++) {
			cb(this._list[i]);
		}
	},

	_list: [
		[1 /*[0]*/, "player,player,玩家;item,item,道具" /*[1]*/, "{0,#55c4dd}在<font color=#c503fd >【剑山拔剑】</font>中提手拔剑获得了{1,#fcf300}" /*[2]*/],
		[2 /*[0]*/, "player,player,玩家;item,item,道具" /*[1]*/, "{0,#55c4dd}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1,#fcf300}" /*[2]*/],
		[3 /*[0]*/, "player,player,玩家;item,item,道具" /*[1]*/, "{0,#55c4dd}在<font color=#c503fd >【命锁】</font>中好运连连获得了{1,#fcf300}" /*[2]*/],
		[4 /*[0]*/, "player,player,玩家;string,level,关卡;item,item,道具" /*[1]*/, "{0,#fcf300}在{1,#fcf300}中突破挑战获得了金色魂侍{2,#fcf300}" /*[2]*/],
		[5 /*[0]*/, "player,player,玩家;item,item,道具" /*[1]*/, "{0,#fcf300}经过了长期累积获得了金色魂侍{1,#fcf300}" /*[2]*/],
		[6 /*[0]*/, "player,player,玩家;item,item,道具" /*[1]*/, "{0,#55c4dd}使用<font color=#c503fd >【圣器召唤】</font>召唤出守护阵印{1,#fcf300}" /*[2]*/],
		[8 /*[0]*/, "player,player,玩家;clique,clique,帮派;clique,dummy_link,点击申请" /*[1]*/, "{0,#55c4dd}成立了帮派【{1,#fcf300}】，欢迎各位侠士道友前去拜会，{2,#fcf300}" /*[2]*/],
		[9 /*[0]*/, "clique,clique,帮派;clique,dummy_link,点击申请" /*[1]*/, "帮派【{0,#55c4dd}】准备大展宏图！欢迎各位侠士道友加入帮派，{1,#fcf300}" /*[2]*/],
		[10 /*[0]*/, "player,player,玩家" /*[1]*/, "{0,#fcf300}被提升为副帮主" /*[2]*/],
		[11 /*[0]*/, "player,old_owner,原帮主;player,new_owner,新帮主" /*[1]*/, "{0,#fcf300}将帮主的位置转让给了{1,#fcf300}" /*[2]*/],
		[12 /*[0]*/, "player,new_owner,新班主;player,player2,原帮主" /*[1]*/, "{0,#fcf300}弹劾了{1,#fcf300}成为了新帮主" /*[2]*/],
		[13 /*[0]*/, "player,player,玩家" /*[1]*/, "{0,#fcf300}离开了帮派" /*[2]*/],
		[14 /*[0]*/, "player,manger,管理员;player,member,成员" /*[1]*/, "{0,#fcf300}将{1,#fcf300}踢出了帮派" /*[2]*/],
		[15 /*[0]*/, "" /*[1]*/, "帮派公告已更新，请注意查看" /*[2]*/],
		[16 /*[0]*/, "string,content,内容" /*[1]*/, "{0}" /*[2]*/],
		[17 /*[0]*/, "player,manger,管理员" /*[1]*/, "{0,#fcf300}被解除了副帮主的职位" /*[2]*/],
		[18 /*[0]*/, "player,new_member,新成员" /*[1]*/, "{0,#fcf300}加入了帮派" /*[2]*/],
		[19 /*[0]*/, "string,build_type,帮派建筑名称;string,clique_level,帮派等级" /*[1]*/, "帮派{0,#fcf300}已升级至 {1,#fcf300} 级" /*[2]*/],
		[20 /*[0]*/, "clique,clique,敌对帮派;player,hijacker,劫持者;player,boat_owner,我帮玩家;boat,dummy_link,镖船" /*[1]*/, "{0,#55c4dd}的{1,#fcf300}夺取了本帮{2,#fcf300}的镖船，{3,#ff0000}" /*[2]*/],
		[21 /*[0]*/, "player,fighter,夺回者;player,boat_owner,镖船主" /*[1]*/, "本帮{0,#fcf300}仗义出手夺回了{1,#fcf300}的镖船" /*[2]*/],
		[22 /*[0]*/, "player,fighter,夺回者" /*[1]*/, "很遗憾{0,#fcf300}击败了你夺回了镖船" /*[2]*/],
		[23 /*[0]*/, "player,fighter,夺回者" /*[1]*/, "{0,#fcf300}仗义出手帮你夺回了镖船" /*[2]*/],
		[24 /*[0]*/, "string,coins,铜钱;string,fame,声望;string,contrib,贡献" /*[1]*/, "恭喜你夺取成功获得{0,#6b6055}铜钱{1,#6b6055}声望{2,#6b6055}帮贡" /*[2]*/],
		[25 /*[0]*/, "string,coins,铜钱;string,fame,声望;string,contrib,贡献" /*[1]*/, "恭喜你护送运镖成功获得{0,#6b6055}铜钱{1,#6b6055}声望{2,#6b6055}帮贡" /*[2]*/],
		[26 /*[0]*/, "player,hijacker,劫持者;string,coins,铜钱;string,fame,声望;string,contrib,贡献" /*[1]*/, "{0,#fcf300}夺取了镖船损失{1,#6b6055}铜钱{2,#6b6055}声望{3,#6b6055}帮贡" /*[2]*/],
		[27 /*[0]*/, "clique,hijacker_clique,劫持者帮派;player,hijacker,劫持者" /*[1]*/, "{0,#55c4dd}的{1,#fcf300}夺取了你的镖船快去找人帮忙" /*[2]*/],
		[28 /*[0]*/, "player,name,玩家名称;string,mission_name,区域名称;string,level_name,关卡名称" /*[1]*/, "【{0,#fcf300}】经过一番激战，通关【{1}】的【{2}】，为讨伐绝望之地做出杰出贡献" /*[2]*/],
		[29 /*[0]*/, "player,name,玩家名称;string,mission_name,区域名称;string,level_name,关卡名称" /*[1]*/, "【{0,#fcf300}】大显神威，三星通关【{1}】的【{2}】，为讨伐绝望之地做出杰出贡献" /*[2]*/],
		[30 /*[0]*/, "string,job,帮派职位;player,name,发放者;string,item_name,战备类型" /*[1]*/, "福利派送！{0}{1,#ffd200}为大家发放了{2}，请查收邮箱" /*[2]*/],
		[31 /*[0]*/, "string,camp,阵营;string,boss,首领" /*[1]*/, "恭喜各位大侠成功击败绝望之地{0}军团首领【{1}】！" /*[2]*/]
	],
};
