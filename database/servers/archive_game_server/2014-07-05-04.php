
<?php
db_execute($db, "
drop table if exists level_star;
create table level_star(
	`level_id` int(11) not null comment '关卡ID',
	`two_star_score` int(11) not null comment '两星要求分数',
	`three_star_score` int(11) not null comment '三星要求分数',
	PRIMARY KEY (`level_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='关卡星级分数表';
");
?>
