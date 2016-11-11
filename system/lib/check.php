<?php

function check_achievement($db, $row)
{
	if ($row['atype'] == 0) {
		return false;
	}

	switch ($row['second_type_id']) {
	case 3:
		if ($row['town_id'] == -1) {
			// 无城镇ID
			return false;
		}	
		$tmp = db_query($db, "select count(`id`) as `num` from quest where award_achievement_id={$row['id']}");
		if ($tmp[0]['num'] != 1) {
			// 与其他系统关联数据有问题
			return false;
		}
		break;
	case 4:
		if ($row['town_id'] == -1) {
			// 无城镇ID
			return false;
		}	
		$tmp = db_query($db, "select count(`id`) as `num` from challenge_mission_set where award_achievement_id={$row['id']}");
		if ($tmp[0]['num'] != 1) {
			// 与其他系统关联数据有问题
			return false;
		}
		break;
	case 5:
		if ($row['town_id'] == -1) {
			// 无城镇ID
			return false;
		}
		break;
	case 6:
		$tmp = db_query($db, "select count(`id`) as `num` from quest where award_achievement_id={$row['id']}");
		if ($tmp[0]['num'] != 1) {
			// 与其他系统关联数据有问题
			return false;
		}
		break;
	case 12:
		$tmp = db_query($db, "select count(`id`) as `num` from multi_mission where award_achievement_id={$row['id']}");
		if ($tmp[0]['num'] != 1) {
			// 与其他系统关联数据有问题
			return false;
		}
		break;
	case 9:
		if ($row['atype'] == 900) {
			if ($row['town_id'] == -1) {
				// 无城镇ID
				return false;
			}
		}
		break;
	case 10:
	case 11:
	case 13:
	case 14:
	case 15:
	case 16:
		return true;
		break;
	case 18:
		if ($row['atype'] == 1800) {
			if ($row['town_id'] == -1) {
				// 无城镇ID
				return false;
			}
		}
		break;
	case 19:
		if ($row['atype'] == 1900) {
			if ($row['bit'] == -1) {
				// 无位值
				return false;
			}
			$tmp = db_query($db, "select count(*) as num from skill where award_achievement_id={$row['id']}");
			if ($tmp[0]['num'] != 1) {
				return false;
			}
		}
		break;
	default:
		return false;
	}

	return true;
}