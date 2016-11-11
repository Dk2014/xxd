<?php
require_once dirname(__FILE__).'/config.php';
require_once dirname(__FILE__).'/mysql.php';
require_once dirname(__FILE__).'/html_component.php';
require_once dirname(__FILE__).'/../pages/item_links.php';

$db = db_connect();


//
// 检查用户是否登录，如未登录则跳转到登录页面
//
function check_login() {
	session_start();

	if (!isset($_SESSION['user_id'])) {
		Header("Location:login.php");
		exit;
	}
}
function is_in($search, $arr){
	foreach ($arr as $key => $value) {
		if ($value == $search) {
			return $key;
		}
	}
	return false;
}

function get_diff($search, $pconfig, $exclude){
	$result = array();
	foreach ($pconfig['columns'] as $value) {
		if(!in_array($value['name'], $exclude) && !in_array($value['name'], $search)){
			array_push($result, $value['text']);
		}
	}
	return $result;
}

// 截取中文字符
function utf_substr( $str, $len ) {
	for ( $i=0;$i<$len;$i++ ) {
		$temp_str=substr( $str, 0, 1 );
		if ( ord( $temp_str ) > 127 ) {
			$i++;
			if ( $i<$len ) {
				$new_str[]=substr( $str, 0, 3 );
				$str=substr( $str, 3 );
			}
		} else {
			$new_str[]=substr( $str, 0, 1 );
			$str=substr( $str, 1 );
		}
	}
	return join( $new_str );
}

function mission_enemy_tip_render($role) {
	if ($role == null)
		return;

	$html  = "<table>";
	$html .= "<tr><td colspan=\"6\"><b style=\"font-size:16px;\">".$role['name']."</b></td><td></td></tr>";
	$html .= "<tr><td colspan=\"6\"><hr style=\"border:solid 1px #ccc; height:0; border-top:0; border-left:0; border-right:0;\"/></td></tr>";

	$html .= "<tr valign=\"top\">";

	$html .= "<td><table style=\"width:130px; border-right:solid 1px #ccc\">";
	$html .= "<tr><td>等级</td><td>".$role['level']."</td></tr>";
	$html .= "<tr><td>生命</td><td>".$role['health']."</td></tr>";
	$html .= "<tr><td>内力</td><td>".$role['cultivation']."</td></tr>";
	$html .= "<tr><td>速度</td><td>".$role['speed']."</td></tr>";
	$html .= "<tr><td>攻击</td><td>".$role['attack']."</td></tr>";
	$html .= "<tr><td>防御</td><td>".$role['defence']."</td></tr>";
	$html .= "<tr><td>暴击</td><td>".$role['critial']."%</td></tr>";
	$html .= "<tr><td>格挡</td><td>".$role['block']."%</td></tr>";
	$html .= "<tr><td>命中</td><td>".$role['hit']."%</td></tr>";
	$html .= "<tr><td>闪避</td><td>".$role['dodge']."%</td></tr>";
	$html .= "<tr><td>必杀</td><td>".$role['critial_hurt']."%</td></tr>";
	$html .= "</table></td>";

	$html .= "<td><table style=\"width:160px; border-right:solid 1px #ccc\">";
	$html .= "<tr><td>护甲上限</td><td>".$role['sunder_max_value']."</td></tr>";
	$html .= "<tr><td>破甲起始百分比</td><td>".$role['sunder_min_hurt_rate']."</td></tr>";
	$html .= "<tr><td>破甲后百分比</td><td>".$role['sunder_end_hurt_rate']."</td></tr>";
	$html .= "</table></td>";

	$html .= "<td><table style=\"width:200px; border-right:solid 1px #ccc\">";
	$html .= "<tr><td>绝招威力1</td><td>".$role['skill_force']."</td></tr>";
	if ($role['skill_id'] > 0) {
		$html .= "<tr><td>绝招名称1</td><td>".$role['skill_name']."</td></tr>";
	} else {
		$html .= "<tr><td>绝招名称1</td><td>无</td></tr>";
	}
	
	
	$html .= "</table></td>";


	$html .= "<td>";
	$html .= html_get_showFrame($role['sign'], "./mobile-out/角色/monster", "236", "236");
	$html .= "</td>";

	$html .= "</tr>";
	$html .= "</table>";

	return $html;
}

