
// enum
var MOD_DAILY_SIGN_IN = {
		    
		  AWARD_TYPE_NEW_PLAYER_AWARD: 0,
  AWARD_TYPE_REGULAR_AWARD: 1,

}

PIModule.Daily_sign_in = PIModule.Daily_sign_in || {};

PIModule.Daily_sign_in.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 22,1, {
  });
};

PIModule.Daily_sign_in.sendSign = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 22,2, {
  });
};

PIModule.Daily_sign_in.sendSign_past_day = function(index) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 22,3, {
    "index":index,
  });
};

