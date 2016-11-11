
// enum
var MOD_FRIEND = {
		    
		  FRIEND_TYPE_GAME_FRIEND: 1,
  FRIEND_TYPE_PLATFORM_FRIEND: 2,

  FRIEND_MODE_STRANGE: 0,
  FRIEND_MODE_LISTEN_ONLY: 1,
  FRIEND_MODE_LISTENED_ONLY: 2,
  FRIEND_MODE_FRIEND: 3,

  ADD_RESULT_SUCCEED: 0,
  ADD_RESULT_FAILED_ADD_SELF: 1,
  ADD_RESULT_FAILED_ADD_NOT_EXIST: 2,
  ADD_RESULT_FAILED_ADD_FOLLOW: 3,
  ADD_RESULT_FAILED_ADD_FULL: 4,
  ADD_RESULT_FAILED_TARGET_FULL: 5,

  LISTEND_STATE_CANCEL: 1,
  LISTEND_STATE_LISTEND: 2,

  MSG_MODE_SEND: 1,
  MSG_MODE_RECEIVE: 2,

}

PIModule.Friend = PIModule.Friend || {};

PIModule.Friend.sendGet_friend_list = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 14,0, {
  });
};

PIModule.Friend.sendListen_by_nick = function(nick) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,1, {
    "nick":nick,
  });
};

PIModule.Friend.sendCancel_listen = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,2, {
    "pid":pid,
  });
};

PIModule.Friend.sendSend_heart = function(nickname, friend_type, pid) {
  PIProtocol.getInstance().send_packet( ((arguments[3]===undefined) ? 'def' : arguments[3]), 14,3, {
    "nickname":nickname,
    "friend_type":friend_type,
    "pid":pid,
  });
};

PIModule.Friend.sendChat = function(pid, message) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 14,4, {
    "pid":pid,
    "message":message,
  });
};

PIModule.Friend.sendGet_chat_history = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,5, {
    "pid":pid,
  });
};

PIModule.Friend.sendBlock = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,7, {
    "pid":pid,
  });
};

PIModule.Friend.sendCancel_block = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,8, {
    "pid":pid,
  });
};

PIModule.Friend.sendClean_chat_history = function(pid) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,9, {
    "pid":pid,
  });
};

PIModule.Friend.sendCurrent_platform_friend_num = function(num) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 14,11, {
    "num":num,
  });
};

PIModule.Friend.sendGet_send_heart_list = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 14,12, {
  });
};

PIModule.Friend.sendSend_heart_to_all_friends = function(friend_type, friends) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 14,13, {
    "friend_type":friend_type,
    "friends":friends,
  });
};