function enemy_role_editor($allEnemyRole, $row, $name) {
	if ($row == NULL) {
		return '
<input class="editor role_editor" style="width:150px" type="text" id="new_role_editor_##id##_'.$name.'" value="无"/>
<input type="hidden" name="'.$name.'[]" id="new_role_editor_##id##_'.$name.'_hidden" value="0" />
';
	}

	return '
<input class="editor role_editor" style="width:150px" type="text" value="'.($row[$name] != 0 ? $allEnemyRole[$row[$name]]['name'] : '无').'" id="role_editor_'.$row['id'].'_'.$name.'"/>
<input type="hidden" name="'.$name.'[]" value="'.$row[$name].'" id="role_editor_'.$row['id'].'_'.$name.'_hidden" />
';
}

function enemy_group_table($allEnemyRole, $row, $sql) {
	global $db;

	$html = '';

	$groups = db_query($db, $sql);

	$allEnemyRole[0] = null;

	$role_exp = 0;
	$coins = 0;

	$gi = 0;
	foreach ($groups as $group) {
		$gi ++;

		$html .= enemy_group_table_one($allEnemyRole, $row, $group, $gi, $role_exp, $coins);
	}


	return $html;
}

function enemy_group_table_one($allEnemyRole, $row, $group, $group_index, &$role_exp, &$coins) {
	$html = '';

	$gi = $group_index;

	$allEnemyRole[0] = null;

	$html .= "<table style=\"width:100px; height:60px; margin:10px 0; background:#fff;\" colspan=\"0\" rowspan=\"0\" cellpadding=\"0\" cellspacing=\"0\">";

	foreach (array(10, 5, 0) as $i) {
		$html .= "<tr>";

		for ($j = 5; $j >= 1; $j --) {
			$id = ($i + $j) < 10 ? ($i + $j) : ''.($i + $j);

			if ($group['pos'.$id] > 0) {
				$role = $allEnemyRole[$group['pos'.$id]];

				$html .= "<td id=\"eg_{$row['id']}_{$gi}_td_$id\" style=\"color:#426DC9;cursor:pointer\">";
				$html .= "<a href=\"?p=enemy_role&m=".$group['pos'.$id]."\" target=\"_blank\">";
				$html .= mb_substr($role['name'], 0, 1, "utf-8");
				$html .= "</a>";
				$html .= "<span style=\"display:none\" id=\"eg_{$row['id']}_{$gi}_tip_$id\">".mission_enemy_tip_render($role)."</span>";
				$html .= "<script>\$(function(){\$('#eg_{$row['id']}_{$gi}_td_{$id}').qtip({content: document.getElementById('eg_{$row['id']}_{$gi}_tip_{$id}').innerHTML, position:{ adjust:{screen:true} }, style: { name:'light', tip:'topLeft', border:{ width:3, radius:8 }, width:{max:1000} }});})</script>";
				$html .= "</td>";

				// $role_exp += $role['role_exp'];
				// $coins += $role['coins'];
			} else {
				$html .= "<td style=\"color:#ccc;cursor:default\">$id</td>";
			}
		}

		$html .= "</tr>";
	}

	$html .= "</table>";

	return $html;
}

function enemy_group_table_editor($allEnemyRole, $row) {
	$html = '';

	$allEnemyRole[0] = null;

	$html .= "<table style=\"width:100px; height:60px; margin:10px 0; background:#fff;\" colspan=\"0\" rowspan=\"0\" cellpadding=\"0\" cellspacing=\"0\">";

	foreach (array(10, 5, 0) as $i) {
		$html .= "<tr>";

		for ($j = 5; $j >= 1; $j --) {
			$id = ($i + $j) < 10 ? ($i + $j) : ''.($i + $j);

			$role = $allEnemyRole[$row['pos'.$id]];

			$html .= "<td>".enemy_role_editor($allEnemyRole, $row, 'pos'.$id)."</td>";
		}

		$html .= "</tr>";
	}

	$html .= "</table>";

	return $html;
}

