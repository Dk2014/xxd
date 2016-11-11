<?php

$this->AddSQL("
alter table `town` add column `exit_x` int(11) comment '出口x坐标';
alter table `town` add column `exit_y` int(11) comment '出口y坐标';
");

