
// enum
var MOD_MISSION = {
		    
		  OUT_RESULT_FAILED: 0,
  OUT_RESULT_SUCCEED: 1,

  EXTEND_TYPE_RESOURCE: 1,
  EXTEND_TYPE_ACTIVITY: 2,

  EXTEND_LEVEL_TYPE_RESOURCE: 1,
  EXTEND_LEVEL_TYPE_BUDDY: 9,
  EXTEND_LEVEL_TYPE_PET: 10,
  EXTEND_LEVEL_TYPE_GHOST: 11,

  EXTEND_LEVEL_SUB_TYPE_NONE: 0,
  EXTEND_LEVEL_SUB_TYPE_COIN: 1,
  EXTEND_LEVEL_SUB_TYPE_EXP: 2,

  BATTLE_TYPE_MISSION: 0,
  BATTLE_TYPE_RESOURCE: 1,
  BATTLE_TYPE_TOWER: 2,
  BATTLE_TYPE_MULTILEVEL: 3,
  BATTLE_TYPE_HARD: 8,
  BATTLE_TYPE_BUDDY: 9,
  BATTLE_TYPE_PET: 10,
  BATTLE_TYPE_GHOST: 11,
  BATTLE_TYPE_RAINBOW: 12,
  BATTLE_TYPE_PVE: 13,
  BATTLE_TYPE_DESPAIR: 19,

}

PIModule.Mission = PIModule.Mission || {};

PIModule.Mission.sendOpen = function(mission_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,0, {
    "mission_id":mission_id,
  });
};

PIModule.Mission.sendGet_mission_level = function(mission_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,1, {
    "mission_id":mission_id,
  });
};

PIModule.Mission.sendEnter_level = function(mission_level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,2, {
    "mission_level_id":mission_level_id,
  });
};

PIModule.Mission.sendOpen_level_box = function(times) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,3, {
    "times":times,
  });
};

PIModule.Mission.sendUse_ingot_relive = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 4,4, {
  });
};

PIModule.Mission.sendUse_item = function(role_id, item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 4,5, {
    "role_id":role_id,
    "item_id":item_id,
  });
};

PIModule.Mission.sendEnter_extend_level = function(level_type, level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 4,7, {
    "level_type":level_type,
    "level_id":level_id,
  });
};

PIModule.Mission.sendGet_extend_level_info = function(level_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,8, {
    "level_type":level_type,
  });
};

PIModule.Mission.sendOpen_small_box = function(box_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,9, {
    "box_id":box_id,
  });
};

PIModule.Mission.sendEnter_hard_level = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,11, {
    "level_id":level_id,
  });
};

PIModule.Mission.sendGet_hard_level = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 4,12, {
  });
};

PIModule.Mission.sendGet_buddy_level_role_id = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 4,13, {
  });
};

PIModule.Mission.sendSet_buddy_level_team = function(role_pos, buddy_pos, tactical) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 4,14, {
    "role_pos":role_pos,
    "buddy_pos":buddy_pos,
    "tactical":tactical,
  });
};

PIModule.Mission.sendAuto_fight_level = function(level_type, level_id, times) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 4,15, {
    "level_type":level_type,
    "level_id":level_id,
    "times":times,
  });
};

PIModule.Mission.sendEnter_rainbow_level = function(mission_level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,16, {
    "mission_level_id":mission_level_id,
  });
};

PIModule.Mission.sendOpen_meng_yao = function(meng_yao_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,17, {
    "meng_yao_id":meng_yao_id,
  });
};

PIModule.Mission.sendGet_mission_level_by_item_id = function(item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,18, {
    "item_id":item_id,
  });
};

PIModule.Mission.sendBuy_hard_level_times = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,19, {
    "level_id":level_id,
  });
};

PIModule.Mission.sendOpen_random_award_box = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,20, {
    "level_id":level_id,
  });
};

PIModule.Mission.sendEvaluate_level = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 4,21, {
  });
};

PIModule.Mission.sendOpen_shaded_box = function(shaded_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,22, {
    "shaded_id":shaded_id,
  });
};

PIModule.Mission.sendGet_mission_total_star_info = function(town_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,23, {
    "town_id":town_id,
  });
};

PIModule.Mission.sendGet_mission_total_star_awards = function(town_id, box_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 4,24, {
    "town_id":town_id,
    "box_type":box_type,
  });
};

PIModule.Mission.sendClear_mission_level_state = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 4,25, {
  });
};

PIModule.Mission.sendBuy_resource_mission_level_times = function(sub_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 4,26, {
    "sub_type":sub_type,
  });
};