// PVE的战斗信息
function enemy_group_info($sql){
	global $db;
	$html = '';
	$enemyGroupInfo = db_query($db, $sql);
	if (empty($enemyGroupInfo)) {
		return $html;
	}
	$html .= "</br>奖励命力：".$enemyGroupInfo[0]['award_soul'];
	$html .= "</br>奖励友情值：".$enemyGroupInfo[0]['award_friend'];
	$html .= "</br>奖励阅历：".$enemyGroupInfo[0]['award_skill_point'];
	$html .= "</br>是否开启几率属性：".($enemyGroupInfo[0]['battle_random'] == 1 ? "是" : "否");
	$html .= "</br>是否开启狂暴：".($enemyGroupInfo[0]['battle_crazy'] == 1 ? "是" : "否");
	if ($enemyGroupInfo[0]['battle_crazy'] == 1) {
		$html .= "</br>狂暴阶段1：第".$enemyGroupInfo[0]['crazy_round1']."回合时，伤害增加".$enemyGroupInfo[0]['crazy_buff_rate1']."%";
		$html .= "</br>狂暴阶段2：第".$enemyGroupInfo[0]['crazy_round2']."回合时，伤害增加".$enemyGroupInfo[0]['crazy_buff_rate2']."%";
		$html .= "</br>狂暴阶段3：第".$enemyGroupInfo[0]['crazy_round3']."回合时，伤害增加".$enemyGroupInfo[0]['crazy_buff_rate3']."%";
	}
	$html .= "</br>最大回合数(胜利条件)：". $enemyGroupInfo[0]['max_round'];
	return $html;
}

function skill_book_info($sql){
	global $db;
	global $item_additional_types;
	$html = '';
	$skillBookInfo = db_query($db, $sql);
	if (empty($skillBookInfo)) {
		return $html;
	}

	$quality = '';
	switch($skillBookInfo[0]['quality']){
		case "0":
			$quality = '蓝色';
			break;
		case "1":
			$quality = '紫色';
			break;
		case "2":
			$quality = '金色';
			break;
		case "3":
			$quality = '橙色';
			break;
	}
	$html .= "</br>品质：".$quality;
	$html .= "</br>绝招类型：".$skillBookInfo[0]['skill_type'];
	$html .= "</br>绝招名称：".$skillBookInfo[0]['name'];
	$html .= "</br>绝招可用：".$skillBookInfo[0]['skill_role_id'];
	$html .= "</br>每次领悟进度：".$skillBookInfo[0]['grasp_schedule']."%";
	$html .= "</br>声望等级：".$skillBookInfo[0]['fame_level'];
	$html .= "</br>修炼完加成：".$item_additional_types[$skillBookInfo[0]['add_data_type']]."+".$skillBookInfo[0]['add_data_value'];

	return $html;
}

function skill_buddy_info($sql){
	global $db;
	$html = '';
	$skillBookInfo = db_query($db, $sql);
	if (empty($skillBookInfo)) {
		return $html;
	}
	$num =$skillBookInfo[0]['release_num'];
	if ($num == '-1') {
		$num = '连续释放';
	}
	$html .= "</br>释放次数：".$num;
	return $html;
}

function enemy_auto_complete($allEnemyRole) {
	$html = "enemy_roles = [\n";

	$html .= "{value:'无', id:'0'},\n";

	foreach ($allEnemyRole as $key => $value) {
		$html .= "{value:'".$value['name']."', id:'".$key."'},\n";
	}

	$html .= "{value:'', id:'0'}\n";
	$html .= "];\n";

	$html .= "
$('.role_editor').autocomplete({
	lookup: enemy_roles, 
	minChars: 0,
	onSelect: function(s){ 
		console.log(s);
		$('#' + this.id + '_hidden')[0].setAttribute('value', s.id); 
		// addToChanges($(this).parent().parent());
	}
});
";

	return $html;
}

function get_items_json($items, $json_name) {
	$html = "$json_name = [\n";

	$html .= "{value:'无', id:'0'},\n";

	foreach ($items as $key => $value) {
		$html .= "{value:'".$value."', id:'".$key."'},\n";
	}

	$html .= "{value:'', id:'0'}\n";
	$html .= "];\n";

	return $html;
}

