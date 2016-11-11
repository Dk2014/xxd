
// enum
var MOD_BATTLE_PVP = {
		    
		}

PIModule.Battle_pvp = PIModule.Battle_pvp || {};

PIModule.Battle_pvp.sendInvite_fight = function(target_pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 40,0, {
    "target_pid":target_pid,
  });
};

PIModule.Battle_pvp.sendReply_invite = function(target_pid, reply_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 40,1, {
    "target_pid":target_pid,
    "reply_type":reply_type,
  });
};

PIModule.Battle_pvp.sendEnter_pvp = function(inviter_pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 40,2, {
    "inviter_pid":inviter_pid,
  });
};

PIModule.Battle_pvp.sendCancel_invite = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 40,3, {
  });
};

