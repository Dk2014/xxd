<?php

$this->AddSQL("

ALTER TABLE `role_friendship`
	MODIFY `favourite_item` smallint(6) NOT NULL COMMENT '喜好品ID',
	MODIFY `favourite_count` int(11) NOT NULL COMMENT '喜好品需求量',
	MODIFY `level_color` varchar(20) NOT NULL COMMENT '名称颜色',
	MODIFY `display_graph` varchar(20) NOT NULL COMMENT '资源标识',
	MODIFY `relationship_name` varchar(20) NOT NULL COMMENT '羁绊名称';

");

