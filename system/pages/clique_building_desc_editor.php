<script src="jquery-ui.min.js"></script>
    <script src="jquery.qtip-1.0.0-rc3.js"></script>
    <script src="jquery.autocomplete.js"></script>
    <link rel="stylesheet" href="jquery-ui.css" />
<style>
.desc_head { border:solid 1px #ccc; border-right:0; border-bottom:0; font-size:13px; width:100%; text-align:right; height:20px;}
.desc_content { border:solid 1px #ccc; border-right:0; border-bottom:0; font-size:13px; width:100%; }
.desc_table td { border:solid 1px #ccc; border-left:0; border-top:0; padding:8px; }
.descs_list {border:solid 1px #ccc;padding-left: 50px;}
.desc_input_text { border:0; border-bottom:solid 1px #ccc; text-align:right; }
</style>

<script>
var index = 1;
var building_id = 0; // 预缓存当前修改描述的建筑物ID
// 加载获取数据
function descrEditor(id){
	$.get("/pages/clique_building_desc_ajax.php?id="+id, function(data){
		$('#descs_list').empty();
		data = JSON.parse(data);
		for (var i = 0; i < data.list.length; i++) {
			$('#descs_list').append("<div class='descs_list' id='descs_list_"+(i+1)+"'>等级:<input type='text' class='desc_input_text' id='descs_list_level_"+(i+1)+"' value='"+data.list[i].level+"' /> &nbsp;&nbsp;描述:<textarea id='descs_list_desc_"+(i+1)+"' rows='3' cols='60'>"+data.list[i].desc+"</textarea> <span class='descs_list_del' onclick='remove_item("+(i+1)+");'> 一 </span>");
			index++;
		}
	});
	building_id = id; 
	$("#clique_building_eidtor").dialog({width: 960, modal: true});
}

function save_editor(){
	dat = {};
	dat.list = [];
	ind = 0;
	$(".descs_list").each(function(i,data){
		level = $(this).children("input").first().val();
		if(parseInt(level) < 0){
			alert('第'+(i+1)+'行的建筑等级应为大于0的整数');
			return false;
		}
		desc = $(this).children("textarea").first().val();
		if(desc.length == 0){
			alert('第'+(i+1)+'行的建筑说明不能为空');
			return false;
		}
		dat.list[ind] = {};
		dat.list[ind].level = level;
		dat.list[ind++].desc = desc;
	});

	$.ajax({
		url: "/pages/clique_building_desc_ajax.php", 
		type: 'POST', 
		data: {
			'id': building_id,
			'descs': JSON.stringify(dat)
		}
	}).done(function(html){
		$('#clique_building_eidtor').dialog('close'); 
	});

	// 数据保存完毕关闭时 还原索引
	index = 1;
	building_id = 0;
}

$(document).ready(function(){

	$("#add").click(function(){
		$('#descs_list').append("<div class='descs_list' id='descs_list_"+index+"'>等级:<input type='text' class='desc_input_text' id='descs_list_level_"+index+"' value='' /> &nbsp;&nbsp;描述:<textarea id='descs_list_desc_"+index+"' rows='3' cols='60'></textarea> <span class='descs_list_del' onclick='remove_item("+index+");'> 一 </span>");
		index++;
	});
});

function remove_item(ind){
	$("#descs_list_"+ind).remove();
}

</script>
<div id="clique_building_eidtor" title="描述配置" style="display:none">
	<div class="desc_head"><span id="add">+</span></div>
<div class="desc_content" id="descs_list">

</div>
<div style="margin-top:8px;text-align:center">
	<input type="button" value="保存" onclick="save_editor()" />
</div>
</div>