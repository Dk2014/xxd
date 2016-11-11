
// enum
var MOD_VIP = {
		    
		  BUY_TIMES_TYPE_BIWUCHANGCISHU: 0,
  BUY_TIMES_TYPE_RAINBOWSAODANG: 1,

}

PIModule.Vip = PIModule.Vip || {};

PIModule.Vip.sendInfo = function(player_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 20,1, {
    "player_id":player_id,
  });
};

PIModule.Vip.sendGet_total = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 20,2, {
  });
};

PIModule.Vip.sendVip_level_total = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 20,3, {
  });
};

PIModule.Vip.sendBuy_times = function(buytype) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 20,4, {
    "buytype":buytype,
  });
};

