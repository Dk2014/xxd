
// enum
var MOD_AWAKEN = {
		    
		}

PIModule.Awaken = PIModule.Awaken || {};

PIModule.Awaken.sendAwaken_info = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 38,0, {
    "role_id":role_id,
  });
};

PIModule.Awaken.sendLevelup_attr = function(role_id, attr_impl) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 38,1, {
    "role_id":role_id,
    "attr_impl":attr_impl,
  });
};

