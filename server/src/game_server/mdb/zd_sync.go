package mdb

func (file *SyncFile) initDB() error {
	file.BeginTrans()
	file.WriteBytes([]byte("global_announcement,id,expire_time,tpl_id,parameters,content,send_time,spacing_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*expire_time - bigint*/, 3/*tpl_id - int*/, 7/*parameters - varchar*/, 7/*content - varchar*/, 4/*send_time - bigint*/, 4/*spacing_time - bigint*/})

	file.WriteBytes([]byte("global_arena_rank,rank,pid"))
	file.WriteBytes([]byte{3/*rank - int*/, 4/*pid - bigint*/})

	file.WriteBytes([]byte("global_clique,id,name,announce,total_donate_coins,owner_pid,owner_login_time,manger_pid1,manger_pid2,center_building_coins,temple_building_coins,health_building_coins,attack_building_coins,defense_building_coins,store_building_coins,bank_building_coins,center_building_level,temple_building_level,health_building_level,attack_building_level,defense_building_level,bank_building_level,member_json,auto_audit,auto_audit_level,contrib,contrib_clear_time,recruit_time,worship_time,worship_cnt,store_send_time,store_send_cnt"))
	file.WriteBytes([]byte{4/*id - bigint*/, 7/*name - varchar*/, 7/*announce - varchar*/, 4/*total_donate_coins - bigint*/, 4/*owner_pid - bigint*/, 4/*owner_login_time - bigint*/, 4/*manger_pid1 - bigint*/, 4/*manger_pid2 - bigint*/, 4/*center_building_coins - bigint*/, 4/*temple_building_coins - bigint*/, 4/*health_building_coins - bigint*/, 4/*attack_building_coins - bigint*/, 4/*defense_building_coins - bigint*/, 4/*store_building_coins - bigint*/, 4/*bank_building_coins - bigint*/, 2/*center_building_level - smallint*/, 2/*temple_building_level - smallint*/, 2/*health_building_level - smallint*/, 2/*attack_building_level - smallint*/, 2/*defense_building_level - smallint*/, 2/*bank_building_level - smallint*/, 8/*member_json - blob*/, 1/*auto_audit - tinyint*/, 2/*auto_audit_level - smallint*/, 4/*contrib - bigint*/, 4/*contrib_clear_time - bigint*/, 4/*recruit_time - bigint*/, 4/*worship_time - bigint*/, 1/*worship_cnt - tinyint*/, 4/*store_send_time - bigint*/, 1/*store_send_cnt - tinyint*/})

	file.WriteBytes([]byte("global_clique_boat,id,clique_id,boat_type,owner_pid,escort_start_timestamp,total_escort_time,recover_pid,hijacker_pid,hijack_timestamp,status"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*clique_id - bigint*/, 2/*boat_type - smallint*/, 4/*owner_pid - bigint*/, 4/*escort_start_timestamp - bigint*/, 2/*total_escort_time - smallint*/, 4/*recover_pid - bigint*/, 4/*hijacker_pid - bigint*/, 4/*hijack_timestamp - bigint*/, 1/*status - tinyint*/})

	file.WriteBytes([]byte("global_despair_boss,camp_type,level,timestamp,dead_timestamp,hp,max_hp"))
	file.WriteBytes([]byte{1/*camp_type - tinyint*/, 2/*level - smallint*/, 4/*timestamp - bigint*/, 4/*dead_timestamp - bigint*/, 4/*hp - bigint*/, 4/*max_hp - bigint*/})

	file.WriteBytes([]byte("global_despair_boss_history,id,version,camp_type,level,timestamp,start_timestamp,dead_timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*version - bigint*/, 1/*camp_type - tinyint*/, 2/*level - smallint*/, 4/*timestamp - bigint*/, 4/*start_timestamp - bigint*/, 4/*dead_timestamp - bigint*/})

	file.WriteBytes([]byte("global_despair_land,id,version,timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*version - bigint*/, 4/*timestamp - bigint*/})

	file.WriteBytes([]byte("global_despair_land_battle_record,id,pid,fightnum,timestamp,tag,battle_version,camp_type,level_id,record"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*fightnum - int*/, 4/*timestamp - bigint*/, 2/*tag - smallint*/, 2/*battle_version - smallint*/, 1/*camp_type - tinyint*/, 3/*level_id - int*/, 13/*record - mediumblob*/})

	file.WriteBytes([]byte("global_despair_land_camp,camp_type,timestamp,dead_timestamp,level"))
	file.WriteBytes([]byte{1/*camp_type - tinyint*/, 4/*timestamp - bigint*/, 4/*dead_timestamp - bigint*/, 2/*level - smallint*/})

	file.WriteBytes([]byte("global_gift_card_record,id,pid,code,timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 6/*code - char*/, 4/*timestamp - bigint*/})

	file.WriteBytes([]byte("global_group_buy_status,id,cid,status"))
	file.WriteBytes([]byte{4/*id - bigint*/, 3/*cid - int*/, 3/*status - int*/})

	file.WriteBytes([]byte("global_mail,id,send_time,expire_time,title,content,priority,min_level,max_level,min_vip_level,max_vip_level,min_create_time,max_create_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*send_time - bigint*/, 4/*expire_time - bigint*/, 7/*title - varchar*/, 7/*content - varchar*/, 1/*priority - tinyint*/, 2/*min_level - smallint*/, 2/*max_level - smallint*/, 2/*min_vip_level - smallint*/, 2/*max_vip_level - smallint*/, 4/*min_create_time - bigint*/, 4/*max_create_time - bigint*/})

	file.WriteBytes([]byte("global_mail_attachments,id,global_mail_id,item_id,attachment_type,item_num"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*global_mail_id - bigint*/, 2/*item_id - smallint*/, 1/*attachment_type - tinyint*/, 4/*item_num - bigint*/})

	file.WriteBytes([]byte("global_tb_xxd_onlinecnt,id,gameappid,timekey,gsid,onlinecntios,onlinecntandroid,cid"))
	file.WriteBytes([]byte{4/*id - bigint*/, 7/*gameappid - varchar*/, 4/*timekey - bigint*/, 4/*gsid - bigint*/, 4/*onlinecntios - bigint*/, 4/*onlinecntandroid - bigint*/, 4/*cid - bigint*/})

	file.WriteBytes([]byte("player,id,user,nick,main_role_id,cid"))
	file.WriteBytes([]byte{4/*id - bigint*/, 7/*user - varchar*/, 7/*nick - varchar*/, 4/*main_role_id - bigint*/, 4/*cid - bigint*/})

	file.WriteBytes([]byte("player_activity,pid,activity,last_update"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*activity - int*/, 4/*last_update - bigint*/})

	file.WriteBytes([]byte("player_addition_quest,id,pid,serial_number,quest_id,lock,progress,state"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*serial_number - int*/, 3/*quest_id - int*/, 2/*lock - smallint*/, 2/*progress - smallint*/, 1/*state - tinyint*/})

	file.WriteBytes([]byte("player_any_sdk_order,id,pid,order_id,present"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*order_id - bigint*/, 4/*present - bigint*/})

	file.WriteBytes([]byte("player_arena,pid,daily_num,failed_cd_time,battle_time,win_times,daily_award_item,new_record_count,daily_fame,buy_times"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*daily_num - smallint*/, 4/*failed_cd_time - bigint*/, 4/*battle_time - bigint*/, 2/*win_times - smallint*/, 3/*daily_award_item - int*/, 2/*new_record_count - smallint*/, 3/*daily_fame - int*/, 2/*buy_times - smallint*/})

	file.WriteBytes([]byte("player_arena_rank,pid,rank,rank1,rank2,rank3,time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*rank - int*/, 3/*rank1 - int*/, 3/*rank2 - int*/, 3/*rank3 - int*/, 4/*time - bigint*/})

	file.WriteBytes([]byte("player_arena_record,id,pid,mode,old_rank,new_rank,fight_num,target_pid,target_nick,target_old_rank,target_new_rank,target_fight_num,record_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*mode - tinyint*/, 3/*old_rank - int*/, 3/*new_rank - int*/, 3/*fight_num - int*/, 4/*target_pid - bigint*/, 7/*target_nick - varchar*/, 3/*target_old_rank - int*/, 3/*target_new_rank - int*/, 3/*target_fight_num - int*/, 4/*record_time - bigint*/})

	file.WriteBytes([]byte("player_awaken_graphic,id,pid,role_id,attr_impl,level"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 2/*attr_impl - smallint*/, 1/*level - tinyint*/})

	file.WriteBytes([]byte("player_battle_pet,id,pid,battle_pet_id,level,exp,skill_level"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*battle_pet_id - int*/, 3/*level - int*/, 4/*exp - bigint*/, 3/*skill_level - int*/})

	file.WriteBytes([]byte("player_battle_pet_grid,id,pid,grid_id,battle_pet_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*grid_id - tinyint*/, 3/*battle_pet_id - int*/})

	file.WriteBytes([]byte("player_battle_pet_state,pid,upgrade_by_ingot_num,upgrade_by_ingot_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*upgrade_by_ingot_num - int*/, 4/*upgrade_by_ingot_time - bigint*/})

	file.WriteBytes([]byte("player_chest_state,pid,free_coin_chest_num,last_free_coin_chest_at,coin_chest_num,coin_chest_ten_num,is_first_coin_ten,coin_chest_consume,last_coin_chest_at,last_free_ingot_chest_at,ingot_chest_num,ingot_chest_ten_num,is_first_ingot_ten,ingot_chest_consume,last_ingot_chest_at,last_free_pet_chest_at"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*free_coin_chest_num - int*/, 4/*last_free_coin_chest_at - bigint*/, 3/*coin_chest_num - int*/, 3/*coin_chest_ten_num - int*/, 1/*is_first_coin_ten - tinyint*/, 4/*coin_chest_consume - bigint*/, 4/*last_coin_chest_at - bigint*/, 4/*last_free_ingot_chest_at - bigint*/, 3/*ingot_chest_num - int*/, 3/*ingot_chest_ten_num - int*/, 1/*is_first_ingot_ten - tinyint*/, 4/*ingot_chest_consume - bigint*/, 4/*last_ingot_chest_at - bigint*/, 4/*last_free_pet_chest_at - bigint*/})

	file.WriteBytes([]byte("player_clique_kongfu_attrib,pid,health,attack,defence"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*health - int*/, 3/*attack - int*/, 3/*defence - int*/})

	file.WriteBytes([]byte("player_coins,pid,buy_time,daily_count,batch_bought"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*buy_time - bigint*/, 2/*daily_count - smallint*/, 2/*batch_bought - smallint*/})

	file.WriteBytes([]byte("player_cornucopia,pid,open_time,daily_count"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*open_time - bigint*/, 2/*daily_count - smallint*/})

	file.WriteBytes([]byte("player_daily_quest,id,pid,quest_id,finish_count,last_update_time,award_status,class"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*quest_id - smallint*/, 2/*finish_count - smallint*/, 4/*last_update_time - bigint*/, 1/*award_status - tinyint*/, 2/*class - smallint*/})

	file.WriteBytes([]byte("player_daily_quest_star_award_info,pid,stars,lastupdatetime,awarded"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*stars - int*/, 4/*lastupdatetime - bigint*/, 7/*awarded - varchar*/})

	file.WriteBytes([]byte("player_daily_sign_in_record,id,pid,sign_in_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*sign_in_time - bigint*/})

	file.WriteBytes([]byte("player_daily_sign_in_state,pid,update_time,record,signed_today"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*update_time - bigint*/, 2/*record - smallint*/, 1/*signed_today - tinyint*/})

	file.WriteBytes([]byte("player_despair_land_camp_state,id,pid,camp_type,timestamp,point,kill_num,dead_num,dead_num_boss,hurt,boss_battle_num,daily_boss_battle_num,boss_battle_timestamp,awarded"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*camp_type - tinyint*/, 4/*timestamp - bigint*/, 4/*point - bigint*/, 4/*kill_num - bigint*/, 4/*dead_num - bigint*/, 4/*dead_num_boss - bigint*/, 4/*hurt - bigint*/, 3/*boss_battle_num - int*/, 3/*daily_boss_battle_num - int*/, 4/*boss_battle_timestamp - bigint*/, 1/*awarded - tinyint*/})

	file.WriteBytes([]byte("player_despair_land_camp_state_history,id,pid,camp_type,timestamp,point,kill_num,dead_num,dead_num_boss,hurt,boss_battle_num,daily_boss_battle_num,boss_battle_timestamp,awarded"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*camp_type - tinyint*/, 4/*timestamp - bigint*/, 4/*point - bigint*/, 4/*kill_num - bigint*/, 4/*dead_num - bigint*/, 4/*dead_num_boss - bigint*/, 4/*hurt - bigint*/, 3/*boss_battle_num - int*/, 3/*daily_boss_battle_num - int*/, 4/*boss_battle_timestamp - bigint*/, 1/*awarded - tinyint*/})

	file.WriteBytes([]byte("player_despair_land_level_record,id,pid,camp_type,timestamp,level_id,round,passed_timestamp,first_passed_timestamp,first_passed_fightnum,passed_award,perfect_award,milestone_award"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*camp_type - tinyint*/, 4/*timestamp - bigint*/, 3/*level_id - int*/, 1/*round - tinyint*/, 4/*passed_timestamp - bigint*/, 4/*first_passed_timestamp - bigint*/, 3/*first_passed_fightnum - int*/, 1/*passed_award - tinyint*/, 1/*perfect_award - tinyint*/, 1/*milestone_award - tinyint*/})

	file.WriteBytes([]byte("player_despair_land_state,pid,point,kill_num,dead_num,daily_battle_num,daily_battle_timestamp,daily_bought_battle_num,daily_bought_timestamp,daily_boss_bought_timestamp,daily_boss_beast_bought_battle_num,daily_boss_walking_dead_bought_battle_num,daily_boss_devil_bought_battle_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*point - bigint*/, 4/*kill_num - bigint*/, 4/*dead_num - bigint*/, 2/*daily_battle_num - smallint*/, 4/*daily_battle_timestamp - bigint*/, 2/*daily_bought_battle_num - smallint*/, 4/*daily_bought_timestamp - bigint*/, 4/*daily_boss_bought_timestamp - bigint*/, 2/*daily_boss_beast_bought_battle_num - smallint*/, 2/*daily_boss_walking_dead_bought_battle_num - smallint*/, 2/*daily_boss_devil_bought_battle_num - smallint*/})

	file.WriteBytes([]byte("player_driving_sword_event,id,pid,cloud_id,x,y,event_type,data_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*cloud_id - smallint*/, 1/*x - tinyint*/, 1/*y - tinyint*/, 1/*event_type - tinyint*/, 1/*data_id - tinyint*/})

	file.WriteBytes([]byte("player_driving_sword_event_exploring,id,pid,status,garrison_start,garrison_time,award_time,role_id,cloud_id,x,y,data_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*status - tinyint*/, 4/*garrison_start - bigint*/, 4/*garrison_time - bigint*/, 4/*award_time - bigint*/, 1/*role_id - tinyint*/, 2/*cloud_id - smallint*/, 1/*x - tinyint*/, 1/*y - tinyint*/, 1/*data_id - tinyint*/})

	file.WriteBytes([]byte("player_driving_sword_event_treasure,id,pid,progress,cloud_id,x,y,data_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*progress - tinyint*/, 2/*cloud_id - smallint*/, 1/*x - tinyint*/, 1/*y - tinyint*/, 1/*data_id - tinyint*/})

	file.WriteBytes([]byte("player_driving_sword_event_visiting,id,pid,status,target_pid,target_side_state,cloud_id,x,y,data_id,target_status"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*status - tinyint*/, 4/*target_pid - bigint*/, 8/*target_side_state - blob*/, 2/*cloud_id - smallint*/, 1/*x - tinyint*/, 1/*y - tinyint*/, 1/*data_id - tinyint*/, 5/*target_status - text*/})

	file.WriteBytes([]byte("player_driving_sword_eventmask,id,pid,cloud_id,events"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*cloud_id - smallint*/, 8/*events - blob*/})

	file.WriteBytes([]byte("player_driving_sword_info,pid,current_cloud,highest_cloud,current_x,current_y,allowed_action,action_refresh_time,action_buy_time,daily_action_bought"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*current_cloud - smallint*/, 2/*highest_cloud - smallint*/, 1/*current_x - tinyint*/, 1/*current_y - tinyint*/, 2/*allowed_action - smallint*/, 4/*action_refresh_time - bigint*/, 4/*action_buy_time - bigint*/, 1/*daily_action_bought - tinyint*/})

	file.WriteBytes([]byte("player_driving_sword_map,id,pid,cloud_id,shadows,obstacle_count,exploring_count,visiting_count,treasure_count"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*cloud_id - smallint*/, 8/*shadows - blob*/, 1/*obstacle_count - tinyint*/, 1/*exploring_count - tinyint*/, 1/*visiting_count - tinyint*/, 1/*treasure_count - tinyint*/})

	file.WriteBytes([]byte("player_equipment,id,pid,role_id,weapon_id,clothes_id,accessories_id,shoe_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 4/*weapon_id - bigint*/, 4/*clothes_id - bigint*/, 4/*accessories_id - bigint*/, 4/*shoe_id - bigint*/})

	file.WriteBytes([]byte("player_event_award_record,pid,record_bytes,json_event_record"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 13/*record_bytes - mediumblob*/, 13/*json_event_record - mediumblob*/})

	file.WriteBytes([]byte("player_event_daily_online,pid,timestamp,total_online_time,awarded_onlinetime"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*timestamp - bigint*/, 3/*total_online_time - int*/, 3/*awarded_onlinetime - int*/})

	file.WriteBytes([]byte("player_event_vn_daily_recharge,id,pid,page,timestamp,awarded,daily_recharge,total_recharge"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*page - int*/, 4/*timestamp - bigint*/, 1/*awarded - tinyint*/, 4/*daily_recharge - bigint*/, 4/*total_recharge - bigint*/})

	file.WriteBytes([]byte("player_event_vn_dragon_ball_record,pid,item_id,createtime"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*item_id - smallint*/, 4/*createtime - bigint*/})

	file.WriteBytes([]byte("player_events_fight_power,pid,lock"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*lock - int*/})

	file.WriteBytes([]byte("player_events_ingot_record,pid,ingot_in,ingot_in_end_time,ingot_out,ingot_out_end_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*ingot_in - bigint*/, 4/*ingot_in_end_time - bigint*/, 4/*ingot_out - bigint*/, 4/*ingot_out_end_time - bigint*/})

	file.WriteBytes([]byte("player_extend_level,pid,coin_pass_time,exp_pass_time,ghost_pass_time,pet_pass_time,buddy_pass_time,coin_daily_num,exp_daily_num,buddy_daily_num,pet_daily_num,ghost_daily_num,rand_buddy_role_id,buddy_pos,buddy_tactical,role_pos,exp_maxlevel,coins_maxlevel,coins_buy_num,exp_buy_num,coins_buy_time,exp_buy_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*coin_pass_time - bigint*/, 4/*exp_pass_time - bigint*/, 4/*ghost_pass_time - bigint*/, 4/*pet_pass_time - bigint*/, 4/*buddy_pass_time - bigint*/, 1/*coin_daily_num - tinyint*/, 1/*exp_daily_num - tinyint*/, 1/*buddy_daily_num - tinyint*/, 1/*pet_daily_num - tinyint*/, 1/*ghost_daily_num - tinyint*/, 1/*rand_buddy_role_id - tinyint*/, 1/*buddy_pos - tinyint*/, 1/*buddy_tactical - tinyint*/, 1/*role_pos - tinyint*/, 2/*exp_maxlevel - smallint*/, 2/*coins_maxlevel - smallint*/, 2/*coins_buy_num - smallint*/, 2/*exp_buy_num - smallint*/, 4/*coins_buy_time - bigint*/, 4/*exp_buy_time - bigint*/})

	file.WriteBytes([]byte("player_extend_quest,id,pid,quest_id,progress,state"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*quest_id - int*/, 2/*progress - smallint*/, 1/*state - tinyint*/})

	file.WriteBytes([]byte("player_fame,pid,fame,level,arena_fame,mult_level_fame"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*fame - int*/, 2/*level - smallint*/, 3/*arena_fame - int*/, 3/*mult_level_fame - int*/})

	file.WriteBytes([]byte("player_fashion,id,pid,fashion_id,expire_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*fashion_id - smallint*/, 4/*expire_time - bigint*/})

	file.WriteBytes([]byte("player_fashion_state,pid,update_time,dressed_fashion_id"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*update_time - bigint*/, 2/*dressed_fashion_id - smallint*/})

	file.WriteBytes([]byte("player_fate_box_state,pid,lock,star_box_free_draw_timestamp,star_box_draw_count,moon_box_free_draw_timestamp,moon_box_draw_count,sun_box_free_draw_timestamp,sun_box_draw_count,hunyuan_box_free_draw_timestamp,hunyuan_box_draw_count"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*lock - int*/, 4/*star_box_free_draw_timestamp - bigint*/, 3/*star_box_draw_count - int*/, 4/*moon_box_free_draw_timestamp - bigint*/, 3/*moon_box_draw_count - int*/, 4/*sun_box_free_draw_timestamp - bigint*/, 3/*sun_box_draw_count - int*/, 4/*hunyuan_box_free_draw_timestamp - bigint*/, 3/*hunyuan_box_draw_count - int*/})

	file.WriteBytes([]byte("player_fight_num,pid,fight_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*fight_num - int*/})

	file.WriteBytes([]byte("player_first_charge,id,pid,product_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 7/*product_id - varchar*/})

	file.WriteBytes([]byte("player_formation,pid,pos0,pos1,pos2,pos3,pos4,pos5,pos6,pos7,pos8"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 1/*pos0 - tinyint*/, 1/*pos1 - tinyint*/, 1/*pos2 - tinyint*/, 1/*pos3 - tinyint*/, 1/*pos4 - tinyint*/, 1/*pos5 - tinyint*/, 1/*pos6 - tinyint*/, 1/*pos7 - tinyint*/, 1/*pos8 - tinyint*/})

	file.WriteBytes([]byte("player_func_key,pid,key,played_key,unique_key"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*key - smallint*/, 2/*played_key - smallint*/, 4/*unique_key - bigint*/})

	file.WriteBytes([]byte("player_ghost,id,pid,ghost_id,star,level,exp,pos,is_new,role_id,skill_level,relation_id,add_growth"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*ghost_id - smallint*/, 1/*star - tinyint*/, 2/*level - smallint*/, 4/*exp - bigint*/, 2/*pos - smallint*/, 1/*is_new - tinyint*/, 1/*role_id - tinyint*/, 2/*skill_level - smallint*/, 4/*relation_id - bigint*/, 2/*add_growth - smallint*/})

	file.WriteBytes([]byte("player_ghost_equipment,id,pid,role_id,ghost_power,pos1,pos2,pos3,pos4"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 3/*ghost_power - int*/, 4/*pos1 - bigint*/, 4/*pos2 - bigint*/, 4/*pos3 - bigint*/, 4/*pos4 - bigint*/})

	file.WriteBytes([]byte("player_ghost_state,pid,train_by_ingot_num,train_by_ingot_time,last_flush_time,ghost_fight_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*train_by_ingot_num - int*/, 4/*train_by_ingot_time - bigint*/, 4/*last_flush_time - bigint*/, 4/*ghost_fight_num - bigint*/})

	file.WriteBytes([]byte("player_global_chat_state,pid,ban_until"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*ban_until - bigint*/})

	file.WriteBytes([]byte("player_global_clique_building,pid,silver_exchange_time,gold_exchange_num,silver_exchange_num,donate_coins_center_building,donate_coins_temple_building,donate_coins_bank_building,donate_coins_health_building,donate_coins_attack_building,donate_coins_defense_building,donate_coins_store_building,health_level,attack_level,defense_level,worship_time,donate_coins_center_building_time,glod_exchange_time,worship_type"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*silver_exchange_time - bigint*/, 2/*gold_exchange_num - smallint*/, 2/*silver_exchange_num - smallint*/, 4/*donate_coins_center_building - bigint*/, 4/*donate_coins_temple_building - bigint*/, 4/*donate_coins_bank_building - bigint*/, 4/*donate_coins_health_building - bigint*/, 4/*donate_coins_attack_building - bigint*/, 4/*donate_coins_defense_building - bigint*/, 4/*donate_coins_store_building - bigint*/, 2/*health_level - smallint*/, 2/*attack_level - smallint*/, 2/*defense_level - smallint*/, 4/*worship_time - bigint*/, 4/*donate_coins_center_building_time - bigint*/, 4/*glod_exchange_time - bigint*/, 1/*worship_type - tinyint*/})

	file.WriteBytes([]byte("player_global_clique_building_quest,id,pid,quest_id,award_status,building_type"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*quest_id - smallint*/, 1/*award_status - tinyint*/, 2/*building_type - smallint*/})

	file.WriteBytes([]byte("player_global_clique_daily_quest,id,pid,quest_id,finish_count,last_update_time,award_status,class"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*quest_id - smallint*/, 2/*finish_count - smallint*/, 4/*last_update_time - bigint*/, 1/*award_status - tinyint*/, 2/*class - smallint*/})

	file.WriteBytes([]byte("player_global_clique_escort,pid,daily_escort_timestamp,daily_escort_num,escort_boat_type,status,hijacked"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*daily_escort_timestamp - bigint*/, 2/*daily_escort_num - smallint*/, 2/*escort_boat_type - smallint*/, 1/*status - tinyint*/, 1/*hijacked - tinyint*/})

	file.WriteBytes([]byte("player_global_clique_escort_message,id,pid,tpl_id,parameters,timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*tpl_id - smallint*/, 7/*parameters - varchar*/, 4/*timestamp - bigint*/})

	file.WriteBytes([]byte("player_global_clique_hijack,pid,daily_hijack_timestamp,daily_hijack_num,hijacked_boat_type,status,buy_timestamp,buy_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*daily_hijack_timestamp - bigint*/, 2/*daily_hijack_num - smallint*/, 2/*hijacked_boat_type - smallint*/, 1/*status - tinyint*/, 4/*buy_timestamp - bigint*/, 2/*buy_num - smallint*/})

	file.WriteBytes([]byte("player_global_clique_info,pid,clique_id,join_time,contrib,contrib_clear_time,donate_coins_time,daily_donate_coins,total_contrib"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*clique_id - bigint*/, 4/*join_time - bigint*/, 4/*contrib - bigint*/, 4/*contrib_clear_time - bigint*/, 4/*donate_coins_time - bigint*/, 4/*daily_donate_coins - bigint*/, 4/*total_contrib - bigint*/})

	file.WriteBytes([]byte("player_global_clique_kongfu,id,pid,kongfu_id,level"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*kongfu_id - int*/, 2/*level - smallint*/})

	file.WriteBytes([]byte("player_global_friend,id,pid,friend_pid,friend_nick,friend_role_id,friend_mode,last_chat_time,friend_time,send_heart_time,block_mode"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*friend_pid - bigint*/, 7/*friend_nick - varchar*/, 1/*friend_role_id - tinyint*/, 1/*friend_mode - tinyint*/, 4/*last_chat_time - bigint*/, 4/*friend_time - bigint*/, 4/*send_heart_time - bigint*/, 1/*block_mode - tinyint*/})

	file.WriteBytes([]byte("player_global_friend_chat,id,pid,friend_pid,mode,send_time,message"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*friend_pid - bigint*/, 1/*mode - tinyint*/, 4/*send_time - bigint*/, 7/*message - varchar*/})

	file.WriteBytes([]byte("player_global_friend_state,pid,delete_day_count,delete_time,exist_offline_chat,platform_friend_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*delete_day_count - int*/, 4/*delete_time - bigint*/, 1/*exist_offline_chat - tinyint*/, 3/*platform_friend_num - int*/})

	file.WriteBytes([]byte("player_global_gift_card_record,id,pid,code"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 6/*code - char*/})

	file.WriteBytes([]byte("player_global_mail_state,pid,max_timestamp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*max_timestamp - bigint*/})

	file.WriteBytes([]byte("player_global_world_chat_state,pid,timestamp,daily_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*timestamp - bigint*/, 2/*daily_num - smallint*/})

	file.WriteBytes([]byte("player_hard_level,pid,lock,award_lock"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*lock - int*/, 3/*award_lock - int*/})

	file.WriteBytes([]byte("player_hard_level_record,id,pid,level_id,open_time,score,round,daily_num,last_enter_time,buy_times,buy_update_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*level_id - int*/, 4/*open_time - bigint*/, 3/*score - int*/, 1/*round - tinyint*/, 1/*daily_num - tinyint*/, 4/*last_enter_time - bigint*/, 2/*buy_times - smallint*/, 4/*buy_update_time - bigint*/})

	file.WriteBytes([]byte("player_heart,pid,value,update_time,add_day_count,add_time,recover_day_count"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*value - smallint*/, 4/*update_time - bigint*/, 3/*add_day_count - int*/, 4/*add_time - bigint*/, 2/*recover_day_count - smallint*/})

	file.WriteBytes([]byte("player_heart_draw,id,pid,draw_type,daily_num,draw_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*draw_type - tinyint*/, 1/*daily_num - tinyint*/, 4/*draw_time - bigint*/})

	file.WriteBytes([]byte("player_heart_draw_card_record,id,pid,award_type,award_num,item_id,draw_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*award_type - tinyint*/, 2/*award_num - smallint*/, 2/*item_id - smallint*/, 4/*draw_time - bigint*/})

	file.WriteBytes([]byte("player_heart_draw_wheel_record,id,pid,award_type,award_num,item_id,draw_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*award_type - tinyint*/, 2/*award_num - smallint*/, 2/*item_id - smallint*/, 4/*draw_time - bigint*/})

	file.WriteBytes([]byte("player_info,pid,ingot,coins,new_mail_num,last_login_time,last_offline_time,total_online_time,first_login_time,new_arena_report_num,last_skill_flush,sword_soul_fragment"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*ingot - bigint*/, 4/*coins - bigint*/, 2/*new_mail_num - smallint*/, 4/*last_login_time - bigint*/, 4/*last_offline_time - bigint*/, 4/*total_online_time - bigint*/, 4/*first_login_time - bigint*/, 2/*new_arena_report_num - smallint*/, 4/*last_skill_flush - bigint*/, 4/*sword_soul_fragment - bigint*/})

	file.WriteBytes([]byte("player_is_operated,pid,operate_value"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*operate_value - bigint*/})

	file.WriteBytes([]byte("player_item,id,pid,item_id,num,is_dressed,buy_back_state,appendix_id,refine_level_bak,refine_fail_times,refine_level,price"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*item_id - smallint*/, 2/*num - smallint*/, 1/*is_dressed - tinyint*/, 1/*buy_back_state - tinyint*/, 4/*appendix_id - bigint*/, 2/*refine_level_bak - smallint*/, 2/*refine_fail_times - smallint*/, 2/*refine_level - smallint*/, 3/*price - int*/})

	file.WriteBytes([]byte("player_item_appendix,id,pid,health,cultivation,speed,attack,defence,dodge_level,hit_level,block_level,critical_level,tenacity_level,destroy_level,recast_attr"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*health - int*/, 3/*cultivation - int*/, 3/*speed - int*/, 3/*attack - int*/, 3/*defence - int*/, 3/*dodge_level - int*/, 3/*hit_level - int*/, 3/*block_level - int*/, 3/*critical_level - int*/, 3/*tenacity_level - int*/, 3/*destroy_level - int*/, 1/*recast_attr - tinyint*/})

	file.WriteBytes([]byte("player_item_buyback,pid,back_id1,back_id2,back_id3,back_id4,back_id5,back_id6"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*back_id1 - bigint*/, 4/*back_id2 - bigint*/, 4/*back_id3 - bigint*/, 4/*back_id4 - bigint*/, 4/*back_id5 - bigint*/, 4/*back_id6 - bigint*/})

	file.WriteBytes([]byte("player_item_recast_state,pid,player_item_id,selected_attr,attr1_type,attr1_value,attr2_type,attr2_value,attr3_type,attr3_value"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*player_item_id - bigint*/, 1/*selected_attr - tinyint*/, 1/*attr1_type - tinyint*/, 3/*attr1_value - int*/, 1/*attr2_type - tinyint*/, 3/*attr2_value - int*/, 1/*attr3_type - tinyint*/, 3/*attr3_value - int*/})

	file.WriteBytes([]byte("player_level_award_info,pid,awarded"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 7/*awarded - varchar*/})

	file.WriteBytes([]byte("player_login_award_record,pid,active_days,record,update_timestamp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*active_days - int*/, 3/*record - int*/, 4/*update_timestamp - bigint*/})

	file.WriteBytes([]byte("player_mail,id,pid,mail_id,state,send_time,parameters,have_attachment,title,content,expire_time,priority"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*mail_id - int*/, 1/*state - tinyint*/, 4/*send_time - bigint*/, 7/*parameters - varchar*/, 1/*have_attachment - tinyint*/, 7/*title - varchar*/, 7/*content - varchar*/, 4/*expire_time - bigint*/, 1/*priority - tinyint*/})

	file.WriteBytes([]byte("player_mail_attachment,id,pid,player_mail_id,attachment_type,item_id,item_num"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*player_mail_id - bigint*/, 1/*attachment_type - tinyint*/, 2/*item_id - smallint*/, 4/*item_num - bigint*/})

	file.WriteBytes([]byte("player_mail_attachment_lg,id,pid,player_mail_id,attachment_type,item_id,item_num,take_timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*player_mail_id - bigint*/, 1/*attachment_type - tinyint*/, 2/*item_id - smallint*/, 4/*item_num - bigint*/, 4/*take_timestamp - bigint*/})

	file.WriteBytes([]byte("player_mail_lg,id,pmid,pid,mail_id,state,send_time,parameters,have_attachment,title,content"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pmid - bigint*/, 4/*pid - bigint*/, 3/*mail_id - int*/, 1/*state - tinyint*/, 4/*send_time - bigint*/, 7/*parameters - varchar*/, 1/*have_attachment - tinyint*/, 7/*title - varchar*/, 7/*content - varchar*/})

	file.WriteBytes([]byte("player_meditation_state,pid,accumulate_time,start_timestamp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*accumulate_time - int*/, 4/*start_timestamp - bigint*/})

	file.WriteBytes([]byte("player_mission,pid,key,max_order"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*key - int*/, 1/*max_order - tinyint*/})

	file.WriteBytes([]byte("player_mission_level,pid,lock,max_lock,award_lock"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*lock - int*/, 3/*max_lock - int*/, 3/*award_lock - int*/})

	file.WriteBytes([]byte("player_mission_level_record,id,pid,mission_id,mission_level_id,open_time,score,round,daily_num,last_enter_time,empty_shadow_bits"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*mission_id - smallint*/, 3/*mission_level_id - int*/, 4/*open_time - bigint*/, 3/*score - int*/, 1/*round - tinyint*/, 1/*daily_num - tinyint*/, 4/*last_enter_time - bigint*/, 2/*empty_shadow_bits - smallint*/})

	file.WriteBytes([]byte("player_mission_level_state_bin,pid,bin"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 8/*bin - blob*/})

	file.WriteBytes([]byte("player_mission_record,id,pid,town_id,mission_id,open_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*town_id - smallint*/, 2/*mission_id - smallint*/, 4/*open_time - bigint*/})

	file.WriteBytes([]byte("player_mission_star_award,id,pid,town_id,box_type"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*town_id - smallint*/, 1/*box_type - tinyint*/})

	file.WriteBytes([]byte("player_money_tree,pid,total,waved_times_total,waved_times,last_waved_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*total - int*/, 1/*waved_times_total - tinyint*/, 1/*waved_times - tinyint*/, 4/*last_waved_time - bigint*/})

	file.WriteBytes([]byte("player_month_card_award_record,pid,last_update"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*last_update - bigint*/})

	file.WriteBytes([]byte("player_month_card_info,pid,starttime,endtime,buytimes,presenttotal"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*starttime - bigint*/, 4/*endtime - bigint*/, 4/*buytimes - bigint*/, 4/*presenttotal - bigint*/})

	file.WriteBytes([]byte("player_multi_level_info,pid,buddy_role_id,buddy_row,tactical_grid,daily_num,battle_time,lock"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 1/*buddy_role_id - tinyint*/, 1/*buddy_row - tinyint*/, 1/*tactical_grid - tinyint*/, 1/*daily_num - tinyint*/, 4/*battle_time - bigint*/, 3/*lock - int*/})

	file.WriteBytes([]byte("player_new_year_consume_record,pid,consume_record"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 5/*consume_record - text*/})

	file.WriteBytes([]byte("player_npc_talk_record,id,pid,npc_id,town_id,quest_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 3/*npc_id - int*/, 2/*town_id - smallint*/, 2/*quest_id - smallint*/})

	file.WriteBytes([]byte("player_opened_town_treasure,id,pid,town_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*town_id - smallint*/})

	file.WriteBytes([]byte("player_physical,pid,value,update_time,buy_count,buy_update_time,daily_count"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*value - smallint*/, 4/*update_time - bigint*/, 4/*buy_count - bigint*/, 4/*buy_update_time - bigint*/, 2/*daily_count - smallint*/})

	file.WriteBytes([]byte("player_purchase_record,id,pid,item_id,num,timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*item_id - smallint*/, 2/*num - smallint*/, 4/*timestamp - bigint*/})

	file.WriteBytes([]byte("player_push_notify_switch,pid,options"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*options - bigint*/})

	file.WriteBytes([]byte("player_pve_state,pid,max_passed_floor,max_awarded_floor,unpassed_floor_enemy_num,enter_time,daily_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*max_passed_floor - smallint*/, 2/*max_awarded_floor - smallint*/, 2/*unpassed_floor_enemy_num - smallint*/, 4/*enter_time - bigint*/, 1/*daily_num - tinyint*/})

	file.WriteBytes([]byte("player_qq_vip_gift,pid,qqvip,surper"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*qqvip - smallint*/, 2/*surper - smallint*/})

	file.WriteBytes([]byte("player_quest,pid,quest_id,state"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*quest_id - smallint*/, 1/*state - tinyint*/})

	file.WriteBytes([]byte("player_rainbow_level,pid,reset_timestamp,reset_num,segment,order,status,best_segment,best_order,best_record_timestamp,max_open_segment,max_pass_segment,auto_fight_num,buy_times,auto_fight_time,buy_timestamp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*reset_timestamp - bigint*/, 3/*reset_num - int*/, 2/*segment - smallint*/, 1/*order - tinyint*/, 1/*status - tinyint*/, 2/*best_segment - smallint*/, 1/*best_order - tinyint*/, 4/*best_record_timestamp - bigint*/, 2/*max_open_segment - smallint*/, 2/*max_pass_segment - smallint*/, 1/*auto_fight_num - tinyint*/, 2/*buy_times - smallint*/, 4/*auto_fight_time - bigint*/, 4/*buy_timestamp - bigint*/})

	file.WriteBytes([]byte("player_rainbow_level_state_bin,pid,bin"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 13/*bin - mediumblob*/})

	file.WriteBytes([]byte("player_role,id,pid,role_id,level,exp,friendship_level,status,coop_id,timestamp,fight_num"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 2/*level - smallint*/, 4/*exp - bigint*/, 3/*friendship_level - int*/, 2/*status - smallint*/, 2/*coop_id - smallint*/, 4/*timestamp - bigint*/, 3/*fight_num - int*/})

	file.WriteBytes([]byte("player_sealedbook,pid,sealed_record"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 13/*sealed_record - mediumblob*/})

	file.WriteBytes([]byte("player_send_heart_record,id,pid,friend_pid,send_heart_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*friend_pid - bigint*/, 4/*send_heart_time - bigint*/})

	file.WriteBytes([]byte("player_skill,id,pid,role_id,skill_id,skill_trnlv"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 2/*skill_id - smallint*/, 3/*skill_trnlv - int*/})

	file.WriteBytes([]byte("player_state,pid,ban_start_time,ban_end_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*ban_start_time - bigint*/, 4/*ban_end_time - bigint*/})

	file.WriteBytes([]byte("player_sword_soul,id,pid,pos,sword_soul_id,exp,level"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*pos - tinyint*/, 2/*sword_soul_id - smallint*/, 3/*exp - int*/, 1/*level - tinyint*/})

	file.WriteBytes([]byte("player_sword_soul_equipment,id,pid,role_id,is_equipped_xuanyuan,type_bits,pos1,pos2,pos3,pos4,pos5,pos6,pos7,pos8,pos9"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 1/*is_equipped_xuanyuan - tinyint*/, 4/*type_bits - bigint*/, 4/*pos1 - bigint*/, 4/*pos2 - bigint*/, 4/*pos3 - bigint*/, 4/*pos4 - bigint*/, 4/*pos5 - bigint*/, 4/*pos6 - bigint*/, 4/*pos7 - bigint*/, 4/*pos8 - bigint*/, 4/*pos9 - bigint*/})

	file.WriteBytes([]byte("player_sword_soul_state,pid,box_state,num,update_time,ingot_num,supersoul_additional_possible,ingot_yello_one,last_ingot_draw_time,sword_soul_num"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 1/*box_state - tinyint*/, 2/*num - smallint*/, 4/*update_time - bigint*/, 4/*ingot_num - bigint*/, 1/*supersoul_additional_possible - tinyint*/, 2/*ingot_yello_one - smallint*/, 4/*last_ingot_draw_time - bigint*/, 3/*sword_soul_num - int*/})

	file.WriteBytes([]byte("player_taoyuan_bless,pid,daily_bless_times,last_bless_time,daily_be_bless_times,last_be_bless_time,bless_pid1,bless_pid2,bless_pid3,bless_pid4,bless_pid5"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*daily_bless_times - int*/, 4/*last_bless_time - bigint*/, 3/*daily_be_bless_times - int*/, 4/*last_be_bless_time - bigint*/, 4/*bless_pid1 - bigint*/, 4/*bless_pid2 - bigint*/, 4/*bless_pid3 - bigint*/, 4/*bless_pid4 - bigint*/, 4/*bless_pid5 - bigint*/})

	file.WriteBytes([]byte("player_taoyuan_fileds,id,pid,filed_id,filed_status,plant_id,grow_time"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*filed_id - smallint*/, 2/*filed_status - smallint*/, 2/*plant_id - smallint*/, 4/*grow_time - bigint*/})

	file.WriteBytes([]byte("player_taoyuan_heart,pid,level,exp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*level - smallint*/, 4/*exp - bigint*/})

	file.WriteBytes([]byte("player_taoyuan_item,id,pid,item_id,num"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*item_id - smallint*/, 2/*num - smallint*/})

	file.WriteBytes([]byte("player_taoyuan_message,id,pid,nick,exp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 7/*nick - varchar*/, 3/*exp - int*/})

	file.WriteBytes([]byte("player_taoyuan_product,pid,skill_id1,skill_id2,same_time_wine_queue,make_wine_times,same_time_food_queue,make_food_times,food_queue1,food_queue1_start_timestamp,food_queue1_end_timestamp,food_queue2,food_queue2_start_timestamp,food_queue2_end_timestamp,food_queue3,food_queue3_start_timestamp,food_queue3_end_timestamp,food_queue4,food_queue4_start_timestamp,food_queue4_end_timestamp,wine_queue1,wine_queue1_start_timestamp,wine_queue1_end_timestamp,wine_queue2,wine_queue2_start_timestamp,wine_queue2_end_timestamp,wine_queue3,wine_queue3_start_timestamp,wine_queue3_end_timestamp,wine_queue4,wine_queue4_start_timestamp,wine_queue4_end_timestamp,is_open_wine,is_open_food"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*skill_id1 - smallint*/, 2/*skill_id2 - smallint*/, 2/*same_time_wine_queue - smallint*/, 4/*make_wine_times - bigint*/, 2/*same_time_food_queue - smallint*/, 4/*make_food_times - bigint*/, 2/*food_queue1 - smallint*/, 4/*food_queue1_start_timestamp - bigint*/, 4/*food_queue1_end_timestamp - bigint*/, 2/*food_queue2 - smallint*/, 4/*food_queue2_start_timestamp - bigint*/, 4/*food_queue2_end_timestamp - bigint*/, 2/*food_queue3 - smallint*/, 4/*food_queue3_start_timestamp - bigint*/, 4/*food_queue3_end_timestamp - bigint*/, 2/*food_queue4 - smallint*/, 4/*food_queue4_start_timestamp - bigint*/, 4/*food_queue4_end_timestamp - bigint*/, 2/*wine_queue1 - smallint*/, 4/*wine_queue1_start_timestamp - bigint*/, 4/*wine_queue1_end_timestamp - bigint*/, 2/*wine_queue2 - smallint*/, 4/*wine_queue2_start_timestamp - bigint*/, 4/*wine_queue2_end_timestamp - bigint*/, 2/*wine_queue3 - smallint*/, 4/*wine_queue3_start_timestamp - bigint*/, 4/*wine_queue3_end_timestamp - bigint*/, 2/*wine_queue4 - smallint*/, 4/*wine_queue4_start_timestamp - bigint*/, 4/*wine_queue4_end_timestamp - bigint*/, 1/*is_open_wine - tinyint*/, 1/*is_open_food - tinyint*/})

	file.WriteBytes([]byte("player_taoyuan_purchase_record,id,pid,item_id,num,timestamp"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*item_id - smallint*/, 2/*num - smallint*/, 4/*timestamp - bigint*/})

	file.WriteBytes([]byte("player_taoyuan_quest,pid,last_quest_flash_time,quest1_item_id,quest1_item_num,quest1_total_exp,quest1_total_coins,quest1_finish_time,quest2_item_id,quest2_item_num,quest2_total_exp,quest2_total_coins,quest2_finish_time,quest3_item_id,quest3_item_num,quest3_total_exp,quest3_total_coins,quest3_finish_time,quest4_item_id,quest4_item_num,quest4_total_exp,quest4_total_coins,quest4_finish_time,quest5_item_id,quest5_item_num,quest5_total_exp,quest5_total_coins,quest5_finish_time"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*last_quest_flash_time - bigint*/, 2/*quest1_item_id - smallint*/, 2/*quest1_item_num - smallint*/, 2/*quest1_total_exp - smallint*/, 4/*quest1_total_coins - bigint*/, 4/*quest1_finish_time - bigint*/, 2/*quest2_item_id - smallint*/, 2/*quest2_item_num - smallint*/, 2/*quest2_total_exp - smallint*/, 4/*quest2_total_coins - bigint*/, 4/*quest2_finish_time - bigint*/, 2/*quest3_item_id - smallint*/, 2/*quest3_item_num - smallint*/, 2/*quest3_total_exp - smallint*/, 4/*quest3_total_coins - bigint*/, 4/*quest3_finish_time - bigint*/, 2/*quest4_item_id - smallint*/, 2/*quest4_item_num - smallint*/, 2/*quest4_total_exp - smallint*/, 4/*quest4_total_coins - bigint*/, 4/*quest4_finish_time - bigint*/, 2/*quest5_item_id - smallint*/, 2/*quest5_item_num - smallint*/, 2/*quest5_total_exp - smallint*/, 4/*quest5_total_coins - bigint*/, 4/*quest5_finish_time - bigint*/})

	file.WriteBytes([]byte("player_tb_xxd_roleinfo,pid,gameappid,openid,regtime,level,iFriends,moneyios,moneyandroid,diamondios,diamondandroid"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 7/*gameappid - varchar*/, 7/*openid - varchar*/, 4/*regtime - bigint*/, 2/*level - smallint*/, 2/*iFriends - smallint*/, 4/*moneyios - bigint*/, 4/*moneyandroid - bigint*/, 4/*diamondios - bigint*/, 4/*diamondandroid - bigint*/})

	file.WriteBytes([]byte("player_team_info,pid,relationship,health_lv,attack_lv,defence_lv"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 3/*relationship - int*/, 2/*health_lv - smallint*/, 2/*attack_lv - smallint*/, 2/*defence_lv - smallint*/})

	file.WriteBytes([]byte("player_totem,id,pid,totem_id,level,skill_id"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*totem_id - smallint*/, 1/*level - tinyint*/, 2/*skill_id - smallint*/})

	file.WriteBytes([]byte("player_totem_info,pid,ingot_call_daily_num,ingot_call_timestamp,rock_rune_num,jade_rune_num,pos1,pos2,pos3,pos4,pos5,ingot_draw_times"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*ingot_call_daily_num - smallint*/, 4/*ingot_call_timestamp - bigint*/, 3/*rock_rune_num - int*/, 3/*jade_rune_num - int*/, 4/*pos1 - bigint*/, 4/*pos2 - bigint*/, 4/*pos3 - bigint*/, 4/*pos4 - bigint*/, 4/*pos5 - bigint*/, 2/*ingot_draw_times - smallint*/})

	file.WriteBytes([]byte("player_town,pid,town_id,lock,at_pos_x,at_pos_y"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 2/*town_id - smallint*/, 3/*lock - int*/, 2/*at_pos_x - smallint*/, 2/*at_pos_y - smallint*/})

	file.WriteBytes([]byte("player_trader_refresh_state,id,pid,last_update_time,auto_update_time,trader_id,refresh_num"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 4/*last_update_time - bigint*/, 4/*auto_update_time - bigint*/, 2/*trader_id - smallint*/, 2/*refresh_num - smallint*/})

	file.WriteBytes([]byte("player_trader_store_state,id,pid,trader_id,grid_id,item_id,num,cost,stock,goods_type"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 2/*trader_id - smallint*/, 3/*grid_id - int*/, 2/*item_id - smallint*/, 2/*num - smallint*/, 4/*cost - bigint*/, 1/*stock - tinyint*/, 1/*goods_type - tinyint*/})

	file.WriteBytes([]byte("player_use_skill,id,pid,role_id,skill_id1,skill_id4,skill_id2,skill_id3"))
	file.WriteBytes([]byte{4/*id - bigint*/, 4/*pid - bigint*/, 1/*role_id - tinyint*/, 2/*skill_id1 - smallint*/, 2/*skill_id4 - smallint*/, 2/*skill_id2 - smallint*/, 2/*skill_id3 - smallint*/})

	file.WriteBytes([]byte("player_vip,pid,ingot,level,card_id,present_exp"))
	file.WriteBytes([]byte{4/*pid - bigint*/, 4/*ingot - bigint*/, 2/*level - smallint*/, 7/*card_id - varchar*/, 4/*present_exp - bigint*/})

	return file.EndTrans()
}

