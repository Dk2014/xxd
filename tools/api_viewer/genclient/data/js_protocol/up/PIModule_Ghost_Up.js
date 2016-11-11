
// enum
var MOD_GHOST = {
		    
		  EQUIP_POS_POS1: 0,
  EQUIP_POS_POS2: 1,
  EQUIP_POS_POS3: 2,
  EQUIP_POS_POS4: 3,

  RESULT_TYPE_GHOST: 0,
  RESULT_TYPE_FRAGMENT: 1,
  RESULT_TYPE_FRUIT: 2,

}

PIModule.Ghost = PIModule.Ghost || {};

PIModule.Ghost.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 9,0, {
  });
};

PIModule.Ghost.sendEquip = function(from_id, role_id, to_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 9,1, {
    "from_id":from_id,
    "role_id":role_id,
    "to_pos":to_pos,
  });
};

PIModule.Ghost.sendUnequip = function(role_id, from_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 9,2, {
    "role_id":role_id,
    "from_id":from_id,
  });
};

PIModule.Ghost.sendSwap = function(role_id, id_bag, id_equip) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 9,3, {
    "role_id":role_id,
    "id_bag":id_bag,
    "id_equip":id_equip,
  });
};

PIModule.Ghost.sendEquip_pos_change = function(role_id, from_id, to_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 9,5, {
    "role_id":role_id,
    "from_id":from_id,
    "to_pos":to_pos,
  });
};

PIModule.Ghost.sendTrain = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,6, {
    "id":id,
  });
};

PIModule.Ghost.sendUpgrade = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,7, {
    "id":id,
  });
};

PIModule.Ghost.sendBaptize = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,8, {
    "id":id,
  });
};

PIModule.Ghost.sendCompose = function(ghost_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,11, {
    "ghost_id":ghost_id,
  });
};

PIModule.Ghost.sendTrain_skill = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,12, {
    "id":id,
  });
};

PIModule.Ghost.sendFlush_attr = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 9,13, {
    "id":id,
  });
};

PIModule.Ghost.sendRelation_ghost = function(master, slave) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 9,14, {
    "master":master,
    "slave":slave,
  });
};

