
// enum
var MOD_FASHION = {
		    
		}

PIModule.Fashion = PIModule.Fashion || {};

PIModule.Fashion.sendFashion_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 25,1, {
  });
};

PIModule.Fashion.sendDress_fashion = function(fashion_id, in_clubhouse) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 25,2, {
    "fashion_id":fashion_id,
    "in_clubhouse":in_clubhouse,
  });
};

