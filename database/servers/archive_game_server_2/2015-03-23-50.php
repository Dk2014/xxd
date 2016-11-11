<?php 
$this->AddSQL("
	update skill set `music_sign`='QuSan', `sign`='ZhuFu' where id in (1004,1421);
"
);

$this->AddSQL("
	INSERT INTO `role` (`id`, `name`, `sign`, `type`, `is_special`, `skill_id1`, `skill_id2`, `buddy_level`, `mission_lock`, `scale`)
	VALUES
		(14, '替身男', 'TiShenNan', 2, 0, 0, 0, 0, 0, 100),
			(15, '替身女', 'TiShenNv', 2, 0, 0, 0, 0, 0, 100);

");
?>
