<?php

function privilege_config($row) {
	return html_get_link("特权配置 | ", "?p=vip_privilege_config&m={$row['level']}", true);
}

function vip_levelup_gift($row) {
	return html_get_link("VIP升级奖励", "?p=vip_levelup_gift&m={$row['level']}", true);
}

?>
