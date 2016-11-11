
// enum
var MOD_CHANNEL = {
		    
		  MESSAGE_TYPE_CHAT: 0,
  MESSAGE_TYPE_RARE_PROPS: 1,
  MESSAGE_TYPE_CLIQUE_MESSAGE: 2,
  MESSAGE_TYPE_CLIQUE_CHAT: 3,
  MESSAGE_TYPE_CLIQUE_NEWS: 4,
  MESSAGE_TYPE_DESPAIR_LAND_BOSS_KILLED: 5,

}

PIModule.Channel = PIModule.Channel || {};

PIModule.Channel.sendGet_latest_world_channel_message = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 29,0, {
  });
};

PIModule.Channel.sendAdd_world_chat = function(content) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 29,1, {
    "content":content,
  });
};

PIModule.Channel.sendWorld_channel_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 29,2, {
  });
};

PIModule.Channel.sendAdd_clique_chat = function(content) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 29,4, {
    "content":content,
  });
};

PIModule.Channel.sendGet_latest_clique_messages = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 29,5, {
  });
};

