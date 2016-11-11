<?php

# 倒地起身
define("ARISEN", "arisen");
# 攻击
define("ATTACK", "attack");
# 被击
define("BEATEN", "beaten");
# 格挡
define("BLOCK", "block");
# 鞠躬
define("BOW", "bow");
# 抱宝箱
define("CLASP", "clasp");
# 画
define("DRAW", "draw");
# 趴
define("GROVEL", "grovel");
# 守卫
define("GUARD", "guard");
# 操作
define("HANDLE", "handle");
# 城镇待机
define("IDLE", "idle");
# 跳跃
define("JUMP", "jump");
# 踢
define("KICK", "kick");
# 跪
define("KNEEL", "kneel");
# 笑
define("LAUGH", "laugh");
# 躺
define("LIE", "lie");
# 躺着
define("LYING", "lying");
# 采矿
define("MINING", "mining");
# 跑
define("RUN", "run");
# 坐
define("SIT", "sit");
# 坐2
define("SIT2", "sit2");
# 坐3
define("SIT3", "sit3");
# 技能
define("SKILL", "skill");
# 技能
define("SKILL2", "skill2");
# 蹲下
define("SQUAT", "squat");
# 战场待机
define("STANDBY", "standby");
# 挣扎
define("STRUGGLE", "struggle");
# 破甲
define("SUNDER", "sunder");
# 拜师
define("TRAINEE", "trainee");
# 倒地
define("TUMBLE", "tumble");
# 法术
define("MAGIC", "magic");
# 走
define("WALK", "walk");
# 胜利
define("WIN", "win");

define("BOTTOM", "bottom");
define("LEFT_BOTTOM", "left_bottom");
define("LEFT_TOP", "left_top");
define("LEFT", "left");
define("RIGHT_BOTTOM", "right_bottom");
define("RIGHT_TOP", "right_top");
define("RIGHT", "right");
define("TOP", "top");

$in_map = array(IDLE, RUN, WALK);
$in_war = array(ATTACK, BEATEN, BLOCK, GUARD, LAUGH, MAGIC, SKILL, SKILL2, STANDBY, SUNDER, WIN);
$in_drama = array(ARISEN, BOW, CLASP, DRAW, GROVEL, HANDLE, JUMP, KICK, KNEEL, LIE, LYING, MINING, SIT, SIT2, SIT3, SQUAT, STRUGGLE, TRAINEE, TUMBLE);


###

$current_path = dirname(__FILE__)."/";

$path = dirname(__FILE__).'/';
include_once($path."/config.php");

$client_path = dirname(dirname($path)).'/';
$assets_path = $client_path.'client_mobile/assets/';

$assets_town_path    = $assets_path."map/town/";
$assets_mission_path = $assets_path."map/mission/";
$assets_ghost_mission_path = $assets_path."map/ghost_mission/";

$assets_drama_path = $assets_path."drama/";
$flv_path          = $assets_drama_path."flv/";
$icon_path         = $assets_drama_path."icon/";
$mc_path           = $assets_drama_path."mc/";
$script_path       = $assets_drama_path."script/";
$ui_path           = $assets_drama_path."ui/";

$dbi = new DBI();
$dbi->connect();

$players = array(
	array("义峰", "YiFeng"),
	array("昕苒", "XinRan"),
	array("朱媛媛", "ZhuYuanYuan"),
	array("袁铭志", "YuanMingZhi"),
	array("车晓芸", "CheXiaoYun"),
	array("燕无名", "YanWuMing"),
);

$npcs = array(
    array("村长", "CunZhang"),
);

$monsters = array(
	array("草妖", "Grass"),
);

$npcs = collect_npc();
$monsters = collect_monster();

$bubble_icons = array(
	/*
	"BaiXianEr",
	"CunZhang",
	"DaShou",
	"DuFangLaoBan",
	"LuRenJia",
	"NianYouYiFeng",
	"PingMinNan",
	"PingMinNv",
	"QianQian",
	"QiGai",
	"YiFeng",
	"YiYeShu",
	*/
);
for ($i = 0; $i < count($players); $i++) {
	if (array_key_exists($players[$i][1], collect_bubble_icon())) {
		array_push($bubble_icons, $players[$i][1]);
	}
}
for ($i = 0; $i < count($npcs); $i++) {
	if (array_key_exists($npcs[$i][1], collect_bubble_icon())) {
		array_push($bubble_icons, $npcs[$i][1]);
	}
}
for ($i = 0; $i < count($monsters); $i++) {
	if (array_key_exists($monsters[$i][1], collect_bubble_icon())) {
		array_push($bubble_icons, $monsters[$i][1]);
	}
}

