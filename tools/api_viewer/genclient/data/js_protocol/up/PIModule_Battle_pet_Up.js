
// enum
var MOD_BATTLE_PET = {
		    
		}

PIModule.Battle_pet = PIModule.Battle_pet || {};

PIModule.Battle_pet.sendGet_pet_info = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 17,1, {
  });
};

PIModule.Battle_pet.sendSet_pet = function(grid_num, pet_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 17,3, {
    "grid_num":grid_num,
    "pet_id":pet_id,
  });
};

PIModule.Battle_pet.sendSet_pet_swap = function(from_grid_num, to_grid_num) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 17,4, {
    "from_grid_num":from_grid_num,
    "to_grid_num":to_grid_num,
  });
};

PIModule.Battle_pet.sendUpgrade_pet = function(pet_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 17,6, {
    "pet_id":pet_id,
  });
};

PIModule.Battle_pet.sendTraining_pet_skill = function(pet_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 17,7, {
    "pet_id":pet_id,
  });
};

