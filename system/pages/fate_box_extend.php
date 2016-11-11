<?php

//$extend_columns = array(
//);
$extend_columns=array(
  'fixed_prop' => array(
    'editor' => array('params'=>array()),
    'render' => array('params'=>array()),
    'data' => array(),
    'range' => array('params'=>array()),
  ),
);
$allItem = get_all_item();
function editor_fixed_prop($column_name, $column_val, $row, $data){
  global $allItem;
  $item_name='';
  if(isset($allItem[$column_val])){
    $item_name=$allItem[$column_val];
  }
 return editor_single_item($item_name, $column_val, $column_name);
}
function render_fixed_prop($column_name, $column_val, $row, $data){
  global $allItem;
  return (isset($allItem[$column_val]) ? $allItem[$column_val] : "0");
}
function range_fixed_prop(){
  global $allItem;
  $all_item_temp = $allItem;
  $all_item_temp[0] = '无';
  return $all_item_temp;
}
function operation($row) {
	return html_get_link("关卡", "?p=fate_box_level&m=".$row['id'], true) 
		. ' | ' .
		html_get_link("宝箱配置", "?p=fate_box_config&m=".$row['id'], true)
		. ' | ' .
		html_get_link("前20次配置", "?p=fate_box_presudo_config&m=-".$row['id'], true) // -id 作为前X次抽奖配置
		; 
}
function sql_where($params){
  return "order by 'id' ASC";
}
function jsFunction($params){
  global $allItem;

    $html = '
$("select").change( function(){
      $(this).css("color","#0033FF");
});';
$html .= choose_single_item($allItem, 'fixed_prop');
return $html;
}
?>