$talk_icons = array(
);

$uis = array(
    "talk",
    "bubble"
);

$emotions = array(
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
	61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
);

$mcs = array(
	array("Q. 苏州奇遇 - 使用入梦符", "Y - 影界/shi_yong_ru_meng_fu"),
	array("Q. 第一个试炼－使用符箓", "Y - 影界/shi_yong_fu_lu"),
	array("M. 追踪乞丐 - 拾取魂力", "Y - 影界/shi_qu_hun_li"),
	array("Z. 紫色龙珠", "Y - 影界/紫色龙珠"),
	
	/*
	# 任务
	array("Q. 青竹村－义峰惊醒", "yi_feng_jing_xing"),
	array("Q. 青竹村－昕冉惊醒", "xin_ran_jing_xing"),
	array("Q. 回村－屋顶(男)", "wu_ding_male"),
	array("Q. 回村－屋顶(女)", "wu_ding_female"),
	array("Q. 赴约 - 叶开剑虚影", "jian_xu_ying"),
	array("Q. 赴约 - 义峰聚火", "ju_huo_male"),
	array("Q. 赴约 - 义峰聚火失败", "ju_huo_shi_bai_male"),
	array("Q. 赴约 - 主角试炼(男)", "shi_lian_male"),
	array("Q. 赴约 - 主角试炼(女)", "shi_lian_female"),
	array("Q. 回城 - 莲花峰(男)", "lian_hua_feng_male"),
	array("Q. 回城 - 莲花峰(女)", "lian_hua_feng_female"),
	array("Q. 苏州奇遇 - 使用入梦符", "shi_yong_ru_meng_fu"),
	array("Q. 金翅鸟 - 云梦山顶(男)", "jin_chi_niao_male"),
	array("Q. 金翅鸟 - 云梦山顶(女)", "jin_chi_niao_male"),
	array("Q. 第一个试炼－使用符箓", "shi_yong_fu_lu"),
	array("Q. 野猪王 - 义峰背朱媛媛(男)", "ye_zhu_wang_male"),
	array("Q. 野猪王 - 昕苒扶义峰(女)", "ye_zhu_wang_female"),
	array("Q. 金翅鸟 - 云梦山顶(男)", "yun_meng_shan_ding_male"),
	array("Q. 金翅鸟 - 云梦山顶(女)", "yun_meng_shan_ding_female"),
	array("Q. 打败血奴-徐福黄一孟(男)", "da_bai_xue_nu_male"),
	array("Q. 打败血奴-徐福黄一孟(女)", "da_bai_xue_nu_female"),
	array("Q. 昆墟内部 - 吸取内容", "xu_qu_nei_rong"),
	array("Q. 打败奸奇 - 大能量刀刃", "da_neng_liang_dao_ren"),
	array("Q. 打败奸奇 - 浣碧变石像", "huan_bi_bian_shi_xiang"),
	array("Q. 打败神庙亡魂 -复活", "fu_huo"),
	array("Q. 打败神庙卫士 - 木听涛被困开始", "mu_ting_tao_bei_kun_kai_shi"),
	array("Q. 打败神庙卫士 - 木听涛被困循环", "mu_ting_tao_bei_kun_xun_huan"),
	array("Q. 打败神庙卫士 - 木听涛被困解除", "mu_ting_tao_bei_kun_jie_chu"),
	array("Q. 毒兽 - 打败血杀军(男)", "da_bai_xue_sha_jun_male"),
	array("Q. 毒兽 - 打败血杀军(女)", "da_bai_xue_sha_jun_female"),
	array("Q. 皇城屋顶 - 皇城屋顶(男)", "huang_cheng_wu_ding_male"),
	array("Q. 皇城屋顶 - 皇城屋顶(女)", "huang_cheng_wu_ding_female"),
	array("Q. 成长 - 冲破", "chong_po"),
	array("Q. 右翼船阵 - 海上风暴", "hai_shang_feng_bao"),
	array("Q. 虚天殿 - 龙脉之危", "long_mai_zhi_wei"),
	array("Q. 枯木虫巢 - 全体治疗", "quan_ti_zhi_liao"),
	array("Q. 血咒死士 - 天剑", "xue_zhou_si_shi"),
	array("Q. 大福船 - 义峰", "da_fu_chuan_male"),
	array("Q. 大福船 - 昕冉", "da_fu_chuan_female"),
	array("Q. 大战后 - 义峰", "da_zhan_hou_male"),
	array("Q. 大战后 - 昕冉", "da_zhan_hou_female"),
	array("Q. 徐福 - 神秘黑衣人", "shen_mi_hei_yi_ren"),
	array("Q. 徐福 - 爆炸", "bao_zha"),
	array("Q. 古月师太 - 拾取画卷", "shi_qu_hua_juan"),
	
	# 副本
	array("M. 遗失的猎刀 - 神秘猎刀", "yi_shi_de_lie_dao"),
	array("M. 追踪乞丐 - 闪电", "shan_dian"),
	array("M. 追踪乞丐 - 拾取命格", "shi_qu_ming_ge"),
	array("M. 追踪乞丐 - 拾取魂力", "shi_qu_hun_li"),
	array("M. 获取冰露 - 获取冰露", "shi_qu_bing_lu"),
	array("M. 古井惊魂 - 古井惊魂", "gu_jing"),
	array("M. 打败螭吻 - 水珠", "shui_zhu"),
	array("M. 打败毒蝎战将 - 风珠", "feng_zhu"),
	array("M. 火麒麟 - 火珠", "huo_zhu"),
	array("M. 冥界帝皇 - 土珠", "tu_zhu"),
	array("M. 雷之试炼 - 雷珠", "lei_zhu"),
	array("M. 开场剧情", "kai_chang_ju_qing"),
	*/
);

