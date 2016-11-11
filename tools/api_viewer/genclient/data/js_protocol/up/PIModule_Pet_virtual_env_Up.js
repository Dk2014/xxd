
// enum
var MOD_PET_VIRTUAL_ENV = {
		    
		}

PIModule.Pet_virtual_env = PIModule.Pet_virtual_env || {};

PIModule.Pet_virtual_env.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 28,1, {
  });
};

PIModule.Pet_virtual_env.sendTake_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 28,2, {
  });
};

PIModule.Pet_virtual_env.sendAuto_fight = function(floor) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 28,3, {
    "floor":floor,
  });
};

