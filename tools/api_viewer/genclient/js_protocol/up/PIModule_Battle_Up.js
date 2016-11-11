
// enum
var MOD_BATTLE = {
		    
		  FIGHTER_TYPE_ATK: 0,
  FIGHTER_TYPE_DEF: 1,

  FIGHTER_KIND_PLAYER: 0,
  FIGHTER_KIND_BUDDY: 1,
  FIGHTER_KIND_ENEMY: 2,

  ROUND_EVENT_NONE: 0,
  ROUND_EVENT_DODGE: 1,
  ROUND_EVENT_CRIT: 2,
  ROUND_EVENT_BLOCK: 3,
  ROUND_EVENT_SQUELCH: 4,

  ROUND_STATUS_NOT_END: 0,
  ROUND_STATUS_ATK_WIN: 1,
  ROUND_STATUS_DEF_WIN: 2,
  ROUND_STATUS_ATK_NEXT: 3,
  ROUND_STATUS_DEF_NEXT: 4,
  ROUND_STATUS_TRIGGER_CALL_ENEMYS: 5,
  ROUND_STATUS_WAITING: 6,

  BUFF_MODE_POWER: 0,
  BUFF_MODE_SPEED: 1,
  BUFF_MODE_ATTACK: 2,
  BUFF_MODE_DEFEND: 3,
  BUFF_MODE_HEALTH: 4,
  BUFF_MODE_DIZZINESS: 5,
  BUFF_MODE_POISONING: 6,
  BUFF_MODE_CLEAN_BAD: 7,
  BUFF_MODE_CLEAN_GOOD: 8,
  BUFF_MODE_REDUCE_HURT: 9,
  BUFF_MODE_RANDOM: 10,
  BUFF_MODE_BLOCK: 11,
  BUFF_MODE_BLOCK_LEVEL: 12,
  BUFF_MODE_DODGE_LEVEL: 13,
  BUFF_MODE_CRITIAL_LEVEL: 14,
  BUFF_MODE_HIT_LEVEL: 15,
  BUFF_MODE_HURT_ADD: 16,
  BUFF_MODE_MAX_HEALTH: 17,
  BUFF_MODE_KEEPER_REDUCE_HURT: 18,
  BUFF_MODE_ATTRACT_FIRE: 19,
  BUFF_MODE_DESTROY_LEVEL: 20,
  BUFF_MODE_TENACITY_LEVEL: 21,
  BUFF_MODE_SUNDER: 22,
  BUFF_MODE_SLEEP: 23,
  BUFF_MODE_DISABLE_SKILL: 24,
  BUFF_MODE_BUFF_DIRECT_REDUCE_HURT: 25,
  BUFF_MODE_BUFF_ABSORB_HURT: 26,
  BUFF_MODE_BUFF_GHOST_POWER: 27,
  BUFF_MODE_BUFF_PET_LIVE_ROUND: 28,
  BUFF_MODE_BUFF_BUDDY_SKILL: 29,
  BUFF_MODE_BUFF_CLEAR_ABSORB_HURT: 30,
  BUFF_MODE_BUFF_SLEEP_LEVEL: 31,
  BUFF_MODE_BUFF_DIZZINESS_LEVEL: 32,
  BUFF_MODE_BUFF_RANDOM_LEVEL: 33,
  BUFF_MODE_BUFF_DISABLE_SKILL_LEVEL: 34,
  BUFF_MODE_BUFF_POISONING_LEVEL: 35,
  BUFF_MODE_BUFF_RECOVER_BUDDY_SKILL: 36,
  BUFF_MODE_BUFF_MAKE_POWER_FULL: 37,
  BUFF_MODE_BUFF_DOGE: 38,
  BUFF_MODE_BUFF_HIT: 39,
  BUFF_MODE_BUFF_CRITIAL: 40,
  BUFF_MODE_BUFF_TENACITY: 41,
  BUFF_MODE_BUFF_TAKE_SUNSER: 42,
  BUFF_MODE_BUFF_DEFEND_PERSENT: 43,
  BUFF_MODE_BUFF_SUNDER_STATE: 44,
  BUFF_MODE_BUFF_HEALTH_PERCENT: 45,
  BUFF_MODE_BUFF_ALL_RESIST: 46,
  BUFF_MODE_BUFF_REBOTH_HEALTH: 47,
  BUFF_MODE_BUFF_REBOTH_HEALTH_PERCENT: 48,

  BATTLE_TYPE_MISSION: 0,
  BATTLE_TYPE_RESOURCE: 1,
  BATTLE_TYPE_TOWER: 2,
  BATTLE_TYPE_MULTILEVEL: 3,
  BATTLE_TYPE_ARENA: 4,
  BATTLE_TYPE_HARD: 8,
  BATTLE_TYPE_BUDDY: 9,
  BATTLE_TYPE_PET: 10,
  BATTLE_TYPE_GHOST: 11,
  BATTLE_TYPE_RAINBOW: 12,
  BATTLE_TYPE_PVE: 13,
  BATTLE_TYPE_FATE_BOX: 14,
  BATTLE_TYPE_DRIVING_EXPLORING: 15,
  BATTLE_TYPE_DRIVING_SWORD_BF_LEVEL: 16,
  BATTLE_TYPE_HIJACK_BOAT: 17,
  BATTLE_TYPE_RECOVER_BOAT: 18,
  BATTLE_TYPE_DESPAIR: 19,
  BATTLE_TYPE_DESPAIR_BOSS: 20,
  BATTLE_TYPE_PVP: 21,
  BATTLE_TYPE_CLIQUE_BOSS: 22,

  JOB_TYPE_NONE: 0,
  JOB_TYPE_ATTACKER: 1,
  JOB_TYPE_DESTROYER: 2,
  JOB_TYPE_DEFENDER: 3,
  JOB_TYPE_TREATER: 4,
  JOB_TYPE_SUPPORTER: 5,
  JOB_TYPE_OBSTRUCTOR: 6,

}