$flvs = array(
	/*
	array("三年后(男)", "learn_male"),
	array("三年后(女)", "learn_female"),
	*/
);

$enters = array(
	array("剧情传送", "ju_qing_chuan_song"),
	array("怪物剧情传送", "guai_wu_ju_qing_chuan_song"),
);

$leaves = array(
	array("黄一孟消失", "huang_yi_meng"),
);

# 以 drama_ 开头 才会显示在场景上
$puppets = array(
	/*
	# 任务
	
	array("Q. 打败浣碧 - 石像", "drama_shi_xiang"),
	array("Q. 地牢 - 牢笼", "drama_lao_long"),
	array("Q. 成长 - 冲破大门", "drama_chong_po_da_men"),
	array("Q. 打败螭吻 - 水珠", "drama_shui_zhu"),
	array("Q. 打败毒蝎战将 - 风珠", "drama_feng_zhu"),
	array("Q. 火麒麟 - 火珠", "drama_huo_zhu"),
	array("Q. 雷之试炼 - 雷珠", "drama_lei_zhu"),
	array("Q. 龙脉之危 - 封印", "drama_feng_yin"),
	array("Q. 龙脉之危 - 吸取龙脉之力", "drama_long_mai_zhi_li"),
	array("Q. 冥界帝皇 - 土珠", "drama_tu_zhu"),
	array("Q. 打败神庙卫士 - 木听涛被困循环", "drama_mu_ting_tao_bei_kun"),
	array("Q. 昆墟内部 - 吸取内容", "drama_xi_qu_nei_rong"),
	array("Q. 进副本的传送阵", "drama_enter_mission_portal"),
	array("Q. 白咒死士 - 吸取龙脉", "drama_xi_qu_long_mai"),
	array("Q. 百花园 - 紫色烟雾", "drama_zi_se_yan_wu"),
	array("Q. 冰窟 - 十绝冰阵(身体)", "drama_shi_jue_bing_zhen_body"),
	array("Q. 冰窟 - 十绝冰阵(底)", "drama_shi_jue_bing_zhen_shadow"),
	array("Q. 江户城 - 樱花飘落", "drama_ying_hua_piao_luo"),
	array("Q. 隧道 - 毒气", "drama_du_qi"),
	array("Q. 赶跑草妖 - 坟墓", "drama_fen_mu"),
	array("Q. 血咒死士 - 落石", "drama_luo_shi"),
	array("Q. 血杀偷袭 - 海浪", "drama_hai_lang"),
	array("Q. 天灯之愿 - 密道入口", "drama_mi_dao_ru_kou"),
	array("Q. 天灯之愿 - 天灯", "drama_tian_deng_1"),
	array("Q. 无名的梦境 - 大树", "drama_da_shu"),
	array("Q. 无名的梦境 - 福字碑", "drama_fu_zi_bei"),
	array("Q. 无名的梦境 - 烘托气氛", "drama_hong_tuo_qi_fen"),
	array("Q. 无名的梦境 - 天灯", "drama_tian_deng"),
	
	# 影界
	
	array("Y. 影界 - 魂侍之卵", "drama_ghost_egg"),
	array("Y. 影界 - 魂侍之卵(金)", "drama_ghost_gold_egg"),
	
	# 副本
	
	# 宝藏关卡
	
	array("T. 宝藏关卡 - 铜钱", "treasure_coin"),
	array("T. 宝藏关卡 - 宝箱（右下）", "treasure_box_right_bottom"),
	array("T. 宝藏关卡 - 宝箱（左下）", "treasure_box_left_bottom"),
	array("T. 宝藏关卡 - 罐子（右下）", "treasure_jar_right_bottom"),
	array("T. 宝藏关卡 - 罐子（左下）", "treasure_jar_left_bottom"),
	array("T. 宝藏关卡 - 石堆（右下）", "treasure_stones_right_bottom"),
	array("T. 宝藏关卡 - 石堆（左下）", "treasure_stones_left_bottom"),
	array("T. 宝藏关卡 - 雪人（右下）", "treasure_snowman_right_bottom"),
	array("T. 宝藏关卡 - 雪人（左下）", "treasure_snowman_left_bottom"),
	array("T. 宝藏关卡 - 枯骨（右下）", "treasure_bones_right_bottom"),
	array("T. 宝藏关卡 - 枯骨（左下）", "treasure_bones_left_bottom"),
	array("T. 宝藏关卡 - 宝箱1（右上）", "treasure_box1_right_top"),
	array("T. 宝藏关卡 - 宝箱1（左上）", "treasure_box1_left_top"),
	array("T. 宝藏关卡 - 宝箱2（右下）", "treasure_box2_right_bottom"),
	array("T. 宝藏关卡 - 宝箱2（左下）", "treasure_box2_left_bottom"),
	array("T. 宝藏关卡 - 宝箱3（右下）", "treasure_box3_right_bottom"),
	array("T. 宝藏关卡 - 宝箱3（左下）", "treasure_box3_left_bottom"),
        array("T. 宝藏关卡 - 花苞", "treasure_flower_common"),
        array("T. 宝藏关卡 - 树枝堆", "treasure_shuzhidui_common"),
        array("T. 宝藏关卡 - 碑林匣子（右下）", "treasure_beilinxiazi_right_bottom"),
	array("T. 宝藏关卡 - 碑林匣子（左下）", "treasure_beilinxiazi_left_bottom"),
        
	# 特殊伙伴
	
	array("S. 一杯黄粱 - 告示牌", "drama_bulletin"),
	*/
);

