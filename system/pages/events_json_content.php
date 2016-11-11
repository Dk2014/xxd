<script src="jquery-ui.min.js"></script>
<script src="jquery.qtip-1.0.0-rc3.js"></script>
<script src="jquery.autocomplete.js"></script>
<link rel="stylesheet" href="jquery-ui.css" />
<style>
.events_table { border:solid 1px #ccc; border-right:0; border-bottom:0; font-size:13px; width:100%; }
.events_table td { border:solid 1px #ccc; border-left:0; border-top:0; padding:8px; }
.enemystable input{width:140px;}
</style>

<div id="events_josn_award_eidtor" title="奖品配置" style="display:none">
<table colspan="0" rowspan="0" cellpadding="0" cellspacing="0" id="events_info_variable_table"  class="events_table">
   <thead>
        <tr>
        <td >条件</td> 
        <td >元宝</td> 
        <td >铜钱</td>
        <td >爱心</td>
        <td >物品1</td>
        <td >数量</td>
        <td >物品2</td>
        <td >数量</td>
        <td >物品3</td>
        <td >数量</td>
        <td >物品4</td>
        <td >数量</td>
        <td >物品5</td>
        <td >数量</td>
        <td><div style="text-align: right;"><input type="button" value="+" onclick="$('#events_info_variable_table tbody').append(new_award_vars(null,'','','','',''))" /></div></td>
        </tr>    
    </thead>
    <tbody>
    
    </tbody>
</table>
</div>

<script>

<?php
    $all_items=get_all_item();
    $g_items=array();
?>
function events_josn_editor(Type, Page) {
    
    <?=get_items_json($all_items, 'items')?>
    g_items=items;

    $("#events_info_variable_table tbody").children().remove();
    $.get("pages/events_json_analyze.php?type="+Type+"&page="+Page, function(temp_vars, status){
        if (status!='success') {
            alert("无法获取配置数据");
            return;
        }
        
        var tvs=JSON.parse(temp_vars);
        for(var item in tvs){   
            var name5 = get_item_name(tvs[item]["Item1Id"]);
            var name7 = get_item_name(tvs[item]["Item2Id"]);
            var name9 = get_item_name(tvs[item]["Item3Id"]);
            var name11 = get_item_name(tvs[item]["Item4Id"]);
            var name13 = get_item_name(tvs[item]["Item5Id"]);
            var tr=new_award_vars(tvs[item],name5,name7,name9,name11,name13);
            $('#events_info_variable_table tbody').append(tr);
        }
  
    });
	

    $("#events_josn_award_eidtor").dialog({width: 1880, modal: true},
           {
            buttons: [
                {
                   text: "保存",
                    click: function() {
                        var events_info_vars={};
                      
                        $('#events_info_variable_table tbody').children("tr").each(function(ind, ele){
                            events_info_vars[ind]=get_events_info_variable(ele);
                        });
                       
                        $.post("pages/events_json_analyze.php",
                            {
                                'type': Type,
                                'page': Page,
                                'content': JSON.stringify(events_info_vars)
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
    });
		
	return false;
}

function get_item_id(name) {
    for ( var i=0 ; i < g_items.length; ++i ) {
        if (g_items[i].value == name) {
            return g_items[i].id;
        } 
    }
    return '0';
}

function get_events_info_variable(tr_tag){    
    var content={
        'Require': $(tr_tag).find("input[name=variable_name_1]").val(),
        'Ingot': $(tr_tag).find("input[name=variable_name_2]").val(),
        'Coin': $(tr_tag).find("input[name=variable_name_3]").val(),
        'Heart': $(tr_tag).find("input[name=variable_name_4]").val(),
        'Item1Id': get_item_id($(tr_tag).find("input[name=dummy_variable_name_5]").val()),
        'Item1Num': $(tr_tag).find("input[name=variable_name_6]").val(),
        'Item2Id': get_item_id($(tr_tag).find("input[name=dummy_variable_name_7]").val()),
        'Item2Num': $(tr_tag).find("input[name=variable_name_8]").val(),
        'Item3Id': get_item_id($(tr_tag).find("input[name=dummy_variable_name_9]").val()),
        'Item3Num': $(tr_tag).find("input[name=variable_name_10]").val(),
        'Item4Id': get_item_id($(tr_tag).find("input[name=dummy_variable_name_11]").val()),
        'Item4Num': $(tr_tag).find("input[name=variable_name_12]").val(),
        'Item5Id': get_item_id($(tr_tag).find("input[name=dummy_variable_name_13]").val()),
        'Item5Num': $(tr_tag).find("input[name=variable_name_14]").val(),
    };
    
    return content;                   
}

function get_item_name(id) {
    for ( var i=0 ; i < g_items.length; ++i ) {
        if (g_items[i].id == id) {
            return g_items[i].value;
        } 
    }
}

function new_award_vars(info_json,name5,name7,name9,name11,name13){
    
    var tr = $(events_award_vars_template
    .replace("variable_name_val_1", info_json==null?"0":info_json["Require"])
    .replace("variable_name_val_2", info_json==null?"0":info_json["Ingot"])
    .replace("variable_name_val_3", info_json==null?"0":info_json["Coin"])
    .replace("variable_name_val_4", info_json==null?"0":info_json["Heart"])
    .replace("variable_name_val_5", info_json==null?"0":info_json["Item1Id"])
    .replace("dummy_5", info_json==null?"0":name5)
    .replace("variable_name_val_6", info_json==null?"0":info_json["Item1Num"])
    .replace("variable_name_val_7", info_json==null?"0":info_json["Item2Id"])
    .replace("dummy_7", info_json==null?"0":name7)
    .replace("variable_name_val_8", info_json==null?"0":info_json["Item2Num"])
    .replace("variable_name_val_9", info_json==null?"0":info_json["Item3Id"])
    .replace("dummy_9", info_json==null?"0":name9)
    .replace("variable_name_val_10", info_json==null?"0":info_json["Item3Num"])
    .replace("variable_name_val_11", info_json==null?"0":info_json["Item4Id"])
    .replace("dummy_11", info_json==null?"0":name11)
    .replace("variable_name_val_12", info_json==null?"0":info_json["Item4Num"])
    .replace("variable_name_val_13", info_json==null?"0":info_json["Item5Id"])
    .replace("dummy_13", info_json==null?"0":name13)
    .replace("variable_name_val_14", info_json==null?"0":info_json["Item5Num"]));
    
    for ( var i=5 ; i < 14; i=i+2 ) {
        tr.find("input[name=dummy_variable_name_"+i+"]").autocomplete({
            lookup: g_items,
            minChars: 0,
            onSelect: function(s){
                $(this).siblings('.json_gameitem_editor[name=value]').val(s.id);
            }
        });
    }

    return tr;
}


var events_award_vars_template = '<?php echo str_replace("\n", "", str_replace("\t", "", <<<EOS
    <tr>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_1" name="variable_name_1" /></td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_2" name="variable_name_2" /></td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_3" name="variable_name_3" /></td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_4" name="variable_name_4" /></td>
        <td>  <input  style="width: 40px; display:none;" type="text" value="variable_name_val_5" name="variable_name_5" />
        <input type="text" name="dummy_variable_name_5" value="dummy_5" />
        </td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_6" name="variable_name_6" /></td>
        <td>  <input  style="width: 40px; display:none;" type="text" value="variable_name_val_7" name="variable_name_7" />
        <input type="text" name="dummy_variable_name_7" value="dummy_7" />
        </td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_8" name="variable_name_8" /></td>
        <td>  <input  style="width: 40px; display:none;" type="text" value="variable_name_val_9" name="variable_name_9" />
        <input type="text" name="dummy_variable_name_9" value="dummy_9" />
        </td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_10" name="variable_name_10" /></td>
        <td>  <input  style="width: 40px; display:none;" type="text" value="variable_name_val_11" name="variable_name_11" />
        <input type="text" name="dummy_variable_name_11" value="dummy_11" />
        </td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_12" name="variable_name_12" /></td>
        <td>  <input  style="width: 40px; display:none;" type="text" value="variable_name_val_13" name="variable_name_13" />
        <input type="text" name="dummy_variable_name_13" value="dummy_13" />
        </td>
        <td>  <input  style="width: 40px;" type="text" value="variable_name_val_14" name="variable_name_14" /></td>
        <td style="text-align: left;"><input type="button" value="-" onclick="$(this).parent().parent().remove()" /></td>
    </tr>
EOS
)); ?>';
     

</script>
