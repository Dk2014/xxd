
// enum
var MOD_DEBUG = {
		    
		}

PIModule.Debug = PIModule.Debug || {};

PIModule.Debug.sendAdd_buddy = function(role_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,0, {
    "role_id":role_id,
  });
};

PIModule.Debug.sendAdd_item = function(item_id, number) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,2, {
    "item_id":item_id,
    "number":number,
  });
};

PIModule.Debug.sendSet_role_level = function(role_id, level) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,3, {
    "role_id":role_id,
    "level":level,
  });
};

PIModule.Debug.sendSet_coins = function(number) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,4, {
    "number":number,
  });
};

PIModule.Debug.sendSet_ingot = function(number) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,5, {
    "number":number,
  });
};

PIModule.Debug.sendAdd_ghost = function(ghost_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,11, {
    "ghost_id":ghost_id,
  });
};

PIModule.Debug.sendSet_player_physical = function(physical) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,12, {
    "physical":physical,
  });
};

PIModule.Debug.sendReset_level_enter_count = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,13, {
    "level_id":level_id,
  });
};

PIModule.Debug.sendAdd_exp = function(role_id, add_exp) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,14, {
    "role_id":role_id,
    "add_exp":add_exp,
  });
};

PIModule.Debug.sendOpen_ghost_mission = function(mission_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,15, {
    "mission_id":mission_id,
  });
};

PIModule.Debug.sendSend_mail = function(mail_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,16, {
    "mail_id":mail_id,
  });
};

PIModule.Debug.sendClear_mail = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,17, {
  });
};

PIModule.Debug.sendOpen_mission_level = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,18, {
    "level_id":level_id,
  });
};

PIModule.Debug.sendStart_battle = function(battle_type, enemy_id) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,19, {
    "battle_type":battle_type,
    "enemy_id":enemy_id,
  });
};

PIModule.Debug.sendListen_by_name = function(name) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,20, {
    "name":name,
  });
};

PIModule.Debug.sendOpen_quest = function(quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,21, {
    "quest_id":quest_id,
  });
};

PIModule.Debug.sendOpen_func = function(lock) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,22, {
    "lock":lock,
  });
};

PIModule.Debug.sendAdd_sword_soul = function(sword_soul_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,23, {
    "sword_soul_id":sword_soul_id,
  });
};

PIModule.Debug.sendAdd_battle_pet = function(petId) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,25, {
    "petId":petId,
  });
};

PIModule.Debug.sendReset_multi_level_enter_count = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,26, {
  });
};

PIModule.Debug.sendOpen_multi_level = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,27, {
    "level_id":level_id,
  });
};

PIModule.Debug.sendOpen_all_pet_grid = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,28, {
  });
};

PIModule.Debug.sendCreate_announcement = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,29, {
  });
};

PIModule.Debug.sendAdd_heart = function(number) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,30, {
    "number":number,
  });
};

PIModule.Debug.sendReset_hard_level_enter_count = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,31, {
  });
};

PIModule.Debug.sendOpen_hard_level = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,32, {
    "level_id":level_id,
  });
};

PIModule.Debug.sendSet_vip_level = function(level) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,33, {
    "level":level,
  });
};

PIModule.Debug.sendSet_resource_level_open_day = function(level_type, open_day) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,34, {
    "level_type":level_type,
    "open_day":open_day,
  });
};

PIModule.Debug.sendReset_resource_level_open_day = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,35, {
  });
};

PIModule.Debug.sendReset_arena_daily_count = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,36, {
  });
};

PIModule.Debug.sendReset_sword_soul_draw_cd = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,37, {
  });
};

PIModule.Debug.sendSet_first_login_time = function(timestamp) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,38, {
    "timestamp":timestamp,
  });
};

PIModule.Debug.sendEarlier_first_login_time = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,39, {
  });
};

PIModule.Debug.sendReset_server_open_time = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,40, {
  });
};

PIModule.Debug.sendClear_trader_refresh_time = function(trader_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,41, {
    "trader_id":trader_id,
  });
};

PIModule.Debug.sendAdd_trader_refresh_time = function(trader_id, timing) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,42, {
    "trader_id":trader_id,
    "timing":timing,
  });
};

PIModule.Debug.sendClear_trader_schedule = function(trader_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,43, {
    "trader_id":trader_id,
  });
};

PIModule.Debug.sendAdd_trader_schedule = function(trader_id, expire, showup, disappear) {
  PIProtocol.getInstance().send_packet( ((arguments[4]===undefined) ? 'def' : arguments[4]), 99,44, {
    "trader_id":trader_id,
    "expire":expire,
    "showup":showup,
    "disappear":disappear,
  });
};

PIModule.Debug.sendOpen_town = function(town_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,45, {
    "town_id":town_id,
  });
};

PIModule.Debug.sendAdd_global_mail = function(send_delay, expire_delay) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,46, {
    "send_delay":send_delay,
    "expire_delay":expire_delay,
  });
};

PIModule.Debug.sendCreate_announcement_without_tpl = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,47, {
  });
};