$blank_screens = array(
	array("白天", "day"),
	array("黑夜", "night"),
	array("一小时", "a_hour"),
	array("黑夜到白天", "day_and_night"),
);

$effect_ons = array(
	"wars" => array(
		#array("防御免伤", "FangYuMianShangB/master_body"),
	),
	"dramas" => array(
		array("睡觉", "shui_jiao"),
		array("迷惑", "mi_huo"),
	)
);

$drama = array(
	"town"    => collect_town(),#$town,
	"mission" => collect_mission(),#$mission,
	"ghost_mission" => collect_ghost_mission(),
	
	"players"  => $players,
	"npcs"     => collect_npc(),#$npcs,
	"monsters" => collect_monster(),#$monsters,
	
	"actions" => collect_actions(),
	
	"bubble_icons" => $bubble_icons,
	"talk_icons"   => $talk_icons,
	
	"uis"      => $uis,
	"emotions" => $emotions,
	
	"mcs" => $mcs,
	"flvs" => $flvs,
	
	"skills" => collect_skill(),
	"musics" => collect_music(),
	
	"enters" => $enters,
	"leaves" => $leaves,
	
	"puppets" => $puppets,
	
	"blank_screens" => $blank_screens,
	
	"effect_ons" => $effect_ons,
);
#var_export($drama);
file_put_contents($assets_drama_path."drama.txt", json_encode($drama));

file_put_contents($assets_path."roles/skills/config.txt", gzcompress(json_encode($drama["skills"])));

/**
 * 收集数据
 */

