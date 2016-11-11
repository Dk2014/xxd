
// enum
var MOD_CLIQUE_BUILDING = {
		    
		  CLIQUE_KONGFU_TRAIN_RESULT_SUCCESS: 0,
  CLIQUE_KONGFU_TRAIN_RESULT_LACK_CONTRIB: 1,
  CLIQUE_KONGFU_TRAIN_RESULT_NO_CLIQUE: 2,
  CLIQUE_KONGFU_TRAIN_RESULT_MAX_LEVEL: 3,

  CLIQUE_BANK_SOLD_RESULT_SUCCESS: 0,
  CLIQUE_BANK_SOLD_RESULT_CD: 1,
  CLIQUE_BANK_SOLD_RESULT_NO_CLIQUE: 2,
  CLIQUE_BANK_SOLD_RESULT_MAX_LEVEL: 3,

  CLIQUE_BUILDING_DONATE_RESULT_SUCCESS: 0,
  CLIQUE_BUILDING_DONATE_RESULT_FAILED: 1,

  CLIQUE_STORE_CHEST_JUNLIANG: 1,
  CLIQUE_STORE_CHEST_BAOXIANG: 2,

  CLIQUE_STORE_SEND_RESULT_SUCCESS: 0,
  CLIQUE_STORE_SEND_RESULT_LACK_COINS: 1,
  CLIQUE_STORE_SEND_RESULT_NO_CLIQUE: 2,
  CLIQUE_STORE_SEND_RESULT_NOT_MANAGER: 3,
  CLIQUE_STORE_SEND_RESULT_TIMES_NOT_ENOUGH: 4,
  CLIQUE_STORE_SEND_RESULT_CLIQUE_NOT_FOUND: 5,
  CLIQUE_STORE_SEND_RESULT_CLIQUE_CHEST_NOT_FOUND: 6,

}

PIModule.Clique_building = PIModule.Clique_building || {};

PIModule.Clique_building.sendClique_base_donate = function(money) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,1, {
    "money":money,
  });
};

PIModule.Clique_building.sendClique_building_status = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 34,2, {
  });
};

PIModule.Clique_building.sendClique_bank_donate = function(money) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,3, {
    "money":money,
  });
};

PIModule.Clique_building.sendClique_bank_buy = function(kind, num) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 34,4, {
    "kind":kind,
    "num":num,
  });
};

PIModule.Clique_building.sendClique_bank_sold = function(kind) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,5, {
    "kind":kind,
  });
};

PIModule.Clique_building.sendClique_kongfu_donate = function(building, money) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 34,6, {
    "building":building,
    "money":money,
  });
};

PIModule.Clique_building.sendClique_kongfu_info = function(building) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,7, {
    "building":building,
  });
};

PIModule.Clique_building.sendClique_kongfu_train = function(kongfu_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,8, {
    "kongfu_id":kongfu_id,
  });
};

PIModule.Clique_building.sendClique_temple_worship = function(worship_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,9, {
    "worship_type":worship_type,
  });
};

PIModule.Clique_building.sendClique_temple_donate = function(money) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,10, {
    "money":money,
  });
};

PIModule.Clique_building.sendClique_temple_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 34,11, {
  });
};

PIModule.Clique_building.sendClique_store_donate = function(money) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,12, {
    "money":money,
  });
};

PIModule.Clique_building.sendClique_store_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 34,13, {
  });
};

PIModule.Clique_building.sendClique_store_send_chest = function(store_chest_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 34,14, {
    "store_chest_type":store_chest_type,
  });
};

