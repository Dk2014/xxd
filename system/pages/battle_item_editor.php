 <?php 
 	require_once("battle_item_const.php");


  $item_mod_type = array();
  foreach ($battle_item_effect_type as $key => $value) {
    $item_mod_type[$key] = $value['name'];
  }

 	$item_buf_mod = array();
 	foreach ($battle_item_effect_buff_mod as $key => $value) {
 		$item_buf_mod[$key] = $value['name'];
 	}

  $item_target = array();
  foreach ($battle_item_target as $key => $value) {
    $item_target[$key] = $value['name'];
  }

  $item_normal = array();
  foreach ($battle_item_effect_normal_mod as $key => $value) {
    $name = (is_array($value)) ? $value['name']:$value;
    $item_normal[$key] = $name;
  }


 ?>
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>jQuery UI Dialog - Default functionality</title>
  <link rel="stylesheet" href="../jquery-ui.css">
  <script src="../jquery.min.js"></script>
  <script src="../jquery-ui.min.js"></script>
  <style type="text/css">
 	.editor {width: 80px;}
  	tr th {padding-left: 30px;}
  	tr td { padding-left: 30px;}
  	.config{border: solid #ccc 1px;width: 450px;height: 100px;}
    .ui-dialog .ui-dialog-buttonpane { 
    text-align: center;
}
.ui-dialog .ui-dialog-buttonpane .ui-dialog-buttonset { 
    float: none;
}
  </style>
</head>
<script>
 
 var current_mod;
 var select_buf_mod = '<?=html_get_select("mod", $item_buf_mod, 0)?>';
 var select_normal_mod = '<?=html_get_select("mod", $item_normal, 0)?>';
 
  function newConfig() {
  	var mod_value = '<?=html_get_input("value", 0)?>';
  	return '<tr><td>属性：'+current_mod+'</td><td>作用效果值：'+mod_value+' <a href="#" onclick="delModRow(this)"> 删除</a></td></tr>';
  }

  	function refreshSelBuff() {
		var attr = $('.buff_attr');
    	var seltext = $('select[name="effect_type[]"]').find("option:selected").text();
      var keep = $('input[name="keep[]"]');
    	var override=$('input[name="max_override[]"]');
    	seltext = seltext.toLowerCase();
    	if (seltext.indexOf("buff")>=0) {
    		attr.show();
    		keep.val(1);
    		override.val(1);
    		current_mod = select_buf_mod;
    	} else {
    		attr.hide();
    		keep.val(0);
    		override.val(0);
    		current_mod = select_normal_mod;
    	}
	}

  function commit(id) {
  		var config = [];
  		$("select[name='mod[]']").each(function(){
  			config.push({"mod":$(this).val(),"val":0});
  		})

  		var i=0;
  		$("input[name='value[]']").each(function(){
  			config[i++]['val'] = $(this).val();
  		})
  	
  	  $.post("pages/battle_item_ajax.php",
  		{
    		item_id:id,
    		target_type:$('select[name="target_type[]"]').val(),
    		effect_type:$('select[name="effect_type[]"]').val(),
        keep:$('input[name="keep[]"]').val(),
    		cost_power:$('input[name="cost_power[]"]').val(),
    		max_override:$('input[name="max_override[]"]').val(),
        can_use_level:$('select[name="can_use_level[]"]').val(),
    		item_config:JSON.stringify(config)
  		},
		  function(data, status){
		    if (status != 'success') {
		    	alert("更新失败");
		    }
		  });
  }

  function delModRow(node) {
  	$(node).parent().parent().remove();
  }

</script>

<body>
<div id="dialog" title="配置战斗道具">
  <table>
  	<tr>
  		<table>
  			<tr>
  				<th>消耗精气 <?=html_get_input('cost_power', 0)?></th>
          <th>目标对象 <?=html_get_select('target_type', $item_target, 1)?></th>
  				<th>产生效果 <?=html_get_select('effect_type', $item_mod_type, 1)?><a href="#" onclick="javascript:$('.config').append(newConfig());"> +</a></th>
          <th>关卡内可用 <?=html_get_select('can_use_level', array('0'=>'否','1'=>'是'), 0)?></th>
  			</tr>
  			<tr class="buff_attr">
	  			<th>持续回合 <?=html_get_input('keep', 1)?></th>
	  			<th>最大叠加数 <?=html_get_input('max_override', 1)?></th>
  			</tr>
  		</table>
  	</tr>

    <tr>
  		<td>
	    	<table class="config">

	    	</table>
  		</td>
  	</tr>
  	
  </table>
</div>
</body>
  </html>
  <script>
  	$('.buff_attr').hide();
	$("select[name='effect_type[]']").change(function(){
		refreshSelBuff();
    	$('.config tr').remove();
		$('.config').append(newConfig());
	});

	function open_editor(id) {
    $('.config tr').remove();
  	$.get("pages/battle_item_ajax.php?item_id=" + id,function(data, status){
		if (status != 'success') {
			alert("无法获取配置数据");
			return
		}
		data = JSON.parse(data);
		if (data.length == 0){
			$('.config').append(newConfig());
			return
		}

		$('select[name="can_use_level[]"]').val(data.can_use_level);
		$('select[name="target_type[]"]').val(data.target_type);
		$('select[name="effect_type[]"]').val(data.effect_type);
		refreshSelBuff();
    $('input[name="keep[]"]').val(data.keep);
		$('input[name="cost_power[]"]').val(data.cost_power);
		$('input[name="max_override[]"]').val(data.max_override);

		var config = JSON.parse(data.config);
		var i;
		for (i=0;i<config.length; i++){
			$('.config').append(newConfig());
		}

		i = 0;
		$("select[name='mod[]']").each(function(){
  			$(this).val(config[i++].mod);
  		});

  		i=0;
  		$("input[name='value[]']").each(function(){
  			$(this).val(config[i++].val);
  		});
  	
	});

    $( "#dialog" ).dialog(
  		{width:600, height:400},
  		{ buttons:[
  			 { text: "保存", click: function() { commit(id); $( this ).dialog( "close" ); }},
  		]}
	   );
  }
  refreshSelBuff();
  </script>