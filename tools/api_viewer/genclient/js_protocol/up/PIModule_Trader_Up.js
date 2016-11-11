
// enum
var MOD_TRADER = {
		    
		}

PIModule.Trader = PIModule.Trader || {};

PIModule.Trader.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 21,1, {
  });
};

PIModule.Trader.sendStore_state = function(trader_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 21,2, {
    "trader_id":trader_id,
  });
};

PIModule.Trader.sendBuy = function(grid_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 21,3, {
    "grid_id":grid_id,
  });
};

PIModule.Trader.sendRefresh = function(trader_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 21,4, {
    "trader_id":trader_id,
  });
};

