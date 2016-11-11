
// enum
var MOD_TOTEM = {
		    
		  CALL_TYPE_BONE: 0,
  CALL_TYPE_HALLOW: 1,
  CALL_TYPE_INGOT: 2,

  EQUIP_POS_POS1: 0,
  EQUIP_POS_POS2: 1,
  EQUIP_POS_POS3: 2,
  EQUIP_POS_POS4: 3,
  EQUIP_POS_POS5: 4,

}

PIModule.Totem = PIModule.Totem || {};

PIModule.Totem.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 31,0, {
  });
};

PIModule.Totem.sendCall_totem = function(call_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 31,1, {
    "call_type":call_type,
  });
};

PIModule.Totem.sendUpgrade = function(target_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 31,2, {
    "target_id":target_id,
  });
};

PIModule.Totem.sendDecompose = function(totem_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 31,3, {
    "totem_id":totem_id,
  });
};

PIModule.Totem.sendEquip = function(totem_id, equip_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 31,4, {
    "totem_id":totem_id,
    "equip_pos":equip_pos,
  });
};

PIModule.Totem.sendUnequip = function(equip_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 31,5, {
    "equip_pos":equip_pos,
  });
};

PIModule.Totem.sendEquip_pos_change = function(from_pos, to_pos) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 31,6, {
    "from_pos":from_pos,
    "to_pos":to_pos,
  });
};

PIModule.Totem.sendSwap = function(equiped_id, inbag_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 31,8, {
    "equiped_id":equiped_id,
    "inbag_id":inbag_id,
  });
};

PIModule.Totem.sendIs_bag_full = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 31,9, {
  });
};

