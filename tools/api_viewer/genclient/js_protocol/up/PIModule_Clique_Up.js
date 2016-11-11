
// enum
var MOD_CLIQUE = {
		    
		  CREATE_CLIQUE_RESULT_SUCCESS: 0,
  CREATE_CLIQUE_RESULT_DUP_NAME: 1,
  CREATE_CLIQUE_RESULT_HAVE_CLIQUE: 2,

  APPLY_CLIQUE_RESULT_SUCCESS: 0,
  APPLY_CLIQUE_RESULT_ALREADY_JOIN: 1,
  APPLY_CLIQUE_RESULT_NOT_EXIST: 2,
  APPLY_CLIQUE_RESULT_REFUSE: 3,
  APPLY_CLIQUE_RESULT_TOO_MUCH_APPLY: 4,

  CANCEL_APPLY_CLIQUE_RESULT_SUCCESS: 0,
  CANCEL_APPLY_CLIQUE_RESULT_EXPIRED: 1,
  CANCEL_APPLY_CLIQUE_RESULT_NOT_EXIST: 2,

  PROCESS_JOIN_APPLY_RESULT_SUCCESS: 0,
  PROCESS_JOIN_APPLY_RESULT_EXPIRED: 1,
  PROCESS_JOIN_APPLY_RESULT_NO_ROOM: 2,
  PROCESS_JOIN_APPLY_RESULT_NO_PERMISSION: 3,
  PROCESS_JOIN_APPLY_RESULT_CANCEL_APPLY: 4,

  MANGE_MEMBER_ACTION_SET_OWNER: 0,
  MANGE_MEMBER_ACTION_SET_MANGER: 1,
  MANGE_MEMBER_ACTION_SET_MEMBER: 2,
  MANGE_MEMBER_ACTION_KICKOFF: 3,

  MANGE_MEMBER_RESULT_SUCCESS: 0,
  MANGE_MEMBER_RESULT_NOT_EXIST: 1,
  MANGE_MEMBER_RESULT_NO_PERMISSION: 2,

  CLIQUE_OPERA_ERROR_SUCCESS: 0,
  CLIQUE_OPERA_ERROR_CLIQUE_NOT_EXIST: 1,
  CLIQUE_OPERA_ERROR_NO_PERMISSION: 2,
  CLIQUE_OPERA_ERROR_MEMBER_NOT_EXIST: 3,
  CLIQUE_OPERA_ERROR_ALREADY_JOIN: 4,

  NOTIFY_LEAVE_CLIQUE_REASON_KICKOFF: 0,
  NOTIFY_LEAVE_CLIQUE_REASON_COLLAPSE: 1,
  NOTIFY_LEAVE_CLIQUE_REASON_LEAVE: 2,

  NOTIFY_JOINCLIQUE_FAILED_REASON_REFUSE: 0,
  NOTIFY_JOINCLIQUE_FAILED_REASON_EXPIRED: 1,
  NOTIFY_JOINCLIQUE_FAILED_REASON_NOROOM: 2,

  CLIQUE_RECUITMENT_RESULT_SUCCESS: 0,
  CLIQUE_RECUITMENT_RESULT_NO_PERMISSION: 1,
  CLIQUE_RECUITMENT_RESULT_CD: 2,

}

PIModule.Clique = PIModule.Clique || {};

PIModule.Clique.sendCreate_clique = function(name, announce) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,0, {
    "name":name,
    "announce":announce,
  });
};

PIModule.Clique.sendApply_join_clique = function(clique_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,1, {
    "clique_id":clique_id,
  });
};

PIModule.Clique.sendCancel_apply_clique = function(clique_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,2, {
    "clique_id":clique_id,
  });
};

PIModule.Clique.sendProcess_join_apply = function(agree, pidlist) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,3, {
    "agree":agree,
    "pidlist":pidlist,
  });
};

PIModule.Clique.sendElect_owner = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,4, {
  });
};

PIModule.Clique.sendMange_member = function(pid, action) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,5, {
    "pid":pid,
    "action":action,
  });
};

PIModule.Clique.sendDestory_clique = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,6, {
  });
};

PIModule.Clique.sendUpdate_announce = function(content) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,7, {
    "content":content,
  });
};

PIModule.Clique.sendLeave_clique = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,8, {
  });
};

PIModule.Clique.sendList_clique = function(offset, limit) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,9, {
    "offset":offset,
    "limit":limit,
  });
};

PIModule.Clique.sendClique_public_info = function(clique_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,10, {
    "clique_id":clique_id,
  });
};

PIModule.Clique.sendClique_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,11, {
  });
};

PIModule.Clique.sendList_apply = function(offset, limit) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,12, {
    "offset":offset,
    "limit":limit,
  });
};

PIModule.Clique.sendEnter_clubhouse = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,14, {
  });
};

PIModule.Clique.sendLeave_clubhouse = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,15, {
  });
};

PIModule.Clique.sendClub_move = function(to_x, to_y) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,16, {
    "to_x":to_x,
    "to_y":to_y,
  });
};

PIModule.Clique.sendClique_public_info_summary = function(clique_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,20, {
    "clique_id":clique_id,
  });
};

PIModule.Clique.sendClique_auto_audit = function(level, enable) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 33,21, {
    "level":level,
    "enable":enable,
  });
};

PIModule.Clique.sendClique_base_donate = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,22, {
  });
};

PIModule.Clique.sendClique_recruitment = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,26, {
  });
};

PIModule.Clique.sendQuick_apply = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 33,29, {
  });
};

PIModule.Clique.sendOther_clique = function(page) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 33,31, {
    "page":page,
  });
};

