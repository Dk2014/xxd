
// enum
var MOD_ROLE = {
		    
		  ATTRIBUTE_NULL: 0,
  ATTRIBUTE_ATTACK: 1,
  ATTRIBUTE_DEFENCE: 2,
  ATTRIBUTE_HEALTH: 3,
  ATTRIBUTE_SPEED: 4,
  ATTRIBUTE_CULTIVATION: 5,
  ATTRIBUTE_HIT_LEVEL: 6,
  ATTRIBUTE_CRITICAL_LEVEL: 7,
  ATTRIBUTE_BLOCK_LEVEL: 8,
  ATTRIBUTE_DESTROY_LEVEL: 9,
  ATTRIBUTE_TENACITY_LEVEL: 10,
  ATTRIBUTE_DODGE_LEVEL: 11,
  ATTRIBUTE_NUM: 11,

  FIGHTNUM_TYPE_ALL: 0,
  FIGHTNUM_TYPE_ROLE_LEVEL: 1,
  FIGHTNUM_TYPE_SKILL: 2,
  FIGHTNUM_TYPE_EQUIP: 3,
  FIGHTNUM_TYPE_GHOST: 4,
  FIGHTNUM_TYPE_REALM: 5,
  FIGHTNUM_TYPE_FASHION: 8,
  FIGHTNUM_TYPE_FRIENDSHIP: 9,
  FIGHTNUM_TYPE_TEAMSHIP: 10,
  FIGHTNUM_TYPE_CLIQUE_KONGFU: 11,

}

PIModule.Role = PIModule.Role || {};

PIModule.Role.sendGet_all_roles = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 3,0, {
  });
};

PIModule.Role.sendGet_role_info = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,1, {
    "role_id":role_id,
  });
};

PIModule.Role.sendGet_player_info = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,2, {
    "pid":pid,
  });
};

PIModule.Role.sendGet_fight_num = function(fight_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,3, {
    "fight_type":fight_type,
  });
};

PIModule.Role.sendGet_player_info_with_openid = function(openid, game_server_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 3,4, {
    "openid":openid,
    "game_server_id":game_server_id,
  });
};

PIModule.Role.sendLevelup_role_friendship = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,5, {
    "role_id":role_id,
  });
};

PIModule.Role.sendRecruit_buddy = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,6, {
    "role_id":role_id,
  });
};

PIModule.Role.sendGet_role_fight_num = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,7, {
    "role_id":role_id,
  });
};

PIModule.Role.sendChange_role_status = function(role_id, status) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 3,8, {
    "role_id":role_id,
    "status":status,
  });
};

PIModule.Role.sendGet_inn_role_list = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 3,9, {
  });
};

PIModule.Role.sendBuddy_coop = function(coop_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 3,10, {
    "coop_id":coop_id,
  });
};

