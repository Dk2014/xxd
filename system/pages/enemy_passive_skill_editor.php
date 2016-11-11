<?php 


require_once 'lib/config.php';
require_once 'lib/mysql.php';

$db = db_connect();

$allEnemySkill = array();
//TODO type = 10 怪物被动技能
$sql = "select `id`,`name` from skill where `type` = 5;";
$allEnemySkills = db_query($db, $sql);

foreach ($allEnemySkills as $value){
	$allEnemySkill[$value['id']] = $value['name'];
}

?>
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>jQuery UI Dialog - Default functionality</title>
  <link rel="stylesheet" href="../jquery-ui.css">
<?
  //<script src="../jquery.min.js"></script>
  //<script src="../jquery-ui.min.js"></script>
?>
  <style type="text/css">
	.editor {width: 70px;}

    .ui-dialog .ui-dialog-buttonpane { 
    text-align: center;
}
.ui-dialog .ui-dialog-buttonpane .ui-dialog-buttonset { 
    float: none;
}
  </style>
</head>
<script>



</script>

<body>
<div id="passive_skill_dialog" title="怪物被动技能配置">
  <table>
	<tr>
		<table id="enemy_passive_skill_config">

		</table>
	</tr>
  </table>
</div>
</body>
  </html>
<script>
var skills_list = [
<?php
$allEnemySkill[0] = '无'; 
foreach ($allEnemySkill as $key => $value) {
	echo "{value:'{$value}', id:'{$key}'},".PHP_EOL;
}
?>
];
//var select_c_html = '<?=html_get_select("conditions", $conditions, 1)?>';
var input_v_html = '<?=html_get_input("value", 0)?>';
var select_s_html = "<?php echo html_get_auto_completed_input('passive_skill##id', 'skills', $allEnemySkill, 0, 120); ?>";
//var select_s_html = '<?=html_get_select("skills", $allEnemySkill, 0)?>';
var input_p_html = '<?=html_get_input("power", 0)?>';
var ind = 1;

function delRow(node) {
	$(node).parent().parent().remove();
}

function enemy_passive_skill_newConfig() {
	$('#enemy_passive_skill_config').append('<tr><th>被动技能'+select_s_html.replace(/##id/g,ind)+' </th><th><a href="#" onclick="delRow(this)">&nbsp;&nbsp;删除</a></th></tr>');
	ind++;
	$('.skills').autocomplete({
		lookup: skills_list, 
			minChars: 0,
			onSelect: function(s){
				$('#' + this.id + '_hidden')[0].setAttribute('value', s.id);
			}
	});
}

function enemy_passive_skill_commit(enemy_id) {
	var config = [];

	$("input[name='skills[]']").each(function(){
		skill_id = parseInt($(this).val());
		if (skill_id > 0) {
			config.push(skill_id);
		}
	});


	$.post("pages/enemy_passive_skill_ajax.php",
		{
			id:enemy_id,
				config:JSON.stringify(config),
		},
		function(data, status){
			console.log(data, status);
			if (status != 'success') {
				alert("保存失败");
			}
		});
}

function open_passive_skill_editor(enemy_id) {
	$('#enemy_passive_skill_config tr').remove();
	ind = 1;
	$.get("pages/enemy_passive_skill_ajax.php?id=" + enemy_id,function(data, status){
		console.log(data);
		if (status != 'success') {
			alert("无法获取配置数据");
			return
		}

		var config = JSON.parse(data);
		if (config.length == 0){
			enemy_passive_skill_newConfig();
			return
		}

		var i=0;
		for (;i<config.length; i++){
			$('#enemy_passive_skill_config').append(enemy_passive_skill_newConfig());
		}

		i = 1;
		for(i=1;i<=config.length;i++){
			items = $.grep(skills_list, function(n){
				return n.id == parseInt(config[i-1]);
			});
			item = items[0];
			$("#passive_skill"+i).val(item.value);
			$("#passive_skill"+i+"_hidden").val(item.id);
		}


	});

	$( "#passive_skill_dialog" ).dialog(
		{width:550, height:400},
		{ buttons:[
		{ text: "添加", click: function() { enemy_passive_skill_newConfig(); }},
		{ text: "保存", click: function() { enemy_passive_skill_commit(enemy_id);$( this ).dialog( "close" ); }},
	]}
);
}
</script>
