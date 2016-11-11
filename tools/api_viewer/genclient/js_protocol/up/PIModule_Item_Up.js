
// enum
var MOD_ITEM = {
		    
		  EQUIPMENT_POS_WEAPON: 0,
  EQUIPMENT_POS_CLOTHES: 1,
  EQUIPMENT_POS_ACCESSORIES: 2,
  EQUIPMENT_POS_SHOE: 3,

  ATTRIBUTE_NULL: 0,
  ATTRIBUTE_ATTACK: 1,
  ATTRIBUTE_DEFENCE: 2,
  ATTRIBUTE_HEALTH: 3,
  ATTRIBUTE_SPEED: 4,
  ATTRIBUTE_CULTIVATION: 5,
  ATTRIBUTE_HIT_LEVEL: 6,
  ATTRIBUTE_CRITICAL_LEVEL: 7,
  ATTRIBUTE_BLOCK_LEVEL: 8,
  ATTRIBUTE_DESTROY_LEVEL: 9,
  ATTRIBUTE_TENACITY_LEVEL: 10,
  ATTRIBUTE_DODGE_LEVEL: 11,
  ATTRIBUTE_NUM: 11,

}

PIModule.Item = PIModule.Item || {};

PIModule.Item.sendGet_all_items = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 7,0, {
  });
};

PIModule.Item.sendDrop_item = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,1, {
    "id":id,
  });
};

PIModule.Item.sendBuy_item = function(item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,2, {
    "item_id":item_id,
  });
};

PIModule.Item.sendSell_item = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,3, {
    "id":id,
  });
};

PIModule.Item.sendDress = function(role_id, id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,4, {
    "role_id":role_id,
    "id":id,
  });
};

PIModule.Item.sendUndress = function(role_id, pos) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,5, {
    "role_id":role_id,
    "pos":pos,
  });
};

PIModule.Item.sendBuy_item_back = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,6, {
    "id":id,
  });
};

PIModule.Item.sendIs_bag_full = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 7,7, {
  });
};

PIModule.Item.sendDecompose = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,8, {
    "id":id,
  });
};

PIModule.Item.sendRefine = function(id, is_batch) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,9, {
    "id":id,
    "is_batch":is_batch,
  });
};

PIModule.Item.sendGet_recast_info = function(id, attr) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,10, {
    "id":id,
    "attr":attr,
  });
};

PIModule.Item.sendRecast = function(attr) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,11, {
    "attr":attr,
  });
};

PIModule.Item.sendUse_item = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,12, {
    "id":id,
  });
};

PIModule.Item.sendRole_use_cost_item = function(role_id, item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,13, {
    "role_id":role_id,
    "item_id":item_id,
  });
};

PIModule.Item.sendBatch_use_item = function(role_id, id, num) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 7,14, {
    "role_id":role_id,
    "id":id,
    "num":num,
  });
};

PIModule.Item.sendOpen_cornucopia = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,16, {
    "id":id,
  });
};

PIModule.Item.sendGet_sealedbooks = function(item_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 7,17, {
    "item_type":item_type,
  });
};

PIModule.Item.sendActivation_sealedbook = function(item_type, item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,18, {
    "item_type":item_type,
    "item_id":item_id,
  });
};

PIModule.Item.sendExchange_ghost_crystal = function(item_id, exchange_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 7,19, {
    "item_id":item_id,
    "exchange_type":exchange_type,
  });
};

