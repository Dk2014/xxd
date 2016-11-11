
// enum
var MOD_RAINBOW = {
		    
		}

PIModule.Rainbow = PIModule.Rainbow || {};

PIModule.Rainbow.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 23,1, {
  });
};

PIModule.Rainbow.sendReset = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 23,2, {
  });
};

PIModule.Rainbow.sendAward_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 23,3, {
  });
};

PIModule.Rainbow.sendTake_award = function(pos1, pos2) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 23,4, {
    "pos1":pos1,
    "pos2":pos2,
  });
};

PIModule.Rainbow.sendJump_to_segment = function(segment) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 23,5, {
    "segment":segment,
  });
};

PIModule.Rainbow.sendAuto_fight = function(segment) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 23,6, {
    "segment":segment,
  });
};

