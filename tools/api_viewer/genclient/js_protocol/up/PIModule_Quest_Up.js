
// enum
var MOD_QUEST = {
		    
		  GUIDE_TYPE_GUIDE_ENTER_MISSION: 0,
  GUIDE_TYPE_GUIDE_SKILL_USE: 1,
  GUIDE_TYPE_GUIDE_ADVANCED_SKILL_USE: 2,
  GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP: 3,
  GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_USE: 4,
  GUIDE_TYPE_GUIDE_EQUIP_REFINE: 5,
  GUIDE_TYPE_GUIDE_BATTLE_ITEM_ZHIXUECAO_USE: 6,
  GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_EQUIP: 7,
  GUIDE_TYPE_GUIDE_BUDDY_ADVANCED_SKILL_USE: 8,
  GUIDE_TYPE_GUIDE_GHOST_EQUIP: 9,
  GUIDE_TYPE_GUIDE_GHOST_POWER_LOOK: 10,
  GUIDE_TYPE_GUIDE_GHOST_BATTLE_USE: 11,
  GUIDE_TYPE_GUIDE_PET_EQUIP: 12,
  GUIDE_TYPE_GUIDE_PET_BATTLE_USE: 13,
  GUIDE_TYPE_GUIDE_PET_CATCH: 14,
  GUIDE_TYPE_GUIDE_HARD_LEVEL: 15,
  GUIDE_TYPE_GUIDE_TIANJIE: 16,
  GUIDE_TYPE_GUIDE_GOTO_ZHUBAO: 17,
  GUIDE_TYPE_GUIDE_SWORD_SOUL: 18,
  GUIDE_TYPE_GUIDE_FIRST_BATTLE: 19,
  GUIDE_TYPE_GUIDE_MAIN_ROLE_SKILL_EQUIP_FAILE: 20,
  GUIDE_TYPE_GUIDE_ENTER_MISSION_SECOND: 21,
  GUIDE_TYPE_GUIDE_BUDDY_EQUIP: 22,
  GUIDE_TYPE_GUIDE_BUDDY_USE_SKILL_DAODUN: 23,
  GUIDE_TYPE_GUIDE_EQUIP_USE_SKILL_FENGJUANCANSHENG: 24,
  GUIDE_TYPE_GUIDE_FRIENDSHIP: 25,
  GUIDE_TYPE_GUIDE_EQUIP_ROLE_3: 26,
  GUIDE_TYPE_GUIDE_EQUIP_ROLE_4: 27,

  GUIDE_ACTION_GUIDE_ACCEPT: 0,
  GUIDE_ACTION_GUIDE_FINISH: 1,

}

PIModule.Quest = PIModule.Quest || {};

PIModule.Quest.sendUpdate_quest = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 13,1, {
  });
};

PIModule.Quest.sendGet_daily_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 13,2, {
  });
};

PIModule.Quest.sendAward_daily = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,3, {
    "id":id,
  });
};

PIModule.Quest.sendGuide = function(guide_type, action) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 13,5, {
    "guide_type":guide_type,
    "action":action,
  });
};

PIModule.Quest.sendGet_extend_quest_info_by_npc_id = function(npc_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,6, {
    "npc_id":npc_id,
  });
};

PIModule.Quest.sendTake_extend_quest_award = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,7, {
    "quest_id":quest_id,
  });
};

PIModule.Quest.sendGet_pannel_quest_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 13,8, {
  });
};

PIModule.Quest.sendGive_up_addition_quest = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,10, {
    "quest_id":quest_id,
  });
};

PIModule.Quest.sendTake_addition_quest = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,11, {
    "quest_id":quest_id,
  });
};

PIModule.Quest.sendTake_addition_quest_award = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,12, {
    "quest_id":quest_id,
  });
};

PIModule.Quest.sendGet_addition_quest = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 13,13, {
  });
};

PIModule.Quest.sendRefresh_addition_quest = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,14, {
    "quest_id":quest_id,
  });
};

PIModule.Quest.sendTake_quest_stars_awaded = function(stars_level) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 13,15, {
    "stars_level":stars_level,
  });
};

