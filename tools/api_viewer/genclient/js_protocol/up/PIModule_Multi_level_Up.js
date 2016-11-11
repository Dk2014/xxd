
// enum
var MOD_MULTI_LEVEL = {
		    
		  ROOM_STATUS_SUCCESS: 0,
  ROOM_STATUS_FAILED_FULL: 1,
  ROOM_STATUS_FAILED_NOT_EXIST: 2,
  ROOM_STATUS_FAILED_FIGHTING: 3,
  ROOM_STATUS_FAILED_REQUIRE_LEVEL: 4,
  ROOM_STATUS_FAILED_REQUIRE_LOCK: 5,

}

PIModule.Multi_level = PIModule.Multi_level || {};

PIModule.Multi_level.sendCreate_room = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,1, {
    "level_id":level_id,
  });
};

PIModule.Multi_level.sendEnter_room = function(room_id, room_token) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 16,2, {
    "room_id":room_id,
    "room_token":room_token,
  });
};

PIModule.Multi_level.sendLeave_room = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,5, {
  });
};

PIModule.Multi_level.sendChange_buddy = function(buddy_role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,7, {
    "buddy_role_id":buddy_role_id,
  });
};

PIModule.Multi_level.sendGet_base_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,10, {
  });
};

PIModule.Multi_level.sendGet_online_friend = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,12, {
  });
};

PIModule.Multi_level.sendInvite_into_room = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,13, {
    "pid":pid,
  });
};

PIModule.Multi_level.sendRefuse_room_invite = function(room_id, inviter_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 16,16, {
    "room_id":room_id,
    "inviter_id":inviter_id,
  });
};

PIModule.Multi_level.sendMatch_player = function(level) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,19, {
    "level":level,
  });
};

PIModule.Multi_level.sendMatch_room = function(pos) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,22, {
    "pos":pos,
  });
};

PIModule.Multi_level.sendCancel_match_room = function(match_token) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,23, {
    "match_token":match_token,
  });
};

PIModule.Multi_level.sendGet_form_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,24, {
  });
};

PIModule.Multi_level.sendSet_form = function(formation) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 16,25, {
    "formation":formation,
  });
};

PIModule.Multi_level.sendGet_battle_role_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,26, {
  });
};

PIModule.Multi_level.sendGet_match_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,28, {
  });
};

PIModule.Multi_level.sendCancel_match_player = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 16,30, {
  });
};

