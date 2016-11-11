
// enum
var MOD_SKILL = {
		    
		  CHEAT_RESULT_SUCCESS: 0,
  CHEAT_RESULT_NO_DEPAND_SKILL: 1,
  CHEAT_RESULT_ALREADY_STUDY: 2,
  CHEAT_RESULT_CAN_NOT_STUDY_BEFORE: 3,
  CHEAT_RESULT_NOT_ROLE_SKILL: 4,
  CHEAT_RESULT_NOT_CHEAT_TYPE: 5,
  CHEAT_RESULT_ROLE_DOES_NOT_EXISTS: 6,
  CHEAT_RESULT_SKILL_NOT_MATCH_ROLE: 7,
  CHEAT_RESULT_LEVEL_NOT_REACHED: 8,

}

PIModule.Skill = PIModule.Skill || {};

PIModule.Skill.sendGet_all_skills_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 5,1, {
  });
};

PIModule.Skill.sendEquip_skill = function(role_id, order_number, skill_id) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 5,2, {
    "role_id":role_id,
    "order_number":order_number,
    "skill_id":skill_id,
  });
};

PIModule.Skill.sendUnequip_skill = function(role_id, order_number) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 5,3, {
    "role_id":role_id,
    "order_number":order_number,
  });
};

PIModule.Skill.sendStudy_skill_by_cheat = function(role_id, item_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 5,4, {
    "role_id":role_id,
    "item_id":item_id,
  });
};

PIModule.Skill.sendTrain_skill = function(role_id, skill_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 5,5, {
    "role_id":role_id,
    "skill_id":skill_id,
  });
};

PIModule.Skill.sendFlush_skill = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 5,6, {
    "role_id":role_id,
  });
};