function choose_something($all_mission_level, $column_name) {
	$html = get_items_json($all_mission_level, $column_name);
	$html .= <<<EOT
	autoopt_{$column_name} = {
		lookup: $column_name, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});
EOT;
	return $html;
}

function choose_single_item($all_items, $column_name) {
	$html = get_items_json($all_items, 'items');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s) {
			$(this).siblings('input[class="real_value"]').val(s.id);
		}
	};

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;

	return $html;
}

function choose_multi_item($all_items, $column_name) {
	$html = get_items_json($all_items, 'items');

	$html .= <<<EOT

	autoopt_{$column_name} = {
		lookup: items, 
		minChars: 0,
		onSelect: function(s){ 
			item_obj = $(this).siblings('input[class="real_value"]');
			v = item_obj.val();
			if (v == ''){
				v = '[]';
			}
			items = JSON.parse(v);

			inputs = $(this).parent().find('.display_{$column_name}');
			index = null;

			for (var i = 0; i < inputs.length; i++) {
				if (inputs[i] == this){
					index = i ;
					break;
				}
			};

			id = parseInt(s.id);

			if (index == null) {
				index = inputs.length;
			}

			// 删除
			if (id == 0) {
				$(this).remove();
				items.splice(index, 1);
			} else {
				items[index] = id;
			}

			item_obj.val(JSON.stringify(items)) ;
		}
	};

	$(".add_{$column_name}").click(function(){
		$(this).before("<input class='display_{$column_name}' value=''> ");
		$(this).prev().autocomplete(autoopt_{$column_name});
	})

	$(".display_{$column_name}").autocomplete(autoopt_{$column_name});

EOT;
	
	return $html;
}

function editor_single_item($item_name, $item_id, $column_name) {
	$display = "<input class='display_{$column_name}' value='{$item_name}' /input>";
	$real_value = "<input type='hidden' class='real_value' name='{$column_name}[]' value='{$item_id}' />";
	return $display.$real_value;
}

function editor_multi_items($all_items, $row, $column_name) {
	$item_ids = json_decode($row[$column_name]);

	if (!is_array($item_ids)){
		$item_ids = array();
	}

	$display = '';
	foreach ($item_ids as $value) {
		$display .= "<input class='display_{$column_name}' value='{$all_items[$value]}' /input> ";
	}

	$btn = "<a class='add_{$column_name}' href='javascript:void(0)'>+</a>";
	$real_value = "<input type='hidden' class='real_value' name='{$column_name}[]' value='{$row[$column_name]}' />";

	return $display.$btn.$real_value;
}

function show_battle_info($all_enemy, $battle_type, $row) {
	$html  = '<table border="0" width="200"><tr><td width="180">';
	$html .= "<a href='?p=enemy_group_form&m={$row['id']}&m2=$battle_type' target='_blank'>怪物布阵</a> | <a href='?p=enemy_role&parent_id={$row['id']}&battle_type=$battle_type' target='_blank'>关联敌人</a>";
	$html .= enemy_group_table($all_enemy, $row, "select * from enemy_group_form where parent_id ={$row['id']} and `battle_type`=$battle_type order by `order` desc");
	$html .= '</td><td>';
	$html .= "<a href='?p=enemy_group&m={$row['id']}&m2=$battle_type' target='_blank'>战斗信息</a> ";
	$html .= enemy_group_info("select * from enemy_group where `parent_id`={$row['id']} and `battle_type`=$battle_type");
	$html .= '</td></tr></table>';
	return $html;
}

function get_all_enemy_group() {
	global $db;	
	$sql = "select l.name,e.id,e.order from mission_level l left join mission_enemy e on l.mission_id>0 && l.parent_type = 0 and e.mission_level_id = l.id order by l.lock,e.order";
	$rows = db_query($db, $sql);
	$allEnemyIds = array();
	$allEnemyIds[0] = '无';
	foreach ($rows as $key => $value) {
		if ($value['id'] > 0 && $value['order']>0) {
			$allEnemyIds[$value['id']] = $value['name'].'第'.$value['order'].'组';
		}
	}
	return $allEnemyIds;
}

