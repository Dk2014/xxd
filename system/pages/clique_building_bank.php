<?php

include "clique_buildings_link.php";

$pconfig = array(
	'title'   => '帮派建筑物 -- 钱庄',
	'table'   => 'clique_building_bank',
	'links'   => $clique_buildings_link,
	
	'columns' => array(
		array('name' => 'level', 'text' => '钱庄等级', 'width' => '80px'),
		array('name' => 'silver_coupon_coins', 'text' => '银劵底价', 'width' => '80px'),
		array('name' => 'silver_coupon_num', 'text' => '银劵限购', 'width' => '80px'),
		array('name' => 'silver_coupon_sold', 'text' => '银劵售价', 'width' => '80px'),
		array('name' => 'gold_coupon_coins', 'text' => '金劵底价', 'width' => '80px'),
		array('name' => 'gold_coupon_num', 'text' => '金劵限购', 'width' => '80px'),
		array('name' => 'gold_coupon_sold', 'text' => '金劵售价', 'width' => '80px'),
		array('name' => 'upgrade', 'text' => '钱庄升级所需铜钱', 'width' => '200px'),
		array('name' => 'desc', 'text' => '对应等级描述', 'width' => '200px'),
		array('name' => 'cur_level_desc', 'text' => '当前钱庄描述', 'width' => '200px'),
		array('name' => 'next_level_desc', 'text' => '下一等级钱庄描述', 'width' => '200px'),
	),
	
	'opera' => array(
	),
	

	'where' 		=> 'sql_where',
);

function sql_where(){
	return ' ORDER BY `level` ASC';
}

?>
