
// enum
var MOD_ANNOUNCEMENT = {
		    
		}

PIModule.Announcement = PIModule.Announcement || {};

PIModule.Announcement.sendGet_list = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 18,0, {
  });
};

