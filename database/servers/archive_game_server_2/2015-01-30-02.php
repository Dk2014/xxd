<?php

$this->AddSQL("
alter table `skill` modify `info` text DEFAULT NULL COMMENT '绝招描述';
");