PIModule.Debug.sendSet_login_day = function(days) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,48, {
    "days":days,
  });
};

PIModule.Debug.sendReset_login_award = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,49, {
  });
};

PIModule.Debug.sendRest_player_award_lock = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,50, {
  });
};

PIModule.Debug.sendReset_rainbow_level = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,51, {
  });
};

PIModule.Debug.sendSet_rainbow_level = function(segment, order) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,52, {
    "segment":segment,
    "order":order,
  });
};

PIModule.Debug.sendSend_push_notification = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,53, {
  });
};

PIModule.Debug.sendReset_pet_virtual_env = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,54, {
  });
};

PIModule.Debug.sendAdd_fame = function(system, val) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,55, {
    "system":system,
    "val":val,
  });
};

PIModule.Debug.sendAdd_world_chat_message = function(num) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,56, {
    "num":num,
  });
};

PIModule.Debug.sendMonth_card = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,57, {
  });
};

PIModule.Debug.sendEnter_sandbox = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,58, {
  });
};

PIModule.Debug.sendSandbox_rollback = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,59, {
  });
};

PIModule.Debug.sendExit_sandbox = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,60, {
  });
};

PIModule.Debug.sendReset_shaded_missions = function(level_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,61, {
    "level_id":level_id,
  });
};

PIModule.Debug.sendClean_cornucopia = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,62, {
  });
};

PIModule.Debug.sendAdd_totem = function(totem_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,63, {
    "totem_id":totem_id,
  });
};

PIModule.Debug.sendAdd_rune = function(jade_num, rock_num) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,64, {
    "jade_num":jade_num,
    "rock_num":rock_num,
  });
};

PIModule.Debug.sendSend_rare_item_message = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,65, {
  });
};

PIModule.Debug.sendAdd_sword_driving_action = function(point) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,66, {
    "point":point,
  });
};

PIModule.Debug.sendReset_driving_sword_data = function(cloud) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,67, {
    "cloud":cloud,
  });
};

PIModule.Debug.sendAdd_sword_soul_fragment = function(number) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,68, {
    "number":number,
  });
};

PIModule.Debug.sendReset_money_tree_status = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,69, {
  });
};

PIModule.Debug.sendReset_today_money_tree = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,70, {
  });
};

PIModule.Debug.sendClean_sword_soul_ingot_draw_nums = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,71, {
  });
};

PIModule.Debug.sendPunch_driving_sword_cloud = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,72, {
  });
};

PIModule.Debug.sendClear_clique_daily_donate = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,73, {
  });
};

PIModule.Debug.sendSet_clique_contrib = function(contrib) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,74, {
    "contrib":contrib,
  });
};

PIModule.Debug.sendRefresh_clique_worship = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,75, {
  });
};

PIModule.Debug.sendClique_escort_hijack_battle_win = function(boat_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,76, {
    "boat_id":boat_id,
  });
};

PIModule.Debug.sendClique_escort_recover_battle_win = function(boat_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,77, {
    "boat_id":boat_id,
  });
};

PIModule.Debug.sendClique_escort_notify_message = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,87, {
  });
};

PIModule.Debug.sendClique_escort_notify_daily_quest = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,88, {
  });
};

PIModule.Debug.sendSet_clique_building_level = function(building, level) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,89, {
    "building":building,
    "level":level,
  });
};

PIModule.Debug.sendSet_clique_building_money = function(building, money) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,90, {
    "building":building,
    "money":money,
  });
};

PIModule.Debug.sendEscort_bench = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,91, {
  });
};

PIModule.Debug.sendReset_clique_escort_daily_num = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,92, {
  });
};

PIModule.Debug.sendTake_addition_quest = function(first_quest_id) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,93, {
    "first_quest_id":first_quest_id,
  });
};

PIModule.Debug.sendSet_mission_star_max = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,94, {
  });
};

PIModule.Debug.sendClique_bank_cd = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,95, {
  });
};

PIModule.Debug.sendReset_despair_land_battle_num = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,96, {
  });
};

PIModule.Debug.sendReset_clique_store_send_times = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,97, {
  });
};

PIModule.Debug.sendAdd_clique_store_donate = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,98, {
  });
};

PIModule.Debug.sendPass_all_despair_land_level = function(star, camp_type) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,99, {
    "star":star,
    "camp_type":camp_type,
  });
};

PIModule.Debug.sendDespair_land_dummy_boss_kill = function() {
  PIProtocol.getInstance().send_packet( ((arguments[0]===undefined) ? 'def' : arguments[0]), 99,100, {
  });
};

PIModule.Debug.sendAdd_taoyuan_item = function(item_id, number) {
  PIProtocol.getInstance().send_packet( ((arguments[2]===undefined) ? 'def' : arguments[2]), 99,101, {
    "item_id":item_id,
    "number":number,
  });
};

PIModule.Debug.sendAdd_taoyuan_exp = function(add_exp) {
  PIProtocol.getInstance().send_packet( ((arguments[1]===undefined) ? 'def' : arguments[1]), 99,102, {
    "add_exp":add_exp,
  });
};

