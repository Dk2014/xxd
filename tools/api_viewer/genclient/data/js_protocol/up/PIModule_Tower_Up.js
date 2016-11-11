
// enum
var MOD_TOWER = {
		    
		}

PIModule.Tower = PIModule.Tower || {};

PIModule.Tower.sendGet_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 15,1, {
  });
};

PIModule.Tower.sendUse_ladder = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 15,2, {
  });
};

