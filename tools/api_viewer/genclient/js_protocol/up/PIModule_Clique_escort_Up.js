
// enum
var MOD_CLIQUE_ESCORT = {
		    
		  BOAT_STATUS_CHANGE_MY_BOAT_HIJACKING: 0,
  BOAT_STATUS_CHANGE_MY_BOAT_HIJACKED: 1,
  BOAT_STATUS_CHANGE_MY_BOAT_RECOVERED: 2,
  BOAT_STATUS_CHANGE_HIJACKED_BOAT_RECOVERED: 3,
  BOAT_STATUS_CHANGE_ESCORT_FINISHED: 4,
  BOAT_STATUS_CHANGE_HIJACK_FINISHED: 5,

  RECOVER_BATTLE_WIN_RESULT_SUCCESS: 0,
  RECOVER_BATTLE_WIN_RESULT_BOAT_EXPIRE: 1,
  RECOVER_BATTLE_WIN_RESULT_NO_PERMISSION: 2,

  HIJACK_BATTLE_WIN_RESULT_SUCCESS: 0,
  HIJACK_BATTLE_WIN_RESULT_ESCORT_FINISHED: 1,
  HIJACK_BATTLE_WIN_RESULT_HIJACKED: 2,
  HIJACK_BATTLE_WIN_RESULT_NO_PERMISSION: 3,

  START_ESCORT_RESULT_SUCCESS: 0,
  START_ESCORT_RESULT_ILLEGAL_TIME: 1,
  START_ESCORT_RESULT_ESCORT_NOT_END: 2,
  START_ESCORT_RESULT_NO_CLIQUE: 3,
  START_ESCORT_RESULT_RUN_OUT: 4,
  START_ESCORT_RESULT_NO_BOAT: 5,

  HIJACK_BOAT_RESULT_START_BATTLE: 0,
  HIJACK_BOAT_RESULT_CLIQUE_MEMBER: 1,
  HIJACK_BOAT_RESULT_HIJACK_NOT_END: 2,
  HIJACK_BOAT_RESULT_NO_CLIQUE: 3,
  HIJACK_BOAT_RESULT_RUN_OUT: 4,
  HIJACK_BOAT_RESULT_NO_BOAT: 5,
  HIJACK_BOAT_RESULT_CAN_NOT_HIJACK: 6,

  RECOVER_BOAT_RESULT_START_BATTLE: 0,
  RECOVER_BOAT_RESULT_NO_PERMISSION: 1,
  RECOVER_BOAT_RESULT_RECOVERING: 2,
  RECOVER_BOAT_RESULT_NO_BOAT: 3,
  RECOVER_BOAT_RESULT_CAN_NOT_RECOVER: 4,

  ESCORT_STATUS_NONE: 0,
  ESCORT_STATUS_ESCORT: 1,
  ESCORT_STATUS_FINISHED: 2,

  BOAT_STATUS_ESCORT: 0,
  BOAT_STATUS_HIJACK: 1,
  BOAT_STATUS_RECOVER: 2,
  BOAT_STATUS_HIJACK_FINISH: 3,

  HIJACK_STATUS_NONE: 0,
  HIJACK_STATUS_HIJACK: 1,
  HIJACK_STATUS_FINISHED: 2,

}

PIModule.Clique_escort = PIModule.Clique_escort || {};

PIModule.Clique_escort.sendEscort_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,0, {
  });
};

PIModule.Clique_escort.sendGet_ingot_boat = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,1, {
  });
};

PIModule.Clique_escort.sendStart_escort = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,2, {
  });
};

PIModule.Clique_escort.sendHijack_boat = function(boat_id, pay) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 36,3, {
    "boat_id":boat_id,
    "pay":pay,
  });
};

PIModule.Clique_escort.sendRecover_boat = function(boat_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 36,4, {
    "boat_id":boat_id,
  });
};

PIModule.Clique_escort.sendList_boats = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,5, {
  });
};

PIModule.Clique_escort.sendGet_random_boat = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,6, {
  });
};

PIModule.Clique_escort.sendTake_hijack_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,11, {
  });
};

PIModule.Clique_escort.sendTake_escort_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,12, {
  });
};

PIModule.Clique_escort.sendGet_clique_boat_messages = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 36,13, {
  });
};

PIModule.Clique_escort.sendRead_clique_boat_message = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 36,15, {
    "id":id,
  });
};

