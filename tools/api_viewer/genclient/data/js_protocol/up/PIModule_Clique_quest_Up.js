
// enum
var MOD_CLIQUE_QUEST = {
		    
		}

PIModule.Clique_quest = PIModule.Clique_quest || {};

PIModule.Clique_quest.sendGet_clique_daily_Quest = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 35,1, {
  });
};

PIModule.Clique_quest.sendAward_clique_daily_Quest = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 35,2, {
    "id":id,
  });
};

PIModule.Clique_quest.sendGet_clique_building_quest = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 35,4, {
  });
};

PIModule.Clique_quest.sendAward_clique_building_Quest = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 35,5, {
    "id":id,
  });
};

