
// enum
var MOD_SERVER_INFO = {
		    
		}

PIModule.Server_info = PIModule.Server_info || {};

PIModule.Server_info.sendGet_version = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 88,0, {
  });
};

PIModule.Server_info.sendGet_api_count = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 88,1, {
  });
};

PIModule.Server_info.sendSearch_player_role = function(openid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 88,2, {
    "openid":openid,
  });
};

PIModule.Server_info.sendUpdate_access_token = function(token, pfkey) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 88,3, {
    "token":token,
    "pfkey":pfkey,
  });
};

PIModule.Server_info.sendUpdate_event_data = function(version) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 88,4, {
    "version":version,
  });
};

PIModule.Server_info.sendTss_data = function(data) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 88,5, {
    "data":data,
  });
};

PIModule.Server_info.sendPatch_info = function(openid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 88,6, {
    "openid":openid,
  });
};