function get_all_enemy_name() {
	global $db;
	$all_enemy = array();
	$sql = "select id,name from enemy_role";
	$query = db_query($db, $sql);
	foreach($query as $row) {
		$all_enemy[$row['id']] = $row['name'];
	}
	return $all_enemy;
}

function get_all_enemy() {
	global $db;	
	$all_enemy = array();

	$sql = "select *, (select name from skill where id = skill_id) skill_name from `enemy_role` order by `level` asc;";
	$tmp = db_query($db, $sql);

	foreach ($tmp as $value)
	{
		$all_enemy[$value['id']] = $value;
	}

	return $all_enemy;
}

//获取所有支线任务系列
function get_all_addition_quest_serial() {
	global $db;
	$sql = "select * from ( select serial_number as sn, name, award_lock, require_lock from addition_quest where serial_number > 0 order by require_lock ) q group by q.sn;";
	$query = db_query($db, $sql);
	$all_serial = array();
	$max_serial = 0;
	foreach($query as $key => $row) {
		$name = "";
		if(intval($row['require_lock']) == 0 && intval($row['award_lock']) == 0) {
			$name = ($row['name'] . '(孤立)');
		} else {
			$name = ($row['name'] . '(系列)');
		}

		$all_serial[$row['sn']] = $name;
		if(intval($row['sn']) > $max_serial) {
			$max_serial = intval($row['sn']);
		}
	}
	$_all_serial[$max_serial+1] = '新支线系列';
	//oh my dog!
	foreach($all_serial as $k => $v) {
		$_all_serial[$k] =  $v;
	}
	return $_all_serial;
}

//获取所有主线任务
function get_all_quest() {
	global $db;
	$all_quest = array();
	//key 900 以上的区域可能有特殊用途 比如剧情和怪物测试
	$sql = "select id,name from quest order by `order`";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_quest[$row['id']] = $row['name'];
	}
	return $all_quest;
}


function get_all_mission() {
	global $db;
	$all_mission = array();
	//key 900 以上的区域可能有特殊用途 比如剧情和怪物测试
	$sql = "select * from mission where  `keys` < 900"; 
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission[$row['id']] = $row['name'];
	}
	return $all_mission;
}

function get_all_town() {
	global $db;
	$all_town = array();
	$sql = "select id, name from town where  `lock` < 900000"; 
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_town[$row['id']] = $row['name'];
	}
	return $all_town;
}

function get_all_npc_role() {
	global $db;
	$all_roles = array();
	$sql = "select * from npc_role";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_roles[$row['id']] = $row['name'];
	}
	//var_dump($all_roles);
	return $all_roles;
}

function get_all_multi_level() {
	global $db;
	$all_mission_level = array();
	$sql = "select * from multi_level";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission_level[$row['id']] = $row['name'];
	}
	return $all_mission_level;
}

function get_all_rainbow_segment() {
	global $db;
	$all_mission_level = array();
	$sql = "select * from rainbow_level order by segment";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission_level[$row['id']] = "彩虹关卡第{$row['segment']}层";
	}
	return $all_mission_level;
}

function get_all_rainbow_level() {
	global $db;
	$all_mission_level = array();
	$sql = "select * from rainbow_level r left join mission_level m on  r.segment=m.parent_id and m.parent_type=12 order by r.segment, m.order;";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission_level[$row['id']] = "{$row['segment']}-{$row['order']}" . $row['name'];
	}
	return $all_mission_level;
}

/*
 * 0 主线
 * 1 资源
 * 8 深渊
 * 9 伙伴关卡
 * 10 灵宠关卡
 * 11 魂侍关卡
 */
function get_all_level_by_type($type) {
	global $db;
	$all_mission_level = array();
	$sql = "select id, name from mission_level where parent_type={$type}";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission_level[$row['id']] = $row['name'];
	}
	return $all_mission_level;
}

function get_all_mission_level_lock() {
	global $db;
	$all_mission_level = array();
	$sql = "select `lock`, name from mission_level where parent_type=0";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_mission_level[$row['lock']] = $row['name'];
	}
	return $all_mission_level;
}


