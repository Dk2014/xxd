 <?php 

    require_once 'lib/config.php';
    require_once 'lib/mysql.php';

    $db = db_connect();

    $allEnemySkill = array();
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
  <script src="../jquery.min.js"></script>
  <script src="../jquery-ui.min.js"></script>
  <script src="../jquery.qtip-1.0.0-rc3.js"></script>
    <script src="../jquery.autocomplete.js"></script>
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
<div id="skillDialog" title="怪物技能配置">
  
  <table>
  	<tr>
  		<table id="skillConfig">
  		
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
  var e_input_rand_html = '<?=html_get_input("e_rand", 0)?>';
  var e_select_skill_html = "<?php echo html_get_auto_completed_input('skill##id', 'es_skills', $allEnemySkill, 0, 120); ?>";
  var e_input_power_html = '<?=html_get_input("e_power", 0)?>';
  var ind = 1;
  function newSkillConfig() {
    $('#skillConfig').append('<tr><th>'+e_input_rand_html+'%概率</th><th> 触发技能'+ e_select_skill_html.replace(/##id/g,ind) +' </th><th>威力'+e_input_power_html+'</th><th><a href="#" onclick="delRow(this)">&nbsp;&nbsp;删除</a></th></tr>');
    ind++;
    $('.es_skills').autocomplete({
        lookup: skills_list, 
        minChars: 0,
        onSelect: function(s){
          $('#' + this.id + '_hidden')[0].setAttribute('value', s.id);
        }
      });
  }

  function commitSkill(id) {
      var config = [];

      $("input[name='e_rand[]']").each(function() {
        config.push({"rand":$(this).val(), "skill_id":0, "power":0});
      });

      var i = 0;
      for(i=1;i<=config.length;i++){
        config[i-1]['skill_id'] = $("#skill"+i+"_hidden").val();
      }
      
      i = 0;
      $("input[name='e_power[]']").each(function(){
        config[i++]['power'] = $(this).val();
      });

      var postConfig = [];
      i = 0;
      for (; i < config.length; i++) {
          if (config[i].rand > 0 && config[i].skill_id > 0) {
             postConfig.push(config[i]);
          }
      }
      
      var postStr = "";
      if (postConfig.length > 0 ) {
        postStr = JSON.stringify(postConfig)
      }

      $.post("pages/enemy_skill_ajax.php",
      {
        id:id,
        config:postStr,
      },
      function(data, status){
        console.log(data, status);
        if (status != 'success') {
          alert("保存失败");
        }
      });
  }

  function open_skillEditor(id) {
    $('#skillConfig tr').remove();
    ind = 1;
    $.get("pages/enemy_skill_ajax.php?id=" + id,function(data, status){
      if (status != 'success') {
        alert("无法获取配置数据");
        return
      }

      var config = JSON.parse(data);
      if (config.length == 0){
        return
      }

      var i=0;
      for (;i<config.length; i++){
        $('#skillConfig').append(newSkillConfig());
      }

      i = 0;
      $("input[name='e_rand[]']").each(function(){
        $(this).val(config[i++]['rand']);
      });
      
      for(i=1;i<=config.length;i++){
        items = $.grep(skills_list, function(n){
          return n.id == parseInt(config[i-1]['skill_id']);
        });
        item = items[0];
        $("#skill"+i).val(item.value);
        $("#skill"+i+"_hidden").val(item.id);
      }

      i = 0;
      $("input[name='e_power[]']").each(function(){
        $(this).val(config[i++]['power']);
      });
      
    });


     $( "#skillDialog" ).dialog(
      {width:550, height:400},
      { buttons:[
         { text: "添加", click: function() { newSkillConfig(); }},
         { text: "保存", click: function() { commitSkill(id);$( this ).dialog( "close" ); }},
      ]}
     );
   }
  </script>