
// enum
var MOD_TAOYUAN = {
		    
		  BLESS_RESULT_SUCCESS: 0,
  BLESS_RESULT_DAILY_BLESS_TIMES_LIMIT: 1,
  BLESS_RESULT_DAILY_BE_BLESS_TIMES_LIMIT: 2,
  BLESS_RESULT_BLESS_SAME_FRIENDS: 3,
  BLESS_RESULT_CAN_NOT_BLESS_SELF: 4,

  GROW_RESULT_SUCCESS: 0,
  GROW_RESULT_HAS_GROWED: 1,
  GROW_RESULT_FILED_NOT_OPEN: 2,
  GROW_RESULT_HEART_LEVEL_NOT_ENOUGH: 3,

  TAKE_RESULT_SUCCESS: 0,
  TAKE_RESULT_NOT_RIPE: 1,
  TAKE_RESULT_FILED_NOT_OPEN: 2,

  UPGRADE_FILED_RESULT_SUCCESS: 0,
  UPGRADE_FILED_RESULT_IS_BLACK: 1,
  UPGRADE_FILED_RESULT_FILED_NOT_OPEN: 2,
  UPGRADE_FILED_RESULT_HEART_LEVEL_NOT_ENOUGH: 3,

  OPEN_FILED_RESULT_SUCCESS: 0,
  OPEN_FILED_RESULT_FILED_HAS_OPEN: 1,
  OPEN_FILED_RESULT_HEART_LEVEL_NOT_ENOUGH: 2,

  QUEST_FINISH_RESULT_SUCCESS: 0,
  QUEST_FINISH_RESULT_QUEST_HAS_REFRESHED: 1,

  QUEST_STATE_RESULT_CAN_NOT_FINISH: 0,
  QUEST_STATE_RESULT_CAN_FINISH: 1,
  QUEST_STATE_RESULT_HAS_FINISHED: 2,

}

PIModule.Taoyuan = PIModule.Taoyuan || {};

PIModule.Taoyuan.sendTaoyuan_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,0, {
  });
};

PIModule.Taoyuan.sendBless = function(other_pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,1, {
    "other_pid":other_pid,
  });
};

PIModule.Taoyuan.sendShop_buy = function(item_id, num) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 39,2, {
    "item_id":item_id,
    "num":num,
  });
};

PIModule.Taoyuan.sendShop_sell = function(id, num) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 39,3, {
    "id":id,
    "num":num,
  });
};

PIModule.Taoyuan.sendGet_all_items = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,4, {
  });
};

PIModule.Taoyuan.sendFileds_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,5, {
  });
};

PIModule.Taoyuan.sendGrow_plant = function(filed_id, plant_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 39,6, {
    "filed_id":filed_id,
    "plant_id":plant_id,
  });
};

PIModule.Taoyuan.sendTake_plant = function(filed_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,7, {
    "filed_id":filed_id,
  });
};

PIModule.Taoyuan.sendUpgrade_filed = function(filed_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,8, {
    "filed_id":filed_id,
  });
};

PIModule.Taoyuan.sendOpen_filed = function(filed_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,9, {
    "filed_id":filed_id,
  });
};

PIModule.Taoyuan.sendStudy_skill = function(skill_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,10, {
    "skill_id":skill_id,
  });
};

PIModule.Taoyuan.sendMake_product = function(item_id, product_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 39,11, {
    "item_id":item_id,
    "product_type":product_type,
  });
};

PIModule.Taoyuan.sendTake_product = function(product_type, is_ingot) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 39,12, {
    "product_type":product_type,
    "is_ingot":is_ingot,
  });
};

PIModule.Taoyuan.sendProduct_make_queue = function(product_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,13, {
    "product_type":product_type,
  });
};

PIModule.Taoyuan.sendQuest_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,14, {
  });
};

PIModule.Taoyuan.sendQuest_finish = function(quset_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,15, {
    "quset_id":quset_id,
  });
};

PIModule.Taoyuan.sendQuest_refresh = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,16, {
  });
};

PIModule.Taoyuan.sendFriend_taoyuan_info = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,17, {
    "pid":pid,
  });
};

PIModule.Taoyuan.sendSkill_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,18, {
  });
};

PIModule.Taoyuan.sendOpen_queue = function(product_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,19, {
    "product_type":product_type,
  });
};

PIModule.Taoyuan.sendPlant_quickly_maturity = function(filed_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,20, {
    "filed_id":filed_id,
  });
};

PIModule.Taoyuan.sendTaoyuan_message_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 39,21, {
  });
};

PIModule.Taoyuan.sendTaoyuan_message_read = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,22, {
    "id":id,
  });
};

PIModule.Taoyuan.sendOpen_product_building = function(product_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 39,23, {
    "product_type":product_type,
  });
};

