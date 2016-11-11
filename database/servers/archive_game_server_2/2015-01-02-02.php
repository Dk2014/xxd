<?php

$max_unique_key_query = $this->NewQuery("select max(`unique_key`) as `max_key` from `func`");
$max_unique_key_row = $max_unique_key_query->GoNext();
$max_unique_key = intval($max_unique_key_row['max_key']);
$this->DropQuery($max_unique_key_query);
$max_bit = intval(log($max_unique_key,2));
echo "{$max_bit}, {$max_unique_key}";

$level_func_query = $this->NewQuery("select * from `level_func`");

while($level_func = $level_func_query->GoNext()) {
	$name = $level_func['name'];
	$sign = $level_func['sign'];
	$level = $level_func['level'];
	$need_play = $level_func['need_play'];
	$max_bit+=1;
	$unique_key = 1<<$max_bit;
	$this->AddSQL("insert into `func` (`name`, `sign`, `level`, `need_play`, `unique_key`, `lock`) " . 
	       	"values('{$name}', '{$sign}', '{$level}', '{$need_play}', '{$unique_key}', 0);");
}

$this->DropQuery($level_func_query);

$this->AddSQL("update `func` set `type`=2 where `level`>0");
?>
