
// enum
var MOD_TOWN = {
		    
		  TOWNPLAYERSTATUS_NOMAL: 0,
  TOWNPLAYERSTATUS_PVP: 1,

}

PIModule.Town = PIModule.Town || {};

PIModule.Town.sendEnter = function(town_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 1,0, {
    "town_id":town_id,
  });
};

PIModule.Town.sendLeave = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 1,1, {
  });
};

PIModule.Town.sendMove = function(to_x, to_y) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 1,2, {
    "to_x":to_x,
    "to_y":to_y,
  });
};

PIModule.Town.sendTalked_npc_list = function(town_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 1,3, {
    "town_id":town_id,
  });
};

PIModule.Town.sendNpc_talk = function(npc_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 1,4, {
    "npc_id":npc_id,
  });
};

PIModule.Town.sendList_opened_town_treasures = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 1,8, {
  });
};

PIModule.Town.sendTake_town_treasures = function(town_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 1,9, {
    "town_id":town_id,
  });
};

PIModule.Town.sendUpdate_player_status = function(status) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 1,10, {
    "status":status,
  });
};

