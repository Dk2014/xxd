
// enum
var MOD_DRIVING_SWORD = {
		    
		  EXPLORING_MOUNTAIN_STATUS_UNEXPLORED: 0,
  EXPLORING_MOUNTAIN_STATUS_TREASURE_EMPTY: 1,
  EXPLORING_MOUNTAIN_STATUS_IN_GARRISON: 2,
  EXPLORING_MOUNTAIN_STATUS_BROKEN: 3,

  VISITING_MOUNTAIN_STATUS_UNWIN: 0,
  VISITING_MOUNTAIN_STATUS_WIN: 1,
  VISITING_MOUNTAIN_STATUS_AWARDED: 2,

  COMMON_EVENT_HOLE: 0,
  COMMON_EVENT_TELEPORT: 1,
  COMMON_EVENT_OBSTACLE: 2,
  COMMON_EVENT_UNKNOW_TELEPORT: 3,

  MOVING_DIRECTION_NORTH: 0,
  MOVING_DIRECTION_SOUTH: 1,
  MOVING_DIRECTION_WEST: 2,
  MOVING_DIRECTION_EAST: 3,

}

PIModule.Driving_sword = PIModule.Driving_sword || {};

PIModule.Driving_sword.sendCloud_map_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,0, {
  });
};

PIModule.Driving_sword.sendCloud_climb = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,1, {
  });
};

PIModule.Driving_sword.sendCloud_teleport = function(cloud) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 30,2, {
    "cloud":cloud,
  });
};

PIModule.Driving_sword.sendArea_teleport = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,3, {
  });
};

PIModule.Driving_sword.sendArea_move = function(direction) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 30,4, {
    "direction":direction,
  });
};

PIModule.Driving_sword.sendExplorer_start_battle = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,5, {
  });
};

PIModule.Driving_sword.sendExplorer_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,6, {
  });
};

PIModule.Driving_sword.sendExplorer_garrison = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 30,7, {
    "role_id":role_id,
  });
};

PIModule.Driving_sword.sendVisit_mountain = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,8, {
  });
};

PIModule.Driving_sword.sendVisiter_start_battle = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,9, {
  });
};

PIModule.Driving_sword.sendVisiter_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,10, {
  });
};

PIModule.Driving_sword.sendMountain_treasure_open = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,11, {
  });
};

PIModule.Driving_sword.sendList_garrisons = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,12, {
  });
};

PIModule.Driving_sword.sendAward_garrison = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 30,13, {
    "role_id":role_id,
  });
};

PIModule.Driving_sword.sendEnd_garrison = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 30,14, {
    "role_id":role_id,
  });
};

PIModule.Driving_sword.sendBuy_allowed_action = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 30,15, {
  });
};