function collect_town () {
	global $dbi;
	global $assets_town_path;
	
	$towns = array();
	
	$list = $dbi->query("select * from `town` order by `lock`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_town_path.$list[$i]["sign"])) continue;
		if (! $list[$i]["sign"]) continue;
		
		array_push(
			$towns,
			array(
				"name" => $list[$i]["name"],
				"sign" => $list[$i]["sign"],
				"level" => 1,
			)
		);
	}
	
	return $towns;
}

function collect_mission () {
	global $dbi;
	global $assets_mission_path;
	
	$missions = array();
	
	$list = $dbi->query("select * from `mission_level` order by `lock`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_mission_path.$list[$i]["sign"])) continue;
		
		array_push(
			$missions,
			array(
				"name" => $list[$i]["name"],
				"sign" => $list[$i]["sign"],
				"level" => 1,
			)
		);
	}
	
	/*
	$list = $dbi->query("select * from `mission` order by `order`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_mission_path.$list[$i]["sign"])) continue;
		$mission_id = $list[$i]["id"];
		
		$temp = $dbi->query("select count(*) as `count` from `mission_map` where `mission_id` = ".$mission_id);
		
		array_push(
			$missions,
			array(
				"name" => $list[$i]["name"],
				"sign" => $list[$i]["sign"],
				"level" => $temp[0]["count"],
			)
		);
	}
	*/
	
	return $missions;
}

function collect_ghost_mission () {
	global $dbi;
	global $assets_ghost_mission_path;
	
	$missions = array();
	
	return $missions;
	
	$list = $dbi->query("select `name`, `sign` from `town` where `id` in (select `town_id` from `ghost_mission`) order by `lock`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_ghost_mission_path.$list[$i]["sign"])) continue;
		
		array_push(
			$missions,
			array(
				"name" => $list[$i]["name"],
				"sign" => $list[$i]["sign"],
				"level" => 1,
			)
		);
	}
	
	$list = $dbi->query("select `name`, `sign` from `town` where `id` in (select `town_id` from `ingot_ghost_mission`) order by `lock`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_ghost_mission_path.$list[$i]["sign"])) continue;
		
		array_push(
			$missions,
			array(
				"name" => $list[$i]["name"],
				"sign" => $list[$i]["sign"],
				"level" => 1,
			)
		);
	}
	
	return $missions;
}

function collect_npc () {
	global $dbi;
	global $assets_path;
	
	$npcs = array();
	$signs = array();
	
	$list = $dbi->query("select * from `town_npc`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_path."roles/actions/".$list[$i]["sign"])) continue;
		
		if (array_key_exists($list[$i]["sign"], $signs)) continue;
		$signs[$list[$i]["sign"]] = 1;
		
		array_push(
			$npcs,
			array($list[$i]["name"], $list[$i]["sign"])
		);
	}
	
	return $npcs;
}

function collect_monster () {
	global $dbi;
	global $assets_path;
	
	$monsters = array();
	$signs = array();
	
	$list = $dbi->query("select * from `enemy_role`");
	for ($i = 0; $i < count($list); $i++) {
		if (! is_dir($assets_path."roles/actions/".$list[$i]["sign"])) continue;
		
		// if (array_key_exists($list[$i]["sign"], $signs)) continue;
		$signs[$list[$i]["sign"]] = 1;
		
		array_push(
			$monsters,
			array($list[$i]["name"], $list[$i]["sign"])
		);
	}
	
	return $monsters;
}

/**
 * 收集角色动作
 */
function collect_actions () {
	global $assets_path, $in_war, $in_drama;
	
	$actions_path = $assets_path."roles/actions/";
	#print $actions_path;
	
	$actions = array();
	
	$list = glob($actions_path."*");
	for ($i = 0; $i < count($list); $i++) {
		$sign = basename($list[$i]);
		#print $sign."\n";
		
		$actions[$sign] = array();
		
		$arr = glob($list[$i]."/*.*");
		for ($j = 0; $j < count($arr); $j++) {
			$type = basename($arr[$j], ".swf");
			$types = explode("_", $type);
			
			if (! array_key_exists($types[0], $actions[$sign])) {
				$actions[$sign][$types[0]] = array();
			}
			
			if (in_array($types[0], $in_war) || in_array($types[0], $in_drama))
			{
				if (count($types) == 1) continue;
				
				array_push(
					$actions[$sign][$types[0]],
					count($types) == 2 ? $types[1] : $types[1]."_".$types[2]
				);
			}
			else {
				$actions[$sign][$types[0]] = "all";
			}
		}
		#var_export($arr);
		
		#exit();
	}
	
	#var_export($actions);
	return $actions;
}

