
// enum
var MOD_SWORD_SOUL = {
		    
		  DRAW_TYPE_COIN: 0,
  DRAW_TYPE_INGOT: 1,

  EQUIP_POS_POS1: 0,
  EQUIP_POS_POS2: 1,
  EQUIP_POS_POS3: 2,
  EQUIP_POS_POS4: 3,
  EQUIP_POS_POS5: 4,
  EQUIP_POS_POS6: 5,
  EQUIP_POS_POS7: 6,
  EQUIP_POS_POS8: 7,
  EQUIP_POS_POS9: 8,
  EQUIP_POS_NUM: 9,

  BOX_A: 0,
  BOX_B: 1,
  BOX_C: 2,
  BOX_D: 3,
  BOX_E: 4,

}

PIModule.Sword_soul = PIModule.Sword_soul || {};

PIModule.Sword_soul.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 10,0, {
  });
};

PIModule.Sword_soul.sendDraw = function(box, draw_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 10,1, {
    "box":box,
    "draw_type":draw_type,
  });
};

PIModule.Sword_soul.sendUpgrade = function(target_id, sword_souls) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 10,2, {
    "target_id":target_id,
    "sword_souls":sword_souls,
  });
};

PIModule.Sword_soul.sendExchange = function(sword_soul_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 10,3, {
    "sword_soul_id":sword_soul_id,
  });
};

PIModule.Sword_soul.sendEquip = function(from_id, role_id, equip_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 10,4, {
    "from_id":from_id,
    "role_id":role_id,
    "equip_pos":equip_pos,
  });
};

PIModule.Sword_soul.sendUnequip = function(role_id, from_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 10,5, {
    "role_id":role_id,
    "from_id":from_id,
  });
};

PIModule.Sword_soul.sendEquip_pos_change = function(role_id, from_id, to_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 10,6, {
    "role_id":role_id,
    "from_id":from_id,
    "to_pos":to_pos,
  });
};

PIModule.Sword_soul.sendSwap = function(role_id, from_id, to_id) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 10,8, {
    "role_id":role_id,
    "from_id":from_id,
    "to_id":to_id,
  });
};

PIModule.Sword_soul.sendIs_bag_full = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 10,9, {
  });
};

PIModule.Sword_soul.sendEmpty_pos_num = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 10,10, {
  });
};

