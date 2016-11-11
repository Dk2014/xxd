<?php

$extend_columns = array(
/*   '字段' => 配置 */
	'lock' => array(
		'editor' => array('params' => array()),
		'render' => array('params' => array()),
		'data' => array(),
		'regexp' => array("reg"=>"/\d{6}/i",
											"text" =>"应为6位数字",
											),
	),
);

function render_lock($column_name, $column_val, $row, $data){
	return format_big_number($column_val);
}

function editor_lock($column_name, $column_val, $row, $data){
	return html_get_input($column_name, format_big_number($column_val));
}

function item_blacksmith($row) {
	return html_get_link("区域", "?p=mission&m=".$row['id'], false). ' | '.
		   html_get_link("任务", "?p=quest&m=".$row['id'], false).' | '. 
		   html_get_link("NPC", "?p=town_npc&m=".$row['id'], false).' | '. 
		   html_get_link("商人", "?p=town_npc_item&m=".$row['id'], false) . ' | ' .
		   html_get_link("难度关卡", "?p=hard_level&m=".$row['id'], false) . ' | ' .
		   html_get_link("NPC对话", "?p=npc_talk&m=".$row['id'], false) . ' | ' .
	       html_get_link("剧情", "?p=game_scene&m=".$row['id'], false) . ' | ' .
	       html_get_link("下载奖励", "?p=town_treasure&m=".$row['id'], false). ' | ' .
	       html_get_link("区域评星奖", "?p=town_star_awards&m=".$row['id'], false);
}

function js_function($params){
	$html = <<<EOT

	$("#form").submit(function(){
		isOK = true;
		$("#table_body input[name='lock[]']").each(function(i){
			// 跳过第一行
			if ($(this).parent().parent().find("input[name='id[]']").val() < 0 ){
				return true;
			}
			newValue = $(this).val().replace(/,/g, '');
			if (newValue < 100000 || newValue > 999999) {
				isOk = false;
				return false;
			}
			$(this).val( $(this).val().replace(/,/g, '') )
		})
		if (!isOk && $("#sign_form").val()==0) {
			alert("权值位数应该设为6位")
			return false;
		} 
	})
EOT;
	return $html;
}


function sql_where($params) {
	return ' ORDER BY `lock`';
}

?>