/**
 * 扫瞄剧情头像
 */
function collect_bubble_icon () {
	global $assets_path;
	
	static $bubbles;
	
	if (! isset($bubbles)) {
		$bubbles = array();
		
		$list = glob($assets_path."drama/bubble_icon/*.swf");
		for ($i = 0; $i < count($list); $i++) {
			$bubbles[basename($list[$i], ".swf")] = 1;
		}
	}
	
	return $bubbles;
}

/**
 * 收集绝招
 */
function collect_skill () {
	global $dbi, $assets_path;
	
	$hash = array();
	
	$list = $dbi->query("select `name`, `sign` from `skill`");
	for ($i = 0; $i < count($list); $i++) {
        $sign = $list[$i]["sign"];
		$letter = substr($sign, strlen($sign) - 1);
        
		$hash[$sign] = $list[$i]["name"].'('.($letter === "M" ? "怪物" : ($letter === "B" ? "伙伴" : "主角")).')';
        
        #if (! is_dir($assets_path."roles/skills/$sign"))
        #    mkdir($assets_path."roles/skills/$sign");
	}
	
	###
	
	$skills = array();
	
	$list = glob($assets_path."roles/skills/*");
	#var_export($list);
	
	$len = count($list);
	for ($i = 0; $i < $len; $i++) {
		$item = $list[$i];
		
		if (is_dir($item)) {
			$sign = basename($item);
			
			if ($sign == "Block"
                || $sign == "BlockForChenAMing"
                || $sign == "RoleDeadL"
                || $sign == "RoleDeadM"
                || $sign == "RoleDeadS") {
				continue;
            }
			
            $has = array_key_exists($sign, $hash);
			$skills[$sign] = array($has ? $hash[$sign] : $sign);
			
			$arr = glob($item."/*.swf");
			$len1 = count($arr);
			for ($j = 0; $j < $len1; $j++) {
				$basename = basename($arr[$j], ".swf");
				array_push($skills[$sign], $basename);
			}
            
            # 后台是否录入
            #if (! $has)
            #    print "[not exists] ".$sign."\n";
		}
		else {
			#array_push($item);
		}
	}
	#var_export($skills);
	return $skills;
}

/**
 * 收集音乐
 */
function collect_music () {
	global $assets_path;
	
	$musics = array();
	
	$arr = glob($assets_path."sound/*.mp3");
	for ($i = 0; $i < count($arr); $i++) {
		array_push($musics, basename($arr[$i], ".mp3"));
	}
	
	return $musics;
}

// =============================================================================
//
// 角色动作
//
// =============================================================================

$actions = array();

$names = array();

$dbi = new DBI();
$dbi->connect();

getName("role");
getName("enemy_role");
getName("town_npc");

$dbi->close();

function getName ($table) {
	global $dbi, $names;
	
	$list = $dbi->query("select `name`, `sign` from `$table`");
	for ($i = 0; $i < count($list); $i++) {
		$sign = $list[$i]["sign"];
		if (! array_key_exists($sign, $names))
			$names[$sign] = array();
		
		if (! in_array($list[$i]["name"], $names[$sign]))
			$names[$sign][] = $list[$i]["name"];
	}
}

/**
 * 收集角色城镇动作
 */
function collect_role_actions () {
	collect_roles_actions("roles/actions/", TRUE);
}

/**
 * 收集角色战场动作
 */
function collect_roles_actions_war () {
	collect_roles_actions("roles/actions_war/", FALSE);
}

function collect_roles_actions ($path, $is_town) {
	global $assets_path, $actions, $names;
	$arr = glob($assets_path.$path."*");
	
	for ($i = 0;  $i < count($arr); $i++) {
		if (! is_dir($arr[$i])) continue;
		
		$sign = basename($arr[$i]);
		
		if (! array_key_exists($sign, $actions)) {
			$actions[$sign] = array(
				"name" => array_key_exists($sign, $names) ? join(",", $names[$sign]) : "",
			);
		}
		$temp = $actions[$sign][$is_town ? "town" : "war"] = array();
		
		$list = glob($arr[$i]."/*");
		for ($j = 0; $j < count($list); $j++) {
			$actions[$sign][$is_town ? "town" : "war"][] = basename($list[$j], ".swf");
		}
	}
}

collect_role_actions();
collect_roles_actions_war();

file_put_contents($assets_path."roles/actions_config.txt", gzcompress(json_encode($actions)));

?>