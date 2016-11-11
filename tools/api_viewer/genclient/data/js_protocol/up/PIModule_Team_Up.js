
// enum
var MOD_TEAM = {
		    
		}

PIModule.Team = PIModule.Team || {};

PIModule.Team.sendGet_formation_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 2,0, {
  });
};

PIModule.Team.sendUp_formation = function(role_id, pos) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 2,2, {
    "role_id":role_id,
    "pos":pos,
  });
};

PIModule.Team.sendDown_formation = function(pos) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 2,3, {
    "pos":pos,
  });
};

PIModule.Team.sendSwap_formation = function(pos_from, pos_to) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 2,4, {
    "pos_from":pos_from,
    "pos_to":pos_to,
  });
};

PIModule.Team.sendReplace_formation = function(role_id, pos) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 2,5, {
    "role_id":role_id,
    "pos":pos,
  });
};

PIModule.Team.sendTraining_teamship = function(attr_ind) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 2,6, {
    "attr_ind":attr_ind,
  });
};

