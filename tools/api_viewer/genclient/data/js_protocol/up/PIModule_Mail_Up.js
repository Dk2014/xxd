
// enum
var MOD_MAIL = {
		    
		}

PIModule.Mail = PIModule.Mail || {};

PIModule.Mail.sendGet_list = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 12,0, {
  });
};

PIModule.Mail.sendRead = function(id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 12,1, {
    "id":id,
  });
};

PIModule.Mail.sendTake_attachment = function(attachment_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 12,2, {
    "attachment_id":attachment_id,
  });
};

PIModule.Mail.sendGet_infos = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 12,3, {
  });
};

PIModule.Mail.sendRequest_global_mail = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 12,4, {
  });
};

