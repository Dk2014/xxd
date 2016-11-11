
// enum
var MOD_CLIQUE_BOSS = {
		    
		  REFRESH_CHALLENGE_RESULT_FAILED: 0,
  REFRESH_CHALLENGE_RESULT_SUCCEED: 1,
  REFRESH_CHALLENGE_RESULT_DO_NOT_NEED_REFRESH: 2,

}

PIModule.Clique_boss = PIModule.Clique_boss || {};

PIModule.Clique_boss.sendBoss_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 41,1, {
  });
};

PIModule.Clique_boss.sendClique_boss_refresh_challenge = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 41,5, {
  });
};

