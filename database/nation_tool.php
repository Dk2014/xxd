<?php
// 区分当前发布版本 -- 国服 or 越南 .etc
function _S() {  // switch
    $target_nation = exec("hg branch");
    switch ($target_nation) {
    case "tw"://台湾
    return func_get_arg(3);
    case "tencent":   // 国服
	return func_get_arg(1);
    case "soha":      // 越南
	return func_get_arg(2);
    default:
	return func_get_arg(0); 
    }
}
?>
