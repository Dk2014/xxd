<!-- <script src="jquery-ui.min.js"></script>
<script src="jquery.qtip-1.0.0-rc3.js"></script>
<link rel="stylesheet" href="jquery-ui.css" /> -->
<script src="lang_rawstr.js"></script>
<style>
.skill_info_table { border:solid 1px #ccc; border-right:0; border-bottom:0; font-size:13px; width:100%; }
.skill_info_table td { border:solid 1px #ccc; border-left:0; border-top:0; padding:8px; }
</style>

<div id="skill_info_editor_dialog" title="绝招描述" style="display:none">
    <table colspan="0" rowspan="0" cellpadding="0" cellspacing="0" class="skill_info_table">
        <tr>
            <td colspan="2"><input id="skill_info_template_text" style="width: 100%;"></td>
        </tr>
        <tr style="background: #f0f0f0;">
            <td style="border-right: 0;">数值定义</td>
            <td><div style="text-align: right;"><input type="button" value="+" onclick="$('#skill_info_variable_table tbody').append(new_skill_info_variable('', null))" /></div></td>
        </tr>
        <tr><td colspan="2">
            <table id="skill_info_variable_table">
                <tbody></tbody>
            </table>
        </td></tr>
    </table>
</div>

<script>
function edit_skill_info(id){
    $("#skill_info_variable_table tbody").children().remove();
    $.get("pages/skill_info_ajax.php?id="+id, function(temp_vars, status){
        if (status!='success') {
            alert("无法获取配置数据");
            return;
        }
        var tvs=JSON.parse(temp_vars);
    	$("#skill_info_template_text").val(tvs.template);
        for(var var_name in tvs.variables){
            $('#skill_info_variable_table tbody').append(new_skill_info_variable(var_name, tvs.variables[var_name]));
        }
    });
    $("#skill_info_editor_dialog").dialog(
        {width: 550, height: 400},
        {
            buttons: [
                {
                    text: "保存",
                    click: function() {
                        var skill_info_vars={};
                        $('#skill_info_variable_table tbody').children("tr").each(function(ind, ele){
                            var_info=get_skill_info_variable(ele);
                            skill_info_vars[var_info.name]=var_info.infos;
                        });
						var content={
                        	'variables': skill_info_vars,
                        	'template': $("#skill_info_template_text").val()
						};
                        $.post("pages/skill_info_ajax.php",
                            {
                                'id': id,
								'content': JSON.stringify(content)
                            },
                            function(data, status){
                                console.log(data, status);
                                if (status != 'success') {
                                    alert("保存失败");
                                }
                            });
                        $(this).dialog("close");
                    }
                },
            ]
        }
    );
}

function new_skill_info_variable(var_name, info_json){
	if(info_json==null) {
		info_json = {};
	}
    if(info_json["RawBase"] == undefined) {
        info_json["RawBase"] = 0;
    }
    if(info_json["Cul2AtkRate"] == undefined) {
        info_json["Cul2AtkRate"] = 0;
    }
    if(info_json["TrnLvRate"] == undefined) {
        info_json["TrnLvRate"] = 0;
    }
    return skill_info_vars_template
    .replace("variable_name_val", var_name)
    .replace("RawBase_val", info_json["RawBase"])
    .replace("Cul2AtkRate_val", info_json["Cul2AtkRate"])
    .replace("TrnLvRate_val", info_json["TrnLvRate"]);
}

function get_skill_info_variable(tr_tag){
    var result={};
    result.name=$(tr_tag).find("input[name=variable_name]").val();
    result.infos={};
    result.infos["RawBase"]=$(tr_tag).find("input[name=RawBase]").val();
    result.infos["Cul2AtkRate"]=$(tr_tag).find("input[name=Cul2AtkRate]").val();
    result.infos["TrnLvRate"]=$(tr_tag).find("input[name=TrnLvRate]").val();
    return result;
}

var skill_info_vars_template = '<?php echo str_replace("\n", "", str_replace("\t", "", <<<EOS
    <tr>
        <td style="text-align: right;">
	    $<input style="width: 40px;" type="text" value="variable_name_val" name="variable_name" /> = 
		<input style="width: 40px" type="text" value="RawBase_val" name="RawBase" /> + 
		内力 * <input style="width: 40px" type="text" value="Cul2AtkRate_val" name="Cul2AtkRate" />% + 
		训练等级 * <input style="width: 40px" type="text" value="TrnLvRate_val" name="TrnLvRate" />
	</td>
        <td style="text-align: left;"><input type="button" value="-" onclick="$(this).parent().parent().remove()" /></td>
    </tr>
EOS
)); ?>';
</script>
