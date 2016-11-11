
// enum
var MOD_DESPAIR_LAND = {
		    
		  DESPAIR_LAND_CAMP_BEAST: 1,
  DESPAIR_LAND_CAMP_WALKING_DEAD: 2,
  DESPAIR_LAND_CAMP_DEVIL: 3,

  CAMP_BOSS_STATUS_CLOSING: 0,
  CAMP_BOSS_STATUS_OPENING: 1,
  CAMP_BOSS_STATUS_FINISHED: 2,

  DESPAIR_LAND_BATTLE_RECORD_TYPE_FIRST: 1,
  DESPAIR_LAND_BATTLE_RECORD_TYPE_LATEST: 2,
  DESPAIR_LAND_BATTLE_RECORD_TYPE_LOW_FIGHT_NUM: 3,

}

PIModule.Despair_land = PIModule.Despair_land || {};

PIModule.Despair_land.sendDespair_land_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 37,1, {
  });
};

PIModule.Despair_land.sendDespair_land_camp_info = function(camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 37,2, {
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_camp_player_info = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 37,3, {
    "level_id":level_id,
  });
};

PIModule.Despair_land.sendDespair_land_pick_box = function(level_id, camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 37,4, {
    "level_id":level_id,
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_watch_record = function(record_type, camp_type, level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 37,5, {
    "record_type":record_type,
    "camp_type":camp_type,
    "level_id":level_id,
  });
};

PIModule.Despair_land.sendDespair_land_buy_battle_num = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 37,6, {
  });
};

PIModule.Despair_land.sendDespair_land_boss_info = function(camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 37,7, {
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_buy_boss_battle_num = function(camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 37,10, {
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_pick_three_star_box = function(level_id, camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 37,12, {
    "level_id":level_id,
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_battle_boss_award_info = function(camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 37,13, {
    "camp_type":camp_type,
  });
};

PIModule.Despair_land.sendDespair_land_boss_open_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 37,14, {
  });
};

