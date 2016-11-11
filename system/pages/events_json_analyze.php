<?php
$jsonfile="../../json_events.json";
// 拉取
if(isset($_GET['type'])  && isset($_GET['page']) ){
    $type=$_GET['type'];
    $page=$_GET['page'];

    $josn_string = file_get_contents($jsonfile);
   
    $obj=json_decode($josn_string);
    $data=$obj->JsonEvents;
    foreach ( $data as $unit ) {
        if ($type == $unit->Type  && $page == $unit->Page) { 
            $require = $unit->AwardList;  
        }
    }
    
    echo json_encode($require, JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE);
}

    
// 提交
if(isset($_POST['type'])  && isset($_POST['page']) ){
    
    $type=$_POST['type'];
    $page=$_POST['page'];
    $content=json_decode($_POST['content'], true);

    $josn_string = file_get_contents($jsonfile);
   
    $obj=json_decode($josn_string);
    $data=$obj->JsonEvents;
    foreach ( $data as $unit ) {
        if ($type == $unit->Type  && $page == $unit->Page) { 
            //replace the content 
            $unit->AwardList =  $content;
        }
    }
    $obj->JsonEvents=$data;
    file_put_contents($jsonfile, json_encode($obj,JSON_PRETTY_PRINT|JSON_UNESCAPED_UNICODE));
}

   