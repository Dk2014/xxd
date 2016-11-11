<?php
$config = array(
	array( 'const',
		'HAS_BEEN_READ' => 1, //已读
		'UN_READ' => 0, // 未读
		'HAVE_ATTACHMENT' => 1, // 有附件
		'NO_ATTACHMENT' => 0, // 没有附件

		'NO_ATTACHMENT_SAVE_DAYS' => 3, // 没有附件保存天数
		'HAVE_ATTACHMENT_SAVE_DAYS' => 7, // 有附件保存天数
		'CLIQUE_HAVE_ATTACHMENT_SAVE_DAYS' => 14, // 帮派工资附件保存天数
		'HEART_MAIL_SAVE_DAYS' => 3, //爱心邮件保存时间
		'HEART_MAIL_ID' => 2, //爱心邮件ID
		'CLIQUE_SALARY_OWNER_MAIL_ID' => 37,//帮主工资邮件ID
		'CLIQUE_SALARY_MANAGER_MAIL_ID' => 38,//副帮主工资邮件ID
		'CLIQUE_SALARY_MAIL_OWNER_COIN_BASE' => 500, //帮主工资基数
		'CLIQUE_SALARY_MAIL_MANAGER_COIN_BASE' => 100, //副帮主工资基数
	),

	// 邮件附件类型
	array('const', 
		'ATTACHMENT_ITEM'  => 0,
		'ATTACHMENT_COINS' => 1, //  '铜钱',
		'ATTACHMENT_INGOT' => 2, //  '元宝',
		'ATTACHMENT_HEART' => 3, //  '爱心',
		'ATTACHMENT_FORMATION_EXP' => 4, // '经验(所有上阵角色)',
		'ATTACHMENT_SINGLE_ROLE_EXP' => 5, //  '经验(单个角色)',
		'ATTACHMENT_GHOST' => 6, //  '魂侍',
		'ATTACHMENT_SWORD_SOUL' => 7, // 剑心,
		'ATTACHMENT_HEART_FROM_FRIEND' => 8, // 还有赠送的爱心
		'ATTACHMENT_RELATIONSHIP' => 9, // 友情值
		'ATTACHMENT_FAME' => 10, // 声望
		'ATTACHMENT_TOTEM' => 11, // 阵印
	),

	// 邮件删除时机
	array('const', 
		'AUTO_DELETE_MAGIC_NUM' => 10, //自动删除魔数，大于这个数字的认为是制定时间删除
		'AUTO_DELETE_AFTER_EXPIRED'  => 0,
		'AUTO_DELETE_AFTER_READ_WITHOUT_ATTACHMENT'  => 1,
	),

);
?>
