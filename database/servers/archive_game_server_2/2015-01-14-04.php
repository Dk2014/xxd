<?php

$this->AddSQL("
alter table `skill` add `info_vars` text COMMENT '绝招描述参量';
");

