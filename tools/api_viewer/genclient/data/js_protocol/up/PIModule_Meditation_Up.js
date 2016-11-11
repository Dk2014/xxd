
// enum
var MOD_MEDITATION = {
		    
		}

PIModule.Meditation = PIModule.Meditation || {};

PIModule.Meditation.sendMeditation_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 27,1, {
  });
};

PIModule.Meditation.sendStart_meditation = function(in_clubhouse) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 27,2, {
    "in_clubhouse":in_clubhouse,
  });
};

PIModule.Meditation.sendStop_meditation = function(in_clubhouse) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 27,3, {
    "in_clubhouse":in_clubhouse,
  });
};

