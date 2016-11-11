
// enum
var MOD_DRAW = {
		    
		  EXCHANGE_GIFT_RESULT_SUCCESS: 0,
  EXCHANGE_GIFT_RESULT_EXPIRE: 1,
  EXCHANGE_GIFT_RESULT_DUP_EXCHANGE: 2,

  CHEST_TYPE_COIN: 0,
  CHEST_TYPE_COIN_FREE: 1,
  CHEST_TYPE_COIN_TEN: 2,
  CHEST_TYPE_INGOT: 3,
  CHEST_TYPE_INGOT_FREE: 4,
  CHEST_TYPE_INGOT_TEN: 5,
  CHEST_TYPE_PET: 6,
  CHEST_TYPE_PET_FREE: 7,
  CHEST_TYPE_PET_TEN: 8,

  ITEM_TYPE_COIN: 1,
  ITEM_TYPE_INGOT: 2,
  ITEM_TYPE_ITEM: 3,
  ITEM_TYPE_GHOST: 4,
  ITEM_TYPE_SWORD_SOUL: 5,
  ITEM_TYPE_PET: 6,
  ITEM_TYPE_GHOST_FRAGMENT: 7,
  ITEM_TYPE_PREFERENCE: 8,
  ITEM_TYPE_EQUIPMENT: 9,

}

PIModule.Draw = PIModule.Draw || {};

PIModule.Draw.sendGet_heart_draw_info = function(draw_type, award_record) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 50,1, {
    "draw_type":draw_type,
    "award_record":award_record,
  });
};

PIModule.Draw.sendHeart_draw = function(draw_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 50,2, {
    "draw_type":draw_type,
  });
};

PIModule.Draw.sendGet_chest_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 50,3, {
  });
};

PIModule.Draw.sendDraw_chest = function(chest_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 50,4, {
    "chest_type":chest_type,
  });
};

PIModule.Draw.sendHeart_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 50,5, {
  });
};

PIModule.Draw.sendGet_fate_box_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 50,6, {
  });
};

PIModule.Draw.sendOpen_fate_box = function(box_type, times) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 50,7, {
    "box_type":box_type,
    "times":times,
  });
};

PIModule.Draw.sendExchange_gift_code = function(code) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 50,8, {
    "code":code,
  });
};

