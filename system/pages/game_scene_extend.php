<?php

$extend_columns = array(
        'quest_order' => array(
                'editor' 	=> array('params' => array()),
                'render' 	=> array('params' => array()),
                'data' 		=> array(),
                'range'     => array('params' => array()),
        ),
);

$match_quest_record_sql = "select `order`, `name` from `quest`";

$town_id = null; //I don't know how to do this, just bear it.
if (isset($_POST['m']))
    $town_id = $_POST['m'];
else if (isset($_GET['m']))
    $town_id = $_GET['m'];

if (isset($town_id))
    $match_quest_record_sql .= " where `related_town`={$town_id}";

$matched_quest_records = db_query($db, $match_quest_record_sql);

$allQuestOrder = array();
foreach($matched_quest_records as $record)
    $allQuestOrder[$record['order']] = $record['name'];

function range_quest_order(){
    global $allQuestOrder;
    return $allQuestOrder;
}

function render_quest_order($column_name, $column_val, $row, $data){
    global $allQuestOrder;
    return $allQuestOrder[$column_val];
    //return html_get_link($allMissionLevel[$column_val], '?p=mission_enemy&m='.$column_val, true);
}

function editor_quest_order($column_name, $column_val, $row, $data){
    global $allQuestOrder;
    $quest_order = $row['quest_order'];
    $quest_order_name = '';
    if ( isset($allQuestOrder[$quest_order])) {
        $quest_order_name = $allQuestOrder[$quest_order];
    }
    return editor_single_item($quest_order_name, $quest_order, "quest_order");
    //return html_get_select($column_name, $allMissionLevel, $column_val);
}


function sql_where($params) {
    global $matched_quest_records;
    
    $where = "";

    $order_sql_set = array_reduce($matched_quest_records, function($snow_ball, $item){
        $snow_ball.=$item['order'].',';
        return $snow_ball;
    }, " ");
    $order_sql_set = substr($order_sql_set, 0, strlen($order_sql_set)-1);
    
    $where = " where `quest_order` in ({$order_sql_set}) ";

    return " {$where} ORDER BY `order` ";
}

function js_function($params) {
    global $allQuestOrder;

    $html = '
		$("select").change( function(){
			$(this).css("color","#0033FF");
});';

    $html .= choose_something($allQuestOrder, 'quest_order');

    return $html;
}
