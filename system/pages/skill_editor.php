<script src="jquery-ui.min.js"></script>
    <script src="jquery.qtip-1.0.0-rc3.js"></script>
    <script src="jquery.autocomplete.js"></script>
    <link rel="stylesheet" href="jquery-ui.css" />
<style>
.skill_table { border:solid 1px #ccc; border-right:0; border-bottom:0; font-size:13px; width:100%; }
.skill_table td { border:solid 1px #ccc; border-left:0; border-top:0; padding:8px; }
.skill_input_text { border:0; border-bottom:solid 1px #ccc; text-align:right; }
.enemystable input{width:140px;}
</style>
<script>
var editor_for_ghost = false;

var input_hightlight_on = false;

function open_editor(title, info, id, for_ghost) {
	editor_for_ghost = for_ghost;
	if (!input_hightlight_on) {
		setInterval(function(){
			inputs = $('.skill_input_text');
			for (var i = 0; i < inputs.length; i ++) {
				if (inputs[i].value != '0') {
					$(inputs[i]).css('color', '#f37700');
				} else {
					$(inputs[i]).css('color', '');
				}
			}
		}, 500);
	}

	$.ajax({
		url: "ajax.php?type=load_skill_config&id=" + id, 
		dataType: "json"
	}).done(function(data){ 
		console.log(data);
		if (data == undefined)
			data = {};

		if (data.TargetMode == undefined) {
			if (editor_for_ghost) {
				data.TargetMode = 4; // 魂侍默认不攻击
			} else {
				data.TargetMode = 0;
			}
		}

		if (data.AttackMode == undefined)
			data.AttackMode = 0;

		if (data.KillSelfRate == undefined)
			data.KillSelfRate = 1;

		if (data.DefaultAttack == undefined)
			data.DefaultAttack = 0;

		if (data.LevelAttack == undefined)
			data.LevelAttack = 0;
		
		if (data.Cul2AtkRate == undefined)
			data.Cul2AtkRate = 0;

		if (data.TrnlvRate == undefined)
			data.TrnlvRate = 0;

		if (data.DecPower == undefined)
			data.DecPower = 0;

		if (data.IncPower == undefined)
			data.IncPower = 0;

		if (data.HurtAdd == undefined)
			data.HurtAdd = 0;

		if (data.HurtAddRate == undefined)
			data.HurtAddRate = 0;

		if (data.CureAdd == undefined)
			data.CureAdd = 0;

		if (data.CureAddRate == undefined)
			data.CureAddRate = 0;

		if (data.Critial == undefined)
			data.Critial = 0;

		if (data.ReduceDefend == undefined)
			data.ReduceDefend = 0;

		if (data.SunderAttack == undefined)
			data.SunderAttack = 0;

		if (data.LevelSunderAttack == undefined)
			data.LevelSunderAttack = 0;

		if (data.MustHit == undefined)
			data.MustHit = false;

		if (data.SelfBuffs == undefined)
			data.SelfBuffs = [];

		if (data.TargetBuffs == undefined)
			data.TargetBuffs = [];

		if (data.BuddyBuffs == undefined)
			data.BuddyBuffs = [];

		if (data.GhostOverrideTargetBuff == undefined)
			data.GhostOverrideTargetBuff = false;

		if (data.GhostOverrideSelfBuff == undefined)
			data.GhostOverrideSelfBuff = false;

		if (data.GhostOverrideBuddyBuff == undefined)
			data.GhostOverrideBuddyBuff = false;

		if (data.AttackNum == undefined)
			data.AttackNum = 1;

		if (data.AppendSpecialType == undefined) {
			data.AppendSpecialType = 0;
			data.AppendSpecialValue = 0;
		}

		if (data.TriggerTargetBuff == undefined)
			data.TriggerTargetBuff = false;

		if (data.EventTriggerTarget == undefined)
			data.TriggerTargetBuff = 0;

		if (data.EventTriggerTarget == undefined)
			data.EventTriggerTarget = 0;

		if (data.EventTriggerBuff == undefined)
			data.EventTriggerBuff = 6;
		if(data.EnemyCalls == undefined)
			data.EnemyCalls = [];
		$(".buff_table").remove();
		$(".enemyss").remove();
		$(".enemystable .role_editor").val('无');
		$(".enemystable input[type='hidden']").val(0);
		enemys_table_id = 0;
		if (editor_for_ghost) {
			$('#prop_box').css('display', 'none');
		} else {
			$('#ghost_props').css('display', 'none');
		}

		document.getElementById('title').innerHTML           = title + "，" + info;
		document.getElementById('targets').selectedIndex     = data.TargetMode;
		document.getElementById('attack').selectedIndex      = data.AttackMode;
		document.getElementById('kill_self_rate').value      = data.KillSelfRate;
		document.getElementById('skill_default_atk').value   = data.DefaultAttack;
		document.getElementById('skill_cul2atk_rate').value  = data.Cul2AtkRate;
		document.getElementById('skill_trnlv_rate').value    = data.TrnlvRate;
 
		document.getElementById('dec_power').value           = data.DecPower;
		document.getElementById('inc_power').value           = data.IncPower;
		document.getElementById('hurt_add').value            = data.HurtAdd;
		document.getElementById('hurt_add_rate').value       = data.HurtAddRate;
		document.getElementById('cure_add').value            = data.CureAdd;
		document.getElementById('cure_add_rate').value       = data.CureAddRate;
		document.getElementById('critial').value             = data.Critial;
		document.getElementById('reduce_defend').value       = data.ReduceDefend;
		document.getElementById('sunder_attack').value       = data.SunderAttack;
		document.getElementById('must_hit').checked          = data.MustHit;

		document.getElementById('ghost_override_target_buff').checked = data.GhostOverrideTargetBuff;
		document.getElementById('ghost_override_self_buff').checked   = data.GhostOverrideSelfBuff;
		document.getElementById('ghost_override_buddy_buff').checked  = data.GhostOverrideBuddyBuff;

		document.getElementById('attack_num').value           = data.AttackNum;

		document.getElementById("config_skill_id").value = id; 

		document.getElementById('append_skill_special_type').selectedIndex = data.AppendSpecialType;
		document.getElementById('append_skill_special_value').value = data.AppendSpecialValue;

		document.getElementById('event_trigger_type').selectedIndex = data.EventTriggerType;
		document.getElementById('event_trigger_target').value = data.EventTriggerTarget;
		document.getElementById('event_trigger_buff').value = data.EventTriggerBuff;
		document.getElementById('trigger_target_buff').checked = data.TriggerTargetBuff;

		display_buff_config('self_buffs', data.SelfBuffs, buff_template1);
		display_buff_config('target_buffs', data.TargetBuffs, buff_template1);
		display_buff_config('buddy_buffs', data.BuddyBuffs, buff_template2);
		display_enemys_config('enemy_calls', data.EnemyCalls,call_template);

		attack_target_change(document.getElementById('targets'));
		attack_mode_change(document.getElementById('attack'));

		$("#skill_eidtor").dialog({width: 960, modal: true});

		if (!editor_for_ghost) {
			$(".ghost_only").css("display", "none");
		}
	}).fail(function(jqXHR, textStatus){
		console.log(textStatus);
	});

	return false;
}

function save_editor() {
	$.ajax({
		url: 'ajax.php?type=save_skill_config', 
		type: 'POST', 
		data: {
			'id': document.getElementById('config_skill_id').value,
			'config': JSON.stringify(get_json())
		}
	}).done(function(html){
		$('#skill_eidtor').dialog('close'); 
	});
}
var enemy_roles = [ /* 怪物召唤*/
{value:'无', id:'0'},
<?php 
$sql = 'SELECT `id`,`name` FROM `enemy_role`';
global $db;
$enemys = db_query($db,$sql);
$optionstr = '';
$allEnemyRole = array();
foreach ($enemys as $index => $enemy) {
	echo "{value:'".$enemy['name']."', data:'".$enemy['id']."'}";
	if($index < count($enemys)-1){
		echo ',';
	}
	$optionstr .= ('<option value="'.$enemy['id'].'">'.$enemy['name'].'</option>');
}

?>
];




function display_enemys_config(id, configs, template){
	var td = document.getElementById(id);
	var random_group = new Array();
	for(var i=0;i < configs.length; i++){
		if(configs[i].Position>=0){
			//固定位置
			var data_temp = $.grep(enemy_roles,function(s){
				return s.data == configs[i].Enemys;
			});
			if(data_temp.length == 0){continue;}
			var data = data_temp[0];
			var position = configs[i].Position + 1;
			$("#role_editor_pos"+position).val(data.value);
			$("#role_editor_pos"+position+"_hidden").val(data.data);
		}else{
			//随机位置
			var key = configs[i].Enemys + "_" + configs[i].Position;
			var item = random_group[key];
			if(typeof(item) != 'undefined'){
				random_group[key]++;
			}else{
				random_group[key] = 1;
			}
		}
	}

	for(var p in random_group){
		var item_now = random_group[p];
		var enemyid_and_position = p.split('_');
		var enemyid = enemyid_and_position[0];
		var position = enemyid_and_position[1];
		
		var data_temp = $.grep(enemy_roles,function(s){
				return s.data == enemyid;
			});
		if(data_temp.length==0) continue;
		var data = data_temp[0];
		new_call_table(td, template);
		//enemys_table_id
		$("#position"+(enemys_table_id-1)).val(position);
		$("#count"+(enemys_table_id-1)).val(item_now);
		$("#range_row_"+(enemys_table_id-1)+"_hidden").val(enemyid);
		$("#range_row_"+(enemys_table_id-1)).val(data.value);
	}


	// $(td).append(call_templatetable)
	// for(var i=0;i < configs.length; i++){
	// 	new_call_table(td, template);
	// }
	if(configs.length>0){
		document.getElementById('is_call_on').checked=true;
		call_on();
	}else{
		document.getElementById('is_call_on').checked=false;
		call_on();
		$('.enemystable input').val('无');
		$('.enemystable input[type="hidden"]').val('');
	}
}

function display_buff_config(id, configs, template) {
	var td = document.getElementById(id);

	for (var i = 0; i < configs.length; i ++) {
		new_buff_table(td, template);
	}

	var buff_type       = $('select[name=buff_type]', td);
	var buff_keep       = $('select[name=buff_keep]', td);
	var buff_override   = $('select[name=buff_override]', td);
	var buff_rate       = $('input[name=buff_rate]', td);
	var buff_count_rate = $('input[name=buff_count_rate]', td);
	var buff_after_atk  = $('input[name=buff_after_atk]', td);
	var buff_target     = $('select[name=buff_target]', td);
	var buff_skill_level = $('input[name=buff_skill_level]', td);
	var buff_uncleanable = $('input[name=buff_uncleanable]', td);

	var buff_sign             = $('select[name=buff_sign]', td);
	var buff_bace_value       = $('input[name=buff_base_value]', td);
	var buff_raw_value_rate   = $('input[name=buff_raw_value_rate]', td);
	var buff_attack_rate      = $('input[name=buff_attack_rate]', td);
	var buff_skill_force_rate = $('input[name=buff_skill_force_rate]', td);
	var buff_hurt_rate        = $('input[name=buff_hurt_rate]', td);
	var buff_cul2value_rate   = $('input[name=buff_cul2value_rate]', td);
	var buff_trnlv_rate       = $('input[name=buff_trnlv_rate]', td);
	var buff_value_count_rate = $('input[name=buff_value_count_rate]', td);

	for (var i = 0; i < configs.length; i ++) {
		if (configs[i].LevelValue == undefined)
			configs[i].LevelValue = 0;

		if (configs[i].Cul2ValueRate == undefined)
			configs[i].Cul2ValueRate = 0;
		
		if (configs[i].TrnlvRate == undefined)
			configs[i].TrnlvRate = 0;

		if (configs[i].LevelRate == undefined)
			configs[i].LevelRate = 0;
		
		if (configs[i].CountRate == undefined)
			configs[i].CountRate = 0;
		
		if (configs[i].ValueCountRate == undefined)
			configs[i].ValueCountRate = 0;

		if (configs[i].EventTriggerType == undefined)
			configs[i].EventTriggerType = 0;
		if (configs[i].EventTriggerTarget == undefined)
			configs[i].EventTriggerTarget = 0;
		if (configs[i].EventTriggerBuff == undefined)
			configs[i].EventTriggerBuff = 0;

		buff_type[i].selectedIndex     = configs[i].Type;
		buff_keep[i].selectedIndex     = configs[i].Keep;
		buff_override[i].selectedIndex = configs[i].Override;
		buff_rate[i].value             = configs[i].Rate;
		buff_count_rate[i].value       = configs[i].CountRate;

		buff_sign[i].selectedIndex     = configs[i].BuffSign;
		buff_bace_value[i].value       = configs[i].BaseValue;
		buff_raw_value_rate[i].value   = configs[i].RawValueRate;
		buff_attack_rate[i].value      = configs[i].AttackRate;
		buff_skill_force_rate[i].value = configs[i].SkillForceRate;
		buff_hurt_rate[i].value        = configs[i].HurtRate;
		buff_cul2value_rate[i].value   = configs[i].Cul2ValueRate;
		buff_trnlv_rate[i].value       = configs[i].TrnlvRate;
		buff_value_count_rate[i].value = configs[i].ValueCountRate;
		buff_skill_level[i].value = configs[i].RequireSkillLevel ? parseInt(configs[i].RequireSkillLevel) : 0;
		console.log(buff_uncleanable[i], configs[i].BuffUncleanable);
		buff_uncleanable[i].checked = Boolean(configs[i].BuffUncleanable);
		//buff_uncleanable[i] = configs[i].BuffUncleanable ? Boolean(configs[i].BuffUncleanable) : false;

		if (configs[i].TargetMode) {
			buff_target[i].selectedIndex = configs[i].TargetMode;
		}

		buff_type_change(buff_type[i]);
	}
}

function buff_type_change(target) {
	var buff_type = target.selectedIndex;

	if (buff_type != 0 &&  //精气
		buff_type != 4 &&  ///生命
		buff_type != 7 &&  //清除负面状态
		buff_type != 8 &&  //清除增益状态
		buff_type != 17 && //最大生命
		buff_type != 22 && //护甲
		buff_type != 27 && //魂力
		buff_type != 28 && //灵宠回合数
		buff_type != 29 && //伙伴进阶技能使用次数
		buff_type != 30 && //清除护盾伤害
		buff_type != 45 && //生命百分比
		buff_type != 47 && //复活并加生命值
		buff_type != 48 && //复活并加生命百分比
		true
	) {
		$("select[name=buff_keep], select[name=buff_override]", target.parentNode.parentNode.parentNode).removeAttr('disabled');
	} else {
		$("select[name=buff_keep], select[name=buff_override]", target.parentNode.parentNode.parentNode).attr('disabled', 'disabled');
	}

	if (buff_type != 5 && buff_type != 7 && buff_type != 8) {
		$("input[name=buff_value]", target.parentNode.parentNode.parentNode).attr('value', '0').removeAttr('readonly');
	} else {
		$("input[name=buff_value]", target.parentNode.parentNode.parentNode).attr('value', '1').attr('readonly', 'readonly');
	}
}

var buff_table_id = 0;

function new_buff_table(target, template) {
	$(target).append(template.replace(/#id#/gm, buff_table_id));
	buff_table_id += 1;

	attack_target_change(document.getElementById('targets'));
	attack_mode_change(document.getElementById('attack'));
}

var enemys_table_id = 0;
function new_call_table(target, template){
	// if(enemys_table_id==0){
	// 	$(target).append("<p>随机怪物组：</p>")
	// }
	$(target).append(template.replace(/#id#/gm, enemys_table_id));
	enemys_table_id++;
}

function get_json() {
	return {
		"TargetMode": parseInt(document.getElementById('targets').selectedIndex),
		"AttackMode": parseInt(document.getElementById('attack').selectedIndex),
		"KillSelfRate": parseInt(document.getElementById('kill_self_rate').value),
		"DefaultAttack": parseInt(document.getElementById('skill_default_atk').value),
		"Cul2AtkRate": parseInt(document.getElementById('skill_cul2atk_rate').value),
		"TrnlvRate": parseInt(document.getElementById('skill_trnlv_rate').value),
		"DecPower": parseInt(document.getElementById('dec_power').value),
		"IncPower": parseInt(document.getElementById('inc_power').value),
		"HurtAdd": parseInt(document.getElementById('hurt_add').value),
		"HurtAddRate": parseInt(document.getElementById('hurt_add_rate').value),
		"CureAdd": parseInt(document.getElementById('cure_add').value),
		"CureAddRate": parseInt(document.getElementById('cure_add_rate').value),
		"Critial": parseInt(document.getElementById('critial').value),
		"ReduceDefend": parseInt(document.getElementById('reduce_defend').value),
		"SunderAttack": parseInt(document.getElementById('sunder_attack').value),
		"MustHit": document.getElementById('must_hit').checked,
		"GhostOverrideBuddyBuff": document.getElementById('ghost_override_buddy_buff').checked,
		"GhostOverrideSelfBuff": document.getElementById('ghost_override_self_buff').checked,
		"GhostOverrideTargetBuff": document.getElementById('ghost_override_target_buff').checked,
		"AttackNum": parseInt(document.getElementById('attack_num').value),
		"SelfBuffs": get_buff_json('self_buffs'),
		"TargetBuffs": get_buff_json('target_buffs'),
		"BuddyBuffs": get_buff_json('buddy_buffs'),
		"EnemyCalls": get_call_json('enemy_calls'),
		"AppendSpecialType": parseInt(document.getElementById('append_skill_special_type').selectedIndex),
		"AppendSpecialValue": parseInt(document.getElementById('append_skill_special_value').value),
		"EventTriggerType": parseInt(document.getElementById('event_trigger_type').value),
		"EventTriggerTarget": parseInt(document.getElementById('event_trigger_target').value),
		"EventTriggerBuff": parseInt(document.getElementById('event_trigger_buff').value),
		"TriggerTargetBuff": document.getElementById('trigger_target_buff').checked
	};
}

function get_call_json(id){
	var td = document.getElementById(id);
	var json = [];
 	$('.enemystable .role_editor').each(function(){
 		if(parseInt($("#" + $(this).attr('id')+"_hidden").val())>0){
			anEnemy = {
				"Enemys": parseInt($("#" + $(this).attr('id')+"_hidden").val()),
				"Position": parseInt($(this).attr('position'))-1
			};
			json.push(anEnemy);
 		}
 	});

 	$('.enemyss').each(function(){
 		var role_input = $(this).children('.role_editor').first();
 		
 		var index = $(this).attr('index');
 		if($(role_input).val()!='无'){
 			var count = $("#count"+index).val();
 			count = parseInt(count);
 			for(var i=0;i<count;i++){
				anEnemy ={
					"Enemys": parseInt($("#" + $(role_input).attr('id')+ "_hidden").val()),
					"Position": parseInt($("#position"+index).val())
				};
				json.push(anEnemy);
 			}
 		}

 	});

	return json;
}

function get_buff_json(id) {
	var td = document.getElementById(id);

	var json = [];

	var buff_type = $('select[name=buff_type]', td);
	var buff_keep = $('select[name=buff_keep]', td);
	var buff_override = $('select[name=buff_override]', td);
	var buff_rate = $('input[name=buff_rate]', td);
	var buff_count_rate = $('input[name=buff_count_rate]', td);
	var buff_target = $('select[name=buff_target]', td);

	var buff_sign = $('select[name=buff_sign]', td);
	var buff_bace_value = $('input[name=buff_base_value]', td);
	var buff_raw_value_rate = $('input[name=buff_raw_value_rate]', td);
	var buff_attack_rate = $('input[name=buff_attack_rate]', td);
	var buff_skill_force_rate = $('input[name=buff_skill_force_rate]', td);
	var buff_hurt_rate = $('input[name=buff_hurt_rate]', td);
	var buff_cul2value_rate = $('input[name=buff_cul2value_rate]', td);
	var buff_trnlv_rate = $('input[name=buff_trnlv_rate]', td);
	var buff_value_count_rate = $('input[name=buff_value_count_rate]', td);
	var buff_skill_level = $('input[name=buff_skill_level]', td);
	var buff_uncleanable = $('input[name=buff_uncleanable]', td);


	for (var i = 0; i < buff_type.length; i ++) {
		var oneBuff = {
			"Type": parseInt(buff_type[i].selectedIndex),
			"Keep": parseInt(buff_keep[i].selectedIndex),
			"Override": parseInt(buff_override[i].selectedIndex),
			"Rate": parseInt(buff_rate[i].value),
			"CountRate": parseInt(buff_count_rate[i].value),
			"BuffSign": parseInt(buff_sign[i].selectedIndex),
			"BaseValue": parseInt(buff_bace_value[i].value),
			"RawValueRate": parseInt(buff_raw_value_rate[i].value),
			"AttackRate": parseInt(buff_attack_rate[i].value),
			"SkillForceRate": parseInt(buff_skill_force_rate[i].value),
			"HurtRate": parseInt(buff_hurt_rate[i].value),
			"Cul2ValueRate": parseInt(buff_cul2value_rate[i].value),
			"TrnlvRate": parseFloat(buff_trnlv_rate[i].value),
			"ValueCountRate": parseInt(buff_value_count_rate[i].value),
			"RequireSkillLevel": parseInt(buff_skill_level[i].value),
			"BuffUncleanable": Boolean(buff_uncleanable[i].checked)
		};
		console.log(oneBuff);

		if (buff_target.length > 0) {
			oneBuff["TargetMode"] = parseInt(buff_target[i].selectedIndex);
		}

		json.push(oneBuff);
	}

	return json;
}

function attack_target_change(target) {
	if (target.selectedIndex == 4 || target.selectedIndex == 8) {
		$('#attack_fun_box').css('display', 'none');
	} else {
		$('#attack_fun_box').css('display', '');
	}

	if (target.selectedIndex == 8) {
		$('#target_buff_tr1').css('display', 'none');
		$('#target_buff_tr2').css('display', 'none');
		$('#buddy_buff_tr1').css('display', 'none');
		$('#buddy_buff_tr2').css('display', 'none');
		$('.skill_props').css('display', 'none');
		$('.buff_sign_box').css('display', 'none');
		$('.buff_atk_box').css('display', 'none');

		$('select[name=buff_sign]').attr('selectedIndex', 0);
	} else {
		$('#target_buff_tr1').css('display', '');
		$('#target_buff_tr2').css('display', '');
		$('#buddy_buff_tr1').css('display', '');
		$('#buddy_buff_tr2').css('display', '');
		$('.skill_props').css('display', '');
		$('.buff_sign_box').css('display', '');
		$('.buff_atk_box').css('display', '');
	}
}

function attack_mode_change(target) {
	switch (target.selectedIndex) {
	case 1: //自爆
		$('.polynomial_zero_value').css('display', '');
		$('.polynomial_base_attack').css('display', 'none');
		$('.polynomial_kill_self_rate').css('display', '');
		$('.polynomial_skill_force').css('display', '');
		$('.polynomial_cul2atk_rate').css('display', 'none');
		$('.polynomial_default_atk').css('display', 'none');
		$('.polynomial_trnlv_rate').css('display', 'none');
		$('#prop_box').css('display', 'none');
		break;
	case 2: //玩家
		$('.polynomial_zero_value').css('display', 'none');
		$('.polynomial_base_attack').css('display', '');
		$('.polynomial_kill_self_rate').css('display', 'none');
		$('.polynomial_skill_force').css('display', 'none');
		$('.polynomial_cul2atk_rate').css('display', '');
		$('.polynomial_default_atk').css('display', '');
		$('.polynomial_trnlv_rate').css('display', '');
		if (editor_for_ghost) {
			$('#prop_box').css('display', 'none');
		} else {
			$('#prop_box').css('display', '');
		}
		break;
	case 3: //魂侍
		$('.polynomial_zero_value').css('display', 'none');
		$('.polynomial_base_attack').css('display', '');
		$('.polynomial_kill_self_rate').css('display', 'none');
		$('.polynomial_skill_force').css('display', 'none');
		$('.polynomial_cul2atk_rate').css('display', '');
		$('.polynomial_default_atk').css('display', '');
		$('.polynomial_trnlv_rate').css('display', '');
		$('#prop_box').css('display', 'none');
		break;
	case 4: //灵宠
		$('.polynomial_zero_value').css('display', 'none');
		$('.polynomial_base_attack').css('display', '');
		$('.polynomial_kill_self_rate').css('display', 'none');
		$('.polynomial_skill_force').css('display', 'none');
		$('.polynomial_cul2atk_rate').css('display', 'none');
		$('.polynomial_default_atk').css('display', '');
		$('.polynomial_trnlv_rate').css('display', '');
		$('#prop_box').css('display', 'none');
		break;
	case 5: //阵印
		$('.polynomial_zero_value').css('display', 'none');
		$('.polynomial_base_attack').css('display', 'none');
		$('.polynomial_kill_self_rate').css('display', 'none');
		$('.polynomial_skill_force').css('display', 'none');
		$('.polynomial_cul2atk_rate').css('display', 'none');
		$('.polynomial_default_atk').css('display', 'none');
		$('.polynomial_trnlv_rate').css('display', 'none');
		$('#prop_box').css('display', 'none');
		break;
	default: //普通攻击
		$('.polynomial_zero_value').css('display', 'none');
		$('.polynomial_base_attack').css('display', '');
		$('.polynomial_kill_self_rate').css('display', 'none');
		$('.polynomial_skill_force').css('display', '');
		$('.polynomial_cul2atk_rate').css('display', 'none');
		$('.polynomial_default_atk').css('display', 'none');
		$('.polynomial_trnlv_rate').css('display', 'none');
		$('#prop_box').css('display', 'none');
		break;
	}
}

$(function(){
	$('#targets').change(function(){ attack_target_change(this); });
	$('#attack').change(function(){ attack_mode_change(this); });
});

<?php /*自身buff或目标buff*/ ?>
var buff_template1 = '<?php echo str_replace("\n", "", str_replace("\t", "", <<<EOS
	<table colspan="0" rowspan="0" cellpadding="0" cellspacing="0" style="margin-top:8px" id="buff_#id#" class="skill_table buff_table">
	<tr>
		<td>
			<select name="buff_type" onchange="javascript:buff_type_change(this)">
			<option value="0">精气</option>
			<option value="1">速度(禁用)</option>
			<option value="2">攻击</option>
			<option value="3">防御</option>
			<option value="4">生命</option>
			<option value="5">眩晕</option>
			<option value="6">中毒</option>
			<option value="7">清除负面</option>
			<option value="8">清除增益</option>
			<option value="9">免伤百分比</option>
			<option value="10">混乱</option>
			<option value="11">格挡</option>
			<option value="12">格挡等级</option>
			<option value="13">闪避等级</option>
			<option value="14">暴击等级</option>
			<option value="15">命中等级</option>
			<option value="16">伤害百分比</option>
			<option value="17">最大生命</option>
			<option value="18">守卫者免伤</option>
			<option value="19">吸引火力</option>
			<option value="20">破击等级</option>
			<option value="21">韧性等级</option>
			<option value="22">护甲</option>
			<option value="23">睡眠</option>
			<option value="24">禁用绝招</option>
			<option value="25">直接免伤</option>
			<option value="26">吸收伤害</option>
			<option value="27">魂力</option>
			<option value="28">灵宠存活回合</option>
			<option value="29">伙伴技能次数</option>
			<option value="30">清除伤害吸收</option>
			<option value="31">睡眠抗性等级</option>
			<option value="32">眩晕抗性等级</option>
			<option value="33">混乱抗性等级</option>
			<option value="34">封魔抗性等级</option>
			<option value="35">中毒抗性等级</option>
			<option value="36">伙伴技能次数恢复</option>
			<option value="37">满精气</option>
			<option value="38">闪避</option>
			<option value="39">命中</option>
			<option value="40">暴击</option>
			<option value="41">韧性</option>
			<option value="42">破甲值</option>
			<option value="43">百分比防御</option>
			<option value="44">破甲状态(直接破甲)</option>
			<option value="45">生命百分比</option>
			<option value="46">所有抗性等级</option>
			<option value="47">复活并加生命值</option>
			<option value="48">复活并加生命百分比</option>


			</select>
			<span class="buff_sign_box">
			<select name="buff_sign">
			<option>+</option>
			<option>-</option>
			</select>
			</span>
			(
				基础值 <input type="text" style="width:48px" class="skill_input_text" name="buff_base_value" value="0" /> + 
				原属性值 × <input type="text" style="width:32px" class="skill_input_text" name="buff_raw_value_rate" value="0" /> % + 
				<span class="buff_atk_box">
				攻击属性 x <input type="text" style="width:32px" class="skill_input_text" name="buff_attack_rate" value="0" /> % + 
				绝招威力 × <input type="text" style="width:32px" class="skill_input_text" name="buff_skill_force_rate" value="0" /> % + 
				本次伤害 × <input type="text" style="width:32px" class="skill_input_text" name="buff_hurt_rate" value="0" /> % + 
				</span> 
				内力 x <input type="text" style="width:48px" class="skill_input_text" name="buff_cul2value_rate" value="0" /> % ＋ 
			    训练等级 x <input type="text" style="width:48px" class="skill_input_text" name="buff_trnlv_rate" value="0" /> - 
				生效次数 × <input type="text" style="width:32px" class="skill_input_text" name="buff_value_count_rate" value="0" />
			)
		</td>
		<td align="center" rowspan="2">
			<input type="button" value="-" onclick="javascript:$(\'#buff_#id#\').remove()" />
		</td>
	</tr>
	<tr>
		<td>
			发生概率 <input name="buff_rate" type="text" style="width:32px" value="100" class="skill_input_text" /> %  - 生效次数 x <input name="buff_count_rate" type="text" style="width:32px" value="0" class="skill_input_text" /> %，
			持续 <select name="buff_keep" disabled="disabled">
			<option>∞</option>
			<option>1</option>
			<option>2</option>
			<option>3</option>
			<option>4</option>
			<option>5</option>
			<option>6</option>
			<option>7</option>
			<option>8</option>
			<option>9</option>
			</select> 回合，
			可叠加 <select name="buff_override" disabled="disabled">
			<option>∞</option>
			<option selected="selected">1</option>
			<option>2</option>
			<option>3</option>
			<option>4</option>
			<option>5</option>
			<option>6</option>
			<option>7</option>
			<option>8</option>
			<option>9</option>
			</select> 次
			要求等级 <input name="buff_skill_level" class="skill_input_text" type="text" style="width:32px" value="0">
			不可清除 <input name="buff_uncleanable" type="checkbox" />
		</td>
	</tr>
	</table>
EOS
)); ?>';

<?php /*队伍buff*/?>
var buff_template2 = '<?php echo str_replace("\n", "", str_replace("\t", "", <<<EOS
	<table colspan="0" rowspan="0" cellpadding="0" cellspacing="0" style="margin-top:8px" id="buff_#id#" class="skill_table buff_table">
	<tr>
		<td>
			<select name="buff_type" onchange="javascript:buff_type_change(this)">
			<option value="0">精气</option>
			<option value="1">速度(禁用)</option>
			<option value="2">攻击</option>
			<option value="3">防御</option>
			<option value="4">生命</option>
			<option value="5">眩晕</option>
			<option value="6">中毒</option>
			<option value="7">清除负面</option>
			<option value="8">清除增益</option>
			<option value="9">免伤百分比</option>
			<option value="10">混乱</option>
			<option value="11">格挡</option>
			<option value="12">格挡等级</option>
			<option value="13">闪避等级</option>
			<option value="14">暴击等级</option>
			<option value="15">命中等级</option>
			<option value="16">伤害百分比</option>
			<option value="17">最大生命</option>
			<option value="18">守卫者免伤</option>
			<option value="19">吸引火力</option>
			<option value="20">破击等级</option>
			<option value="21">韧性等级</option>
			<option value="22">护甲</option>
			<option value="23">意志</option>
			<option value="24">禁用绝招</option>
			<option value="25">直接免伤</option>
			<option value="26">吸收伤害</option>
			<option value="27">魂力</option>
			<option value="28">灵宠存活回合</option>
			<option value="29">伙伴技能次数</option>
			<option value="30">清除伤害吸收</option>
			<option value="31">睡眠抗性等级</option>
			<option value="32">眩晕抗性等级</option>
			<option value="33">混乱抗性等级</option>
			<option value="34">封魔抗性等级</option>
			<option value="35">中毒抗性等级</option>
			<option value="36">伙伴技能次数恢复</option>
			<option value="37">满精气</option>
			<option value="38">闪避</option>
			<option value="39">命中</option>
			<option value="40">暴击</option>
			<option value="41">韧性</option>
			<option value="42">破甲</option>
			<option value="43">百分比防御</option>
			<option value="44">破甲状态(直接破甲)</option>
			<option value="45">生命百分比</option>
			<option value="46">所有抗性等级</option>
			<option value="47">复活并加生命值</option>
			<option value="48">复活并加生命百分比</option>

			</select>
			<select name="buff_sign">
			<option>+</option>
			<option>-</option>
			</select>
			(
				基础值 <input type="text" style="width:48px" class="skill_input_text" name="buff_base_value" value="0" /> + 
				原属性值 × <input type="text" style="width:32px" class="skill_input_text" name="buff_raw_value_rate" value="0" /> % + 
				攻击属性 x <input type="text" style="width:32px" class="skill_input_text" name="buff_attack_rate" value="0" /> % + 
				绝招威力 × <input type="text" style="width:32px" class="skill_input_text" name="buff_skill_force_rate" value="0" /> % + 
				本次伤害 × <input type="text" style="width:32px" class="skill_input_text" name="buff_hurt_rate" value="0" /> % +  
				内力 x <input type="text" style="width:48px" class="skill_input_text" name="buff_cul2value_rate" value="0" /> % + 
				训练等级 x <input type="text" style="width:48px" class="skill_input_text" name="buff_trnlv_rate" value="0" /> - 
			    生效次数 × <input type="text" style="width:32px" class="skill_input_text" name="buff_value_count_rate" value="0" />
			)
		</td>
		<td align="center" rowspan="2">
			<input type="button" value="-" onclick="javascript:$(\'#buff_#id#\').remove()" />
		</td>
	</tr>
	<tr>
		<td>
			发生概率 <input name="buff_rate" type="text" style="width:32px" value="100" class="skill_input_text"> % - 生效次数 x <input name="buff_count_rate" type="text" style="width:32px" value="0" class="skill_input_text" /> %，
			作用到 <select name="buff_target">
			<option>我方全体</option>
			<option>自己以外的一个伙伴</option>
			<option>血最少的伙伴</option>
			<option>我方灵宠</option>
			<option>全体伙伴</option>
			<option>前排横向固定</option>
			<option>中排横向固定</option>
			<option>后排横向固定</option>
			<option>我方一个死亡角色</option>
			</select>，
			持续 <select name="buff_keep" disabled="disabled">
			<option>∞</option>
			<option>1</option>
			<option>2</option>
			<option>3</option>
			<option>4</option>
			<option>5</option>
			<option>6</option>
			<option>7</option>
			<option>8</option>
			<option>9</option>
			</select> 回合，
			可叠加 <select name="buff_override" disabled="disabled">
			<option>∞</option>
			<option selected="selected">1</option>
			<option>2</option>
			<option>3</option>
			<option>4</option>
			<option>5</option>
			<option>6</option>
			<option>7</option>
			<option>8</option>
			<option>9</option>
			</select> 次
			要求等级 <input name="buff_skill_level" class="skill_input_text" type="text" style="width:32px" value="0">
			不可清除 <input name="buff_uncleanable" type="checkbox" />
		</td>
	</tr>
	</table>
EOS
)); ?>';

function call_on(){
	if(document.getElementById('is_call_on').checked==true){
		document.getElementById('enemy_calls').style.display='';
		$('.role_editor').autocomplete({
			lookup: enemy_roles, 
			minChars: 0,
			onSelect: function(s){ 
				$('#' + this.id + '_hidden')[0].setAttribute('value', s.data);
			}
});
		document.getElementById("add_range").style.display = '';
	}else{
		document.getElementById('enemy_calls').style.display='none';
		document.getElementById("add_range").style.display = 'none';
	}
}

function add_range(){
	if(document.getElementById('is_call_on').checked==true){
		new_call_table(document.getElementById('enemy_calls'), call_template);
		document.getElementById('enemy_calls').style.display='';
		$('.role_editor').autocomplete({
			lookup: enemy_roles, 
			minChars: 0,
			onSelect: function(s){ 
				$('#' + this.id + '_hidden')[0].setAttribute('value', s.data);
			}
});
	}
}




var call_templatetable = '';
var call_template = '<?php echo "<div style=\\'margin-left:25px;margin-bottom:15px\\' index=\\'#id#\\' id=\\'enemys_#id#\\' class=\\'enemyss\\'>#id#. 怪物：<input id=\\'range_row_#id#\\' class=\\'editor role_editor\\' type=\\'text\\' value=\\'无\\' style=\\'width: 140px;\\' autocomplete=\\'off\\'><input id=\\'range_row_#id#_hidden\\' type=\\'hidden\\' value=\\'0\\' name=\\'range_pos#id#\\'>" ?>';
call_template += "<?php echo str_replace("\n","",str_replace("\r", "", <<<EOS
&nbsp;&nbsp;位置：<select id='position#id#'>
	<option value='-5'>完全随机</option>
	<option value='-4'>左侧随机</option>
	<option value='-3'>右侧随机</option>
	<option value='-2'>前排随机</option>
	<option value='-1'>后排随机</option>
</select> 数量：<input id=\\'count#id#\\' /><input type='button' value='-' onclick='javascript:$(\"#enemys_#id#\").remove()' />
EOS
))?>";
</script>
<div id="skill_eidtor" title="绝招配置" style="display:none">
<table colspan="0" rowspan="0" cellpadding="0" cellspacing="0" class="skill_table">
	<tr>
		<td width="72">绝招名称</td><td id="title"></td>
	<tr>
	<tr>
		<td width="72">攻击方式</td>
		<td>
			<input type="hidden" id="config_skill_id" />
			<select id="targets">
			<option value="0">单体</option>
			<option value="1">全体</option>
			<option value="2">横向</option>
			<option value="3">纵向</option>
			<option value="4">不攻击</option>
			<option value="5">后排横向</option>
			<option value="6">血量最少</option>
			<option value="7">单体重复两次</option>
			<option value="8">被动技</option>
			<option value="9">纵向攻击，并额外随机攻击另一列</option>
			<option value="10">随机攻击两个目标</option>
			<option value="11">单体攻击，并额外随机攻击另一人</option>
			<option value="12">血量最多</option>
			<option value="13">V字攻击</option>
			<option value="14">后排单体攻击</option>
			<option value="15">十字攻击</option>
			<option value="16">横排穿透</option>
			<option value="17">十字攻击(固定)</option>
			<option value="18">前排横向(固定)</option>
			<option value="19">中排横向(固定)</option>
			<option value="20">后排横向(固定)</option>
			</select>
			<span id="attack_fun_box">，
				<select id="attack">
				<option value="0">普通</option>
				<option value="1">自爆</option>
				<option value="2">玩家</option>
				<option value="3">魂侍</option>
				<option value="4">灵宠</option>
				<option value="5">阵印</option>
				</select>
				=
				<span class="polynomial_zero_value" style="display: none;">0</span>
				<span class="polynomial_base_attack">攻击力</span>
				<span class="polynomial_kill_self_rate"> + 剩余血量 × <input id="kill_self_rate" type="text" style="width:32px" value="0" class="skill_input_text" /> %</span>
				<span class="polynomial_skill_force"> + 绝招威力</span>
				<span class="polynomial_cul2atk_rate"> + 内力 × <input type="text" style="width:32px" class="skill_input_text" id="skill_cul2atk_rate" value="0" /> %</span>
				<span class="polynomial_default_atk"> + <input type="text" style="width:60px" class="skill_input_text" id="skill_default_atk" value="0" /></span>
				<span class="polynomial_trnlv_rate"> + 训练等级 × <input type="text" style="width:32px" class="skill_input_text" id="skill_trnlv_rate" value="0" /></span>
				<!-- 
				<span id="attack_fun0">攻击力 + 绝招威力</span>
				<span id="attack_fun1" style="display:none">剩余血量 × <input id="kill_self_rate" type="text" style="width:32px" value="0" class="skill_input_text" /> % + 绝招威力</span>
				<span id="attack_fun2" style="display:none">攻击力 + 内力 × <input type="text" style="width:32px" class="skill_input_text" id="skill_cul2atk_rate" value="0" /> % + <input type="text" style="width:60px" class="skill_input_text" id="skill_default_atk" value="0" /> + 训练等级 × <input type="text" style="width:32px" class="skill_input_text" id="skill_trnlv_rate" value="0" /></span>
				<span id="attack_fun3" style="display:none">攻击力 + 内力 × <input type="text" style="width:32px" class="skill_input_text" id="skill_cul2atk_rate" value="0" /> % + <input type="text" style="width:60px" class="skill_input_text" id="skill_default_atk" value="0" /> + 训练等级 × <input type="text" style="width:32px" class="skill_input_text" id="skill_trnlv_rate" value="0" /></span>
				<span id="attack_fun4" style="display:none">攻击力 + <input type="text" style="width:60px" class="skill_input_text" id="skill_default_atk" value="0" /> + 训练等级 × <input type="text" style="width:32px" class="skill_input_text" id="skill_trnlv_rate" value="0" /></span>
				 -->
			</span>
		</td>
	</tr>
	<tr id="ghost_props">
		<td width="72">魂侍特性</td>
		<td>
			攻击次数 <input type="text" style="width:32px" id="attack_num" class="skill_input_text" />，
			伤害增加(直接值) <input type="text" style="width:32px" id="hurt_add" class="skill_input_text" />，
			伤害增加(百分比) <input type="text" style="width:32px" id="hurt_add_rate" class="skill_input_text" />，
			治疗增加(直接值) <input type="text" style="width:32px" id="cure_add" class="skill_input_text" />，
			治疗增加(百分比) <input type="text" style="width:32px" id="cure_add_rate" class="skill_input_text" />
		</td>
	</tr>
	<tr class="skill_props">
		<td width="72">绝招特性</td>
		<td>
			<span id="prop_box">
			消耗精气 <input type="text" style="width:32px" id="dec_power" class="skill_input_text" />，
			恢复精气 <input type="text" style="width:32px" id="inc_power" class="skill_input_text" />，
			</span>
			暴击概率 <input type="text" style="width:32px" id="critial" class="skill_input_text" /> %，
			无视对方防御 <input type="text" style="width:32px" id="reduce_defend" class="skill_input_text" /> %，
			破甲值 <input type="text" style="width:32px" id="sunder_attack" class="skill_input_text" /> + 

			百分百命中 <input type="checkbox" id="must_hit" />
		</td>
	</tr>
	<tr class="skill_props">
		<td width="72">绝招附加特性</td>
		<td>
			类型 <select id="append_skill_special_type">
			<option value="0">无</option>
			<option value="1">每命中1个敌人增加精气</option>
			</select>，
			属性值 <input type="text" style="width:50px" id="append_skill_special_value" value="0" class="skill_input_text" />，
		</td>
	</tr>
	<tr class="event_trigger">
		<td width="72">事件触发</td>
		<td>
			事件 <select id="event_trigger_type">
			<option  selected = "selected"  value="0">无</option>
			<option value="1">闪避</option>
			<option value="2">暴击</option>
			<option value="3">格挡</option>
			<option value="4">反击（未支持）</option>
			</select>，

			条件 <select id="event_trigger_target">
			<option  selected = "selected"  value="0">攻击方</option>
			<option value="1">防守方</option>
			</select> 有
			Buff <select id="event_trigger_buff">
			<option  selected = "selected"  value="6">中毒</option>
			<option value="26">吸收伤害</option>
			</select>，
			关联目标BUFF <input type="checkbox" id="trigger_target_buff" />

		<td>
	</tr>	
	<tr style="background: #f0f0f0;" id="target_buff_tr1">
		<td style="border-right:0;">目标BUFF</td>
		<td>
			<div style="text-align:right;">
				<label class="ghost_only"><input type="checkbox" id="ghost_override_target_buff" />覆盖原buff&nbsp;&nbsp;</label>
				<input type="button" value="+" onclick="javascript:new_buff_table(document.getElementById('target_buffs'), buff_template1)" />
			</div>
		</td>
	</tr>
	<tr id="target_buff_tr2">
		<td id="target_buffs" colspan="2"></td>
	</tr>
	<tr style="background: #f0f0f0;" id="self_buff_tr1">
		<td style="border-right:0;">自身BUFF</td>
		<td>
			<div style="text-align:right;">
				<label class="ghost_only"><input type="checkbox" id="ghost_override_self_buff" />覆盖原buff&nbsp;&nbsp;</label>
				<input type="button" value="+" onclick="javascript:new_buff_table(document.getElementById('self_buffs'), buff_template1)" />
			</div>
		</td>
	</tr>
	<tr id="self_buff_tr1">
		<td id="self_buffs" colspan="2"></td>
	</tr>
	<tr style="background: #f0f0f0;" id="buddy_buff_tr1">
		<td style="border-right:0;">队伍BUFF</td>
		<td>
			<div style="text-align:right;">
				<label class="ghost_only"><input type="checkbox" id="ghost_override_buddy_buff" />覆盖原buff&nbsp;&nbsp;</label>
				<input type="button" value="+" onclick="javascript:new_buff_table(document.getElementById('buddy_buffs'), buff_template2)" />
			</div>
		</td>
	</tr>
	<tr id="buddy_buff_tr2">
		<td id="buddy_buffs" colspan="2"></td>
	</tr>

	<tr style="background: #f0f0f0;<?php if($_GET['p']!='skill_enemy') echo "display:none" ?>" id="enemy_call_tr1" >
		<td style="border-right:0;">怪物召唤</td>
		<td>
			<div style="text-align:right;">
				<input id="is_call_on" type="checkbox" onclick="javascript:call_on()"/>是否需要召唤？&nbsp;&nbsp;<input id="add_range" type="button" style="display:none" value="增加随机怪物" onclick="javascript:add_range()" />
			</div>
		</td>
	</tr>
	<tr id="enemy_call_tr2">
		<td id="enemy_calls" colspan="2" style="display:none"><div style="margin-left:25px;margin-bottom:15px;width:880px"><table class="enemystable"><caption align="top">固定怪物组</caption> <tr><td><input id="role_editor_pos15" position="15" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos15_hidden" type="hidden" value="0" name="pos15"></td><td><input id="role_editor_pos14" position="14" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos14_hidden" type="hidden" value="0" name="pos14"></td><td><input id="role_editor_pos13" position="13" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos13_hidden" type="hidden" value="0" name="pos13"></td>
		<td><input id="role_editor_pos12" position="12" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos12_hidden" type="hidden" value="0" name="pos12"></td>
		<td><input id="role_editor_pos11" position="11" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos11_hidden" type="hidden" value="0" name="pos11"></td></tr>
		<tr><td><input id="role_editor_pos10" position="10" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos10_hidden" type="hidden" value="0" name="pos10"></td>
			<td><input id="role_editor_pos9" position="9" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos9_hidden" type="hidden" value="0" name="pos9"></td>
			<td><input id="role_editor_pos8" position="8" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos8_hidden" type="hidden" value="0" name="pos8"></td>
			<td><input id="role_editor_pos7" position="7" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos7_hidden" type="hidden" value="0" name="pos7"></td>
			<td><input id="role_editor_pos6" position="6" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos6_hidden" type="hidden" value="0" name="pos6"></td></tr>
			<tr><td><input id="role_editor_pos5" position="5" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos5_hidden" type="hidden" value="0" name="pos5"></td>
				<td><input id="role_editor_pos4" position="4" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos4_hidden" type="hidden" value="0" name="pos4"></td>
				<td><input id="role_editor_pos3" position="3" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos3_hidden" type="hidden" value="0" name="pos3"></td>
				<td><input id="role_editor_pos2" position="2" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos2_hidden" type="hidden" value="0" name="pos2"></td>
				<td><input id="role_editor_pos1" position="1" class="editor role_editor" type="text" value="无" style="width: 140px;" autocomplete="off"><input id="role_editor_pos1_hidden" type="hidden" value="0" name="pos1"></td></tr></table></div> <hr /><p>随机怪物组</p></td>
	</tr>
</table>
<div style="margin-top:8px;text-align:center">
	<input type="button" value="保存" onclick="save_editor()" />
</div>
</div>
