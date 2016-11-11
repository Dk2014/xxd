
// enum
var MOD_PLAYER = {
		    
		  LOGIN_STATUS_FAILED: 0,
  LOGIN_STATUS_SUCCEED: 1,
  LOGIN_STATUS_FIRST_TIME: 2,
  LOGIN_STATUS_WAIT_CLOSE: 3,
  LOGIN_STATUS_RESTORE_SUCCEED: 4,
  LOGIN_STATUS_BAN: 5,
  LOGIN_STATUS_RELOGIN: 6,
  LOGIN_STATUS_INVALID_PAYTOKEN: 7,

  PLATFORM_TYPE_MOBILE_IOS_WEIXIN: 1,
  PLATFORM_TYPE_MOBILE_IOS_QQ: 2,
  PLATFORM_TYPE_MOBILE_IOS_GUEST: 5,
  PLATFORM_TYPE_MOBILE_ANDROID_WEIXIN: 17,
  PLATFORM_TYPE_MOBILE_ANDROID_QQ: 18,
  PLATFORM_TYPE_MOBILE_ANDROID_GUEST: 21,

  CHANNEL_ID_MOBILE: 0,
  CHANNEL_ID_MOBILE_IOS_VN: 1,
  CHANNEL_ID_MOBILE_ANDROID_VN: 2,
  CHANNEL_ID_MOBILE_IOS_TW: 3,
  CHANNEL_ID_MOBILE_ANDROID_TW: 4,

  RANK_TYPE_FIGHTNUM: 0,
  RANK_TYPE_LEVEL: 1,
  RANK_TYPE_MAINROLE_FIGHTNUM: 2,
  RANK_TYPE_FAME: 4,
  RANK_TYPE_BUDDY_FIGHTNUM: 5,
  RANK_TYPE_GHOST_FIGHTNUM: 6,
  RANK_TYPE_SWORD_SOUL_NUM: 7,
  RANK_TYPE_NULL: 8,

}

PIModule.Player = PIModule.Player || {};

PIModule.Player.sendInfo = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,3, {
  });
};

PIModule.Player.sendRelogin = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,4, {
  });
};

PIModule.Player.sendBuy_physical = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,5, {
  });
};

PIModule.Player.sendSet_play_key = function(lock) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 0,6, {
    "lock":lock,
  });
};

PIModule.Player.sendGet_time = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,7, {
  });
};

PIModule.Player.sendFrom_platform_login = function(platform_type, channel_id, username, nickname, role_id, hashcode, unixtime, restore, recvCount, gsid, openkey, pay_token, pfkey, zoneid, pf, channel, telecom_oper, client_version, system_hardware, network) {
  PIProtocol.getInstance().send_packet( ((arguments[20]===undefined) ? 'def' : arguments[20]), 0,8, {
    "platform_type":platform_type,
    "channel_id":channel_id,
    "username":username,
    "nickname":nickname,
    "role_id":role_id,
    "hashcode":hashcode,
    "unixtime":unixtime,
    "restore":restore,
    "recvCount":recvCount,
    "gsid":gsid,
    "openkey":openkey,
    "pay_token":pay_token,
    "pfkey":pfkey,
    "zoneid":zoneid,
    "pf":pf,
    "channel":channel,
    "telecom_oper":telecom_oper,
    "client_version":client_version,
    "system_hardware":system_hardware,
    "network":network,
  });
};

PIModule.Player.sendBuy_coins = function(number) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 0,9, {
    "number":number,
  });
};

PIModule.Player.sendGet_login_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,10, {
  });
};

PIModule.Player.sendCross_login_game_server = function(pid, openid, nick, role_id, role_level, time, hash) {
  PIProtocol.getInstance().send_packet( ((arguments[7]===undefined) ? 'def' : arguments[7]), 0,20, {
    "pid":pid,
    "openid":openid,
    "nick":nick,
    "role_id":role_id,
    "role_level":role_level,
    "time":time,
    "hash":hash,
  });
};

PIModule.Player.sendGlobal_login = function(pid, openid, nick, role_id, role_level, fight_num, time, hash, gsid, platform_type, channel_id, username, openkey, pay_token, pfkey, zoneid, pf) {
  PIProtocol.getInstance().send_packet( ((arguments[17]===undefined) ? 'def' : arguments[17]), 0,26, {
    "pid":pid,
    "openid":openid,
    "nick":nick,
    "role_id":role_id,
    "role_level":role_level,
    "fight_num":fight_num,
    "time":time,
    "hash":hash,
    "gsid":gsid,
    "platform_type":platform_type,
    "channel_id":channel_id,
    "username":username,
    "openkey":openkey,
    "pay_token":pay_token,
    "pfkey":pfkey,
    "zoneid":zoneid,
    "pf":pf,
  });
};

PIModule.Player.sendGet_ingot = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 0,27, {
  });
};

PIModule.Player.sendSystem_fame = function(system) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 0,28, {
    "system":system,
  });
};

PIModule.Player.sendGet_ranks = function(flag, page_index) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 0,29, {
    "flag":flag,
    "page_index":page_index,
  });
};

