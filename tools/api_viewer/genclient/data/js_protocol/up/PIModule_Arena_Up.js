
// enum
var MOD_ARENA = {
		    
		  NOTIFY_ARENA_MODE_ATTACKED_SUCC: 3,
  NOTIFY_ARENA_MODE_ATTACKED_FAILED: 4,

}

PIModule.Arena = PIModule.Arena || {};

PIModule.Arena.sendEnter = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 19,1, {
  });
};

PIModule.Arena.sendGet_records = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 19,2, {
  });
};

PIModule.Arena.sendAward_box = function(num) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 19,3, {
    "num":num,
  });
};

PIModule.Arena.sendStart_battle = function(player_id, player_rank) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 19,5, {
    "player_id":player_id,
    "player_rank":player_rank,
  });
};

PIModule.Arena.sendGet_top_rank = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 19,8, {
  });
};

PIModule.Arena.sendClean_failed_cd_time = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 19,9, {
  });
};

