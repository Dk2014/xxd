
// enum
var MOD_CLIQUE_WAR = {
		    
		  SIGN_UP_CLIQUE_WAR_RESULT_FAILED: 0,
  SIGN_UP_CLIQUE_WAR_RESULT_SUCCEED: 1,
  SIGN_UP_CLIQUE_WAR_RESULT_REPEAT: 2,
  SIGN_UP_CLIQUE_WAR_RESULT_NO_CLIQUE: 3,
  SIGN_UP_CLIQUE_WAR_RESULT_NOT_MANAGER: 4,
  SIGN_UP_CLIQUE_WAR_RESULT_MEMBER_NOT_ENOUGH: 5,
  SIGN_UP_CLIQUE_WAR_RESULT_OUT_OF_TIME: 6,

}

PIModule.Clique_war = PIModule.Clique_war || {};

PIModule.Clique_war.sendWar_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 42,1, {
  });
};

PIModule.Clique_war.sendSign_up_clique_war = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 42,5, {
  });
};

PIModule.Clique_war.sendStart_clique_war_battle = function(target_playerid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 42,6, {
    "target_playerid":target_playerid,
  });
};

PIModule.Clique_war.sendSign_up_clique_war_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 42,7, {
  });
};

