package battle

//怪物被动技能与光环
func getEnemyPassiveSkill(enemy *Fighter) ([]passiveSkill) {
	switch enemy.RoleId {
	case 0:
	return nil
	case 981:
		return []passiveSkill{passiveSkill{1760, 50, float32(0/100)}}
	case 982:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 983:
		return []passiveSkill{passiveSkill{1764, 5000, float32(0/100)}}
	case 986:
		return []passiveSkill{passiveSkill{1765, 10000, float32(0/100)}}
	case 998:
		return []passiveSkill{passiveSkill{1755, 8000, float32(0/100)}}
	case 1000:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 1004:
	return nil
	case 1006:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 1013:
		return []passiveSkill{passiveSkill{1766, 10000, float32(0/100)}}
	case 1024:
		return []passiveSkill{passiveSkill{1767, 6000, float32(0/100)}}
	case 1026:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 1036:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 1039:
		return []passiveSkill{passiveSkill{1768, 1, float32(0/100)}}
	case 1040:
		return []passiveSkill{passiveSkill{1769, 800, float32(0/100)}}
	case 1043:
		return []passiveSkill{passiveSkill{1762, 2000, float32(0/100)}}
	case 1044:
		return []passiveSkill{passiveSkill{1771, 20000, float32(0/100)}}
	case 1047:
		return []passiveSkill{passiveSkill{1772, 3720, float32(0/100)}}
	case 1049:
		return []passiveSkill{passiveSkill{1767, 3000, float32(0/100)}}
	case 1058:
		return []passiveSkill{passiveSkill{1773, 20000, float32(0/100)}}
	case 1069:
		return []passiveSkill{passiveSkill{1765, 50000, float32(0/100)}}
	case 1076:
		return []passiveSkill{passiveSkill{1774, 1000, float32(0/100)}}
	case 1092:
		return []passiveSkill{passiveSkill{1762, 5000, float32(0/100)}}
	case 1094:
	return nil
	case 1097:
		return []passiveSkill{passiveSkill{1754, 1000, float32(0/100)}}
	case 1098:
		return []passiveSkill{passiveSkill{1755, 2000, float32(0/100)}}
	case 1102:
		return []passiveSkill{passiveSkill{1756, 8000, float32(0/100)}}
	case 1107:
		return []passiveSkill{passiveSkill{1759, 1, float32(100/100)}}
	case 1108:
		return []passiveSkill{passiveSkill{1758, 1, float32(100/100)}}
	case 1116:
		return []passiveSkill{passiveSkill{1760, 100, float32(0/100)}}
	case 1120:
		return []passiveSkill{passiveSkill{1761, 5000, float32(0/100)}}
	case 1126:
		return []passiveSkill{passiveSkill{1762, 5000, float32(0/100)}}
	case 1136:
		return []passiveSkill{passiveSkill{1762, 5000, float32(0/100)}}
	case 1138:
		return []passiveSkill{passiveSkill{1781, 15000, float32(0/100)}}
	case 1143:
		return []passiveSkill{passiveSkill{1753, 4000, float32(100/100)}}
	case 1145:
		return []passiveSkill{passiveSkill{1763, 5000, float32(100/100)}}
	case 1149:
	return nil
	case 1150:
	return nil
	case 1151:
		return []passiveSkill{passiveSkill{1753, 5000, float32(100/100)}}
	case 1152:
		return []passiveSkill{passiveSkill{1762, 2000, float32(0/100)}}
	case 1153:
		return []passiveSkill{passiveSkill{1759, 1, float32(100/100)}}
	case 1154:
		return []passiveSkill{passiveSkill{1763, 1000, float32(100/100)}}
	case 1157:
		return []passiveSkill{passiveSkill{1759, 1, float32(100/100)},passiveSkill{1771, 10000, float32(100/100)}}
	case 1158:
		return []passiveSkill{passiveSkill{1769, 1000, float32(100/100)},passiveSkill{1771, 1000, float32(100/100)}}
	case 1159:
	return nil
	case 1160:
	return nil
	default:
		return nil
	}
	return nil
}

