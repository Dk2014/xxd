<?php

function foot() {
    require dirname(__FILE__).'/events_json_content.php';
}

function version_opera($row) {
    return '<a href="javascript:;" onclick="events_josn_editor(\''.$row['Type'].'\', \''.$row['Page'].'\',false)">配置奖品参数 </a>';
    
}

?>