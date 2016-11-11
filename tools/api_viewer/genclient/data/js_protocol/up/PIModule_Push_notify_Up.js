
// enum
var MOD_PUSH_NOTIFY = {
		    
		}

PIModule.Push_notify = PIModule.Push_notify || {};

PIModule.Push_notify.sendPush_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 26,1, {
  });
};

PIModule.Push_notify.sendPush_notification_switch = function(notification_id, turn_on) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 26,2, {
    "notification_id":notification_id,
    "turn_on":turn_on,
  });
};

