<?php

$extend_columns = array(
    /*'字段' => 配置 */

    'item_id' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
    ),
    
    'type' => array(
        'editor' => array('params' => array()),
        'render' => array('params' => array()),
        'data' => array(),
        'range' => array('params' => array()),
    ),
);

function sql_where($params) {
    if(isset($_GET['m'])){
	    return " where `type`={$_GET['m']} ";
    }
    return " order by `type`, `coins` ASC";
}

// 类型
$all_type = array(
    0 => '无',
    1 => '伙伴',
    2 => '魂侍',
    3 => '剑心',
    4 => '灵宠',
    5 => '装备',
    6 => '阵印',
    7 => '资源',
    8 => '战斗道具',
    9 => '武器',
    10 => '衣服',
    11 => '鞋',
    12 => '装饰',
);

function render_type($column_name, $column_val, $row, $data) {
    global $all_type;
    return $all_type[$column_val];
}

function editor_type($column_name, $column_val, $row, $data) {
    global $all_type;

    $data = $all_type;
    $field_name = 'type';

    $html = "<select class={$field_name} name={$field_name}[]>";

    foreach ($data as $key => $value) {
        if (isset($row[$field_name]) && $row[$field_name] == $key) {
            $html .= '<option value="'.$key.'" selected="selected" >'.$value.'</option>';
        } else {
            $html .= '<option value="'.$key.'">'.$value.'</option>';
        }
    }

    $html .= '</select>';
    return $html;
}

function range_type(){
    global $all_type;
    return $all_type;
}

function get_teoms() {
    global $db;
    $all_totem = array();

    $tmp = db_query($db, "select `id`, `name` from `totem` ");
    foreach ($tmp as $value)
    {
        $all_totem[$value['id']] = $value['name'];
    }

    return $all_totem;
}

function get_notequips() {

    global $db;
    $all_notequip = array();

    $tmp = db_query($db, "select `id`, `name` from `item` where `type_id` not in (3,4,5,6,8,10,11) ");
    foreach ($tmp as $value)
    {
        $all_notequip[$value['id']] = $value['name'];
    }

    return $all_notequip;
}

$all_roles = get_all_role(); //所有伙伴 
$all_ghosts = get_all_ghost();
$all_pets = get_all_pet();
$all_sword_souls = get_all_sword_soul();
$all_totems = get_teoms();
$all_notequips = get_notequips();

$all_equipments = get_all_item_by_types(array(3,4,5,6));
$all_battletools = get_all_item_by_types(array(8));

$all_items = get_all_item();

function editor_item_id($column_name, $column_val, $row, $data){
    $item_name = '';
    $item_id = $row['item_id'];
    switch ($row['type']) {


    case 1:
        global $all_roles;
        if (isset($all_roles[$item_id])) {
            $item_name = $all_roles[$item_id];
        }
	break;
    case 2:
        global $all_ghosts;
        if (isset($all_ghosts[$item_id])) {
            $item_name = $all_ghosts[$item_id];
        }
	break;
    case 3:
        global $all_sword_souls;
        if (isset($all_sword_souls[$item_id])) {
            $item_name = $all_sword_souls[$item_id];
        }
	break;
    case 4:
        global $all_pets;
        if (isset($all_pets[$item_id])) {
            $item_name = $all_pets[$item_id];
        }
    case 9:
    case 10:
    case 11:
    case 12:
        global $all_equipments;
        if (isset($all_equipments[$item_id])) {
            $item_name = $all_equipments[$item_id];
        }
	break;
    case 6:
        global $all_totems;
        if (isset($all_totems[$item_id])) {
            $item_name = $all_totems[$item_id];
        }
    case 7:
        global $all_notequips;
        if (isset($all_notequips[$item_id])) {
            $item_name = $all_notequips[$item_id];
        }
    
    case 8:
        global $all_battletools;
        if (isset($all_battletools[$item_id])) {
            $item_name = $all_battletools[$item_id];
        }
    }

    return editor_single_item($item_name, $item_id, "item_id");
}

