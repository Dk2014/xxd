
// enum
var MOD_MONEY_TREE = {
		    
		}

PIModule.Money_tree = PIModule.Money_tree || {};

PIModule.Money_tree.sendGet_tree_status = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 32,0, {
  });
};

PIModule.Money_tree.sendGet_tree_money = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 32,1, {
  });
};

PIModule.Money_tree.sendWave_tree = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 32,2, {
  });
};

