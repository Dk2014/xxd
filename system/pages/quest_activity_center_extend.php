<?php

$extend_columns = array(
    /*   '字段' => 配置 */
    'relative' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(
            1 => '等级活动',
            2 => '战力活动',
            3 => '累积登入活动',
            4 => '充值返利活动',
            5 => '比武场排名活动 ',
            6 => '翻倍奖励活动',
            7 => '活跃度活动',
            8 => '月卡活动',
	        9 => '午餐活动',
	        10 => '晚餐活动',
	        11 => 'QQ特权赠送物品',
	        12 => 'QQ特权奖励加成',
	        13 => 'VIP俱乐部',
	        14 => '土豪俱乐部',
	        15 => '冲级返利',
	        16 => '首充返利',
	        17 => '新手七天乐',
            18 => '分享送好礼',
            19 => '团购活动',
            20 => '累计充值元宝活动',
            21 => '累计消费元宝活动',
            22 => '元宝购买伙伴活动',
            23 => '神龙宝箱十连抽活动',
            24 => '每日首冲奖励',
            25 => '连续登陆奖励',
        ),
    ),
    'content' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
    ),
    'is_go' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(0 => '否', 1 => '是', 2 => '文字'),
    ),
    'is_mail' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(0 => '否', 1 => '是'),
    ),
    'tag' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(0 => '无', 1 => '最新', 2 => '限时', 3 => '推荐', 4=>'每日'),
    ),
    'is_relative' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(0 => '否', 1 => '是'),
    ),
    'mail_content' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
    ),
);

//活动配置索引
$award_config_arr = array(
		1 => 'events_level_up',
		2 => 'events_fight_power',
		3 => 'login_award',
		5 => 'events_arena_rank_awards',
		6 => 'event_multiply_config',
		7 => 'events_physical_awards',
		8 => 'events_month_card_awards',
		9 => 'events_dinner_awards',
		10 => 'events_supper_awards',
		11 => 'events_qqvip_gift_awards',
		12 => 'event_multiply_config',
		13 => 'events_vip_club_awards',
        14 => 'events_richman_club_awards',
		17 => 'events_seven_day_awards',
		18 => 'events_share_awards',
        19 => 'events_group_buy',
        20 => 'events_total_recharge',
        21 => 'events_total_consume',
        22 => 'events_buy_partner',
        23 => 'events_ten_draw_awards',
        24 => 'event_first_recharge_daily',
        25 => 'events_total_login',
	);

function event_config_index($row){
	global $award_config_arr;
	if(isset($award_config_arr[intval($row['relative'])])){
		return '<a href="./index.php?p='.$award_config_arr[intval($row['relative'])].'">配置</a>';
	}
	return '';
}

function event_status($row){
    if($row['is_relative'] == 1){
        return '<hr /><font color="green">开服活动</font>';
    }else if($row['start'] == 0){
        return '<hr /><font color="green">永久</font>';
    }else if($row['start'] > time()){
        return '<hr /><font color="red">未开启</font>';
    }else if($row['dispose'] > time()){
        return '<hr /><font color="green">进行中</font>';
    }else{
        return '<hr /><font color="red">已过期</font>';
    }
}

function editor_mail_content($column_name, $column_val, $row, $data)
{
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function render_mail_content($column_name, $column_val, $row, $data)
{
    return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_relative($column_name, $column_val, $row, $data)
{
    return html_get_select($column_name, $data, $column_val);
}

function render_relative($column_name, $column_val, $row, $data)
{
    return $data[$column_val];
}

function editor_content($column_name, $column_val, $row, $data)
{
    return html_get_textarea($column_name, $column_val, true, 4, 30);
}

function editor_is_relative($column_name, $column_val, $row, $data)
{
    return html_get_select($column_name, $data, $column_val);
}

function render_is_relative($column_name, $column_val, $row, $data)
{
    return $data[$column_val];
}

function render_content($column_name, $column_val, $row, $data)
{
    return html_get_textarea($column_name, $column_val, false, 4, 30);
}

function editor_is_go($column_name, $column_val, $row, $data)
{
    return html_get_select($column_name, $data, $column_val);
}

function render_is_go($column_name, $column_val, $row, $data)
{
    return $data[$column_val];
}

function editor_is_mail($column_name, $column_val, $row, $data)
{
    return html_get_select($column_name, $data, $column_val);
}

function render_is_mail($column_name, $column_val, $row, $data)
{
    return $data[$column_val];
}

function editor_tag($column_name, $column_val, $row, $data)
{
    return html_get_select($column_name, $data, $column_val);
}

function render_tag($column_name, $column_val, $row, $data)
{
    return $data[$column_val];
}
function sql_where($params) {
	$where = "";
	return " ORDER BY `weight`";
}

?>