function render_item_id($column_name, $column_val, $row, $data){
    if ($row['type'] == 1) {
        global $all_roles;
        return normal_render($all_roles, $row['item_id']);
    } else if ($row['type'] == 2) {
        global $all_ghosts;
        return normal_render($all_ghosts, $row['item_id']);
    } else if ($row['type'] == 3) {
        global $all_sword_souls;
        return normal_render($all_sword_souls, $row['item_id']);
    } else if ($row['type'] == 4) {
        global $all_pets;
        return normal_render($all_pets, $row['item_id']);
    } else if ($row['type'] == 9 ||
	    $row['type'] == 10  ||
	    $row['type'] == 11 ||
	    $row['type'] == 12 ) {
        global $all_equipments;
        return normal_render($all_equipments, $row['item_id']);
    } else if ($row['type'] == 6) {
        global $all_totems;
        return normal_render($all_totems, $row['item_id']);
    } else if ($row['type'] == 7) {
        global $all_notequips;
        return normal_render($all_notequips, $row['item_id']);
    }  else if ($row['type'] == 8) {
        global $all_battletools;
        return normal_render($all_battletools, $row['item_id']);
    }  
}

function js_function($params) {
    global $all_roles;
    global $all_sword_souls;
    global $all_ghosts;
    global $all_pets;
    global $all_totems;
    global $all_notequips;
    global $all_equipments;
    global $all_battletools;

    global $all_items;

    $column_name = 'item_id';

    $html = get_items_json($all_items, 'items');

    $html .= <<<EOT

    autoopt_{$column_name} = {
        lookup: items, 
        minChars: 0,
        onSelect: function(s) {
            $(this).siblings('input[class="real_value"]').val(s.id);
        }
    };

    $(".display_{$column_name}").autocomplete(autoopt_{$column_name});
EOT;
    $html .= get_items_json($all_roles, 'roles');
    $html .= get_items_json($all_ghosts, 'ghosts');
    $html .= get_items_json($all_sword_souls, 'sword_souls');
    $html .= get_items_json($all_pets, 'pets');
    $html .= get_items_json($all_totems, 'totems');
    $html .= get_items_json($all_notequips, 'notequips');
    $html .= get_items_json($all_equipments, 'equipments');
    $html .= get_items_json($all_battletools, 'battletools');
    

    $html .= <<<EOT

    $(".display_item_id").each(function(){
        switch (parseInt($(this).parent().parent().find(".type").val())) {
        case 1:
            autoopt_item_id.lookup = roles;
            break;
        case 2:
            autoopt_item_id.lookup = ghosts;
            break;
        case 3:
            autoopt_item_id.lookup = sword_souls;
            break;
        case 4:
            autoopt_item_id.lookup = pets;
            break;
	case 9:
	case 10:
	case 11:
	case 12:
            autoopt_item_id.lookup = equipments;
            break;
        case 6:
            autoopt_item_id.lookup = totems;
            break;
        case 7:
            autoopt_item_id.lookup = notequips;
            break;
        
        case 8:
            autoopt_item_id.lookup = battletools;
            break;
        default:
            autoopt_item_id.lookup = [{value:'无', id:'0'}];
        }
        $(this).autocomplete(autoopt_item_id);
    });
    
    $(".type").change(function() {
        switch (parseInt($(this).val())) {
        case 1:
            autoopt_item_id.lookup = roles;
            break;
        case 2:
            autoopt_item_id.lookup = ghosts;
            break;
        case 3:
            autoopt_item_id.lookup = sword_souls;
            break;
        case 4:
            autoopt_item_id.lookup = pets;
            break;
	case 9:
	case 10:
	case 11:
	case 12:
            autoopt_item_id.lookup = equipments;
            break;
        case 6:
            autoopt_item_id.lookup = totems;
            break;
        case 7:
            autoopt_item_id.lookup = notequips;
            break;
        
        case 8:
            autoopt_item_id.lookup = battletools;
            break;
        default:
            autoopt_item_id.lookup = [{value:'无', id:'0'}];
        }
        $(this).parent().parent().find(".display_item_id").autocomplete(autoopt_item_id);
    })

EOT;

return $html;
}

?>