PIModule.Battle = PIModule.Battle || {};

PIModule.Battle.sendStart_battle = function(battle_type, battle_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 6,0, {
    "battle_type":battle_type,
    "battle_id":battle_id,
  });
};

PIModule.Battle.sendNext_round = function(use_skill, use_item, auto_fight, is_attacker, position, job_index, send_num, use_sword_soul, use_ghost_skill_position, use_ghost_skill_id, use_totem) {
  PIProtocol.getInstance().send_packet( ((arguments[11]===undefined) ? 'def' : arguments[11]), 6,1, {
    "use_skill":use_skill,
    "use_item":use_item,
    "auto_fight":auto_fight,
    "is_attacker":is_attacker,
    "position":position,
    "job_index":job_index,
    "send_num":send_num,
    "use_sword_soul":use_sword_soul,
    "use_ghost_skill_position":use_ghost_skill_position,
    "use_ghost_skill_id":use_ghost_skill_id,
    "use_totem":use_totem,
  });
};

PIModule.Battle.sendEscape = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 6,3, {
  });
};

PIModule.Battle.sendStart_ready = function(ok) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 6,6, {
    "ok":ok,
  });
};

PIModule.Battle.sendCall_battle_pet = function(grid_num) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 6,8, {
    "grid_num":grid_num,
  });
};

PIModule.Battle.sendUse_buddy_skill = function(pos, use_skill) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 6,9, {
    "pos":pos,
    "use_skill":use_skill,
  });
};

PIModule.Battle.sendStart_battle_for_hijack_boat = function(pid, boat_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 6,12, {
    "pid":pid,
    "boat_id":boat_id,
  });
};

PIModule.Battle.sendStart_battle_for_recover_boat = function(pid, boat_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 6,13, {
    "pid":pid,
    "boat_id":boat_id,
  });
};

PIModule.Battle.sendRound_ready = function(is_auto) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 6,14, {
    "is_auto":is_auto,
  });
};

PIModule.Battle.sendInit_round = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 6,15, {
  });
};

PIModule.Battle.sendSet_auto = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 6,17, {
  });
};

PIModule.Battle.sendCancel_auto = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 6,18, {
  });
};

PIModule.Battle.sendSet_skill = function(is_attacker, pos_idx, skill_idx) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 6,19, {
    "is_attacker":is_attacker,
    "pos_idx":pos_idx,
    "skill_idx":skill_idx,
  });
};

PIModule.Battle.sendUse_item = function(is_attacker, position, item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 6,20, {
    "is_attacker":is_attacker,
    "position":position,
    "item_id":item_id,
  });
};

PIModule.Battle.sendUse_ghost = function(is_attacker, position) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 6,21, {
    "is_attacker":is_attacker,
    "position":position,
  });
};

PIModule.Battle.sendBattle_reconnect = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 6,23, {
  });
};

