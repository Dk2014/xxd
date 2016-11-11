<?php
$config = array(
	array( 'const',
	"POS0"        => 0, // 位置索引
	"POS1"        => 1,
	"POS2"        => 2,
	"POS3"        => 3,
	"POS4"        => 4,
	"POS5"        => 5,
	"POS6"        => 6,
	"POS7"        => 7,
	"POS8"        => 8,
	"POS_NO_ROLE" => -1, // -1 表示布阵没有角色

	"TEAMSHIP_HEALTH_IND" => 0,
	"TEAMSHIP_ATTACK_IND" => 1,
	"TEAMSHIP_DEFENCE_IND" => 2,

	"TEAMSHIP_MAX_LEVEL" => 100, // 伙伴配合最大等级
	),

	array('between', 'level' => 'int16', 'GetMaxInFormRoleNum' => 'int8', 'assert' => true, 'default' => 3,
	array('(,29]', 3),
	array('[30,49]', 4),
	array('[50,)', 5),
	),
); 
?>