function get_all_npc() {
	$towns = get_all_town();

	global $db;
	$all_npc = array();
	$sql = "select id, town_id, name from town_npc";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$town_name = "";
		if( isset($towns[$row['town_id']])) {
			$town_name = $towns[$row['town_id']];
		} else {
			continue;
		}

		$all_npc[$row['id']] = $town_name . "__" . $row['name'];
	}
	return $all_npc;
}

function get_all_role() {
	global $db;
	$all_roles = array();
	$sql = "select * from role";
	$rows = db_query($db, $sql);
	foreach($rows as $row) {
		$all_roles[$row['id']] = $row['name'];
	}
	return $all_roles;
}

//获取装备之外的普通物品
function get_item_exclude_equip() {
	global $db;
	$all_item = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id` not in (3,4,5,6,10,11)");
	foreach ($tmp as $value)
	{
		$all_item[$value['id']] = $value['name'];
	}
	return $all_item;
}

function get_all_item() {
	global $db;
	$all_item = array();

	$tmp = db_query($db, "select `id`, `name` from `item`");
	foreach ($tmp as $value)
	{
		$all_item[$value['id']] = $value['name'];
	}

	return $all_item;
}

function get_all_item_by_types($types, $resv = false) {
	global $db;
	$all_item = array();
	$sql = "";
	if(is_array($types)) {
		$ids = implode(',', $types);
		if($resv) {
			$sql = "select `id`, `name` from item where `type_id` NOT in ({$ids})";
		} else {
			$sql = "select `id`, `name` from item where `type_id` in ({$ids})";
		}
	} else {
		if($resv) {
			$sql = "select `id`, `name` from item where `type_id` != {$types}";
		} else {
			$sql = "select `id`, `name` from item where `type_id` = {$types}";
		}
	}
	$tmp = db_query($db, $sql);
	foreach ($tmp as $value)
	{
		$all_item[$value['id']] = $value['name'];
	}
	return $all_item;

}

function get_all_fashion_item() {
	global $db;
	$all_item = array();

	$tmp = db_query($db, "select `id`, `name` from `item` where `type_id` = 16");
	foreach ($tmp as $value)
	{
		$all_item[$value['id']] = $value['name'];
	}

	return $all_item;
}

function get_all_sword_soul() {
	global $db;
	$all_sword_soul = array();

	$tmp = db_query($db, "select `id`, `name` from `sword_soul`");
	foreach ($tmp as $value)
	{
		$all_sword_soul[$value['id']] = $value['name'];
	}

	return $all_sword_soul;
}

function get_all_ghost() {
	global $db;
	$all_ghosts = array();

	$tmp = db_query($db, "select `id`, `name` from `ghost`");
	foreach ($tmp as $value)
	{
		$all_ghosts[$value['id']] = $value['name'];
	}

	return $all_ghosts;
}

function get_all_pet() {
	global $db;
	$all_pets = array();

	$query = db_query($db, "select battle_pet.pet_id as id, enemy_role.name as name  from battle_pet inner join enemy_role on battle_pet.pet_id=enemy_role.id;");
	foreach ($query as $row) {
		$all_pets[$row['id']] = $row['name'];
	}
	return $all_pets;

}

function get_all_pet_ball() {
	global $db;
	$all_pets = array();

	$tmp = db_query($db, "select * from item where `type_id`='11'");

	foreach ($tmp as $value)
	{
		$all_pets[$value['id']] = $value['name'];
	}

	return $all_pets;
}

function normal_render($data, $row_value) {
	if (isset($data[$row_value])) {
		return $data[$row_value];
	}
}

function normal_editor($data, $row, $field_name) {
	$html = "<select class={$field_name} name={$field_name}[]>";
	$html .= '<option value="0" selected="selected">无</option>';

	foreach ($data as $key => $value) {
		if (isset($row[$field_name]) && $row[$field_name] == $key) {
			$html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
		} else {
			$html .= '<option value="'.$key.'">'.$value.'</option>';
		}
	}

	$html .= '</select>';
	return $html;
}

?>
