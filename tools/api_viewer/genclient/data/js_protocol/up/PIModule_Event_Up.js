
// enum
var MOD_EVENT = {
		    
		}

PIModule.Event = PIModule.Event || {};

PIModule.Event.sendLogin_award_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,1, {
  });
};

PIModule.Event.sendTake_login_award = function(day) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 24,2, {
    "day":day,
  });
};

PIModule.Event.sendGet_events = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,3, {
  });
};

PIModule.Event.sendGet_event_award = function(event_id, page, param1, param2, param3) {
  PIProtocol.getInstance().send_packet( ((arguments[5]===undefined) ? 'def' : arguments[5]), 24,4, {
    "event_id":event_id,
    "page":page,
    "param1":param1,
    "param2":param2,
    "param3":param3,
  });
};

PIModule.Event.sendPlayer_event_physical_cost = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,6, {
  });
};

PIModule.Event.sendPlayer_month_card_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,7, {
  });
};

PIModule.Event.sendGet_seven_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,8, {
  });
};

PIModule.Event.sendGet_seven_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,9, {
  });
};

PIModule.Event.sendGet_richman_club_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,10, {
  });
};

PIModule.Event.sendGet_richman_club_award = function(column, sequence) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 24,11, {
    "column":column,
    "sequence":sequence,
  });
};

PIModule.Event.sendInfo_share = function(is_share) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 24,12, {
    "is_share":is_share,
  });
};

PIModule.Event.sendInfo_group_buy = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,13, {
  });
};

PIModule.Event.sendGet_ingot_change_total = function(is_in) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 24,14, {
    "is_in":is_in,
  });
};

PIModule.Event.sendGet_event_total_award = function(order) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 24,15, {
    "order":order,
  });
};

PIModule.Event.sendGet_event_arena_rank = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,16, {
  });
};

PIModule.Event.sendGet_event_ten_draw_times = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,17, {
  });
};

PIModule.Event.sendGet_event_recharge_award = function(page, requireid) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 24,18, {
    "page":page,
    "requireid":requireid,
  });
};

PIModule.Event.sendGet_event_new_year = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 24,19, {
  });
};

PIModule.Event.sendQq_vip_continue = function(kind) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 24,20, {
    "kind":kind,
  });
};

