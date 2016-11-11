<?php

// ############### 在此配置 ############### //

$trans_const_config = array(
	array('rpc' => false, 'const' => "DISPOSE_EVENT", ),
	array('rpc' => false, 'const' => "CMD", ),
	array('rpc' => false, 'const' => "PREPARE_STORE_EVENT", ),
	array('rpc' => false, 'const' => "SESSION_STATE_TIMER", ),

	array('rpc' => false, 'const' => "BattleRoom", ),

	array('rpc' => true, 'const' => "AddRecord", ),
	array('rpc' => true, 'const' => "GetFriendsWithMultiLevel", ),
	array('rpc' => true, 'const' => "CreateAnnouncement", ),
	array('rpc' => true, 'const' => "EnterArena", ),
	array('rpc' => true, 'const' => "SwapPlayerRank", ),
	array('rpc' => true, 'const' => "MailSend", ),
	array('rpc' => true, 'const' => "InviteFriendWithMultiLevel", ),
	array('rpc' => true, 'const' => "AwardMultiLevel", ),
	array('rpc' => true, 'const' => "GetPlayerMultiLevelInfo", ),
	array('rpc' => true, 'const' => "SetPlayerInfo", ),
	array('rpc' => true, 'const' => "GetPlayerInfo", ),
	array('rpc' => true, 'const' => "GetPlayerInfoWithOpenId", ),
	array('rpc' => true, 'const' => "GetTargetPlayerRank", ),
	array('rpc' => true, 'const' => "NewPlayerFighter", ),
	array('rpc' => true, 'const' => "AwardArenaBox", ),
	array('rpc' => true, 'const' => "GetTargetFighterForDrivingSwordVisting", ),

	array('rpc' => false, 'const' => "TlogHeart", ),
	array('rpc' => true, 'const' => 'IdipBanUser', ),
	array('rpc' => true, 'const' => 'IdipUnBanUser', ),
	array('rpc' => true, 'const' => 'IdipGetUserInfo', ),
	array('rpc' => true, 'const' => 'IdipGetEquipInfo', ),
	array('rpc' => true, 'const' => 'IdipGetBagInfo', ),
	array('rpc' => true, 'const' => 'IdipGetSoulInfo', ),
	array('rpc' => true, 'const' => 'IdipSendItem', ),
	array('rpc' => true, 'const' => 'IdipDelItem', ),
	array('rpc' => true, 'const' => 'IdipUpdateMoney', ),
	array('rpc' => true, 'const' => 'IdipUpdateExp', ),
	array('rpc' => true, 'const' => 'IdipUpdatePhysical', ),
	array('rpc' => true, 'const' => 'IdipUpdateHeart', ),

	array('rpc' => true, 'const' => 'SendHeartToPlatformFriend', ),
	array('rpc' => true, 'const' => 'SendHeartToGameFriend', ),
	array('rpc' => true, 'const' => 'SendHeartToAllGameFriends', ),
	array('rpc' => true, 'const' => 'GetOnlineNumber', ),
	array('rpc' => false, 'const' => 'GlobalMail', ),
	array('rpc' => false, 'const' => 'GlobalMailLoading', ), //启动时需要从数据库加载数据
	array('rpc' => true, 'const' => "MailBatchSend", ),
	array('rpc' => true, 'const' => 'IdipGetSwordInfo', ),
	array('rpc' => true, 'const' => 'IdipGetPetInfo', ),
	array('rpc' => true, 'const' => 'IdipUpdateGameLamp', ),
	array('rpc' => true, 'const' => 'IdipSendMail', ),
	array('rpc' => false, 'const' => 'GlobalPostMan', ),
	array('rpc' => true, 'const' => 'DeleteItem', ),
	array('rpc' => true, 'const' => 'RefuseMultiLevelRoomInvite', ),
	array('rpc' => true, 'const' => 'StartMultiLevel', ),
	array('rpc' => true, 'const' => 'MultiLevelChangeBuddy', ),
	array('rpc' => true, 'const' => 'MultiLevelChangeForm', ),
	array('rpc' => false, 'const' => "DoLogin", ),
	array('rpc' => true, 'const' => 'IdipAqBanUser', ),
	array('rpc' => true, 'const' => 'IdipAqRelievePunish', ),
	array('rpc' => true, 'const' => 'IdipAqSendMail', ),
	array('rpc' => true, 'const' => 'IdipGetRankInfoGs', ),
	array('rpc' => true, 'const' => 'IdipGetRankInfoHd', ),
	array('rpc' => true, 'const' => 'IdipCleanData', ),
	array('rpc' => true, 'const' => 'IdipSetSoulLevel', ),
	array('rpc' => true, 'const' => 'IdipSetSwordLevel', ),
	array('rpc' => true, 'const' => 'IdipAddPet', ),
	array('rpc' => true, 'const' => 'GetArenaRank', ),
	array('rpc' => true, 'const' => 'IdipSendAllMail', ),
	array('rpc' => true, 'const' => 'IdipGameLampInfo', ),
	array('rpc' => true, 'const' => 'IdipDelNotice', ),

	array('rpc' => true, 'const' => 'TssUserLogin', ),
	array('rpc' => true, 'const' => 'TssUserLogout', ),
	array('rpc' => true, 'const' => 'TssRecvData', ),
	array('rpc' => true, 'const' => 'TssSendData', ),
	array('rpc' => true, 'const' => 'GetGlobalVipCount',),
	array('rpc' => true, 'const' => 'GetGlobalLevelVipCount',),
	array('rpc' => true, 'const' => 'OperateGroupBuyCount',),
	array('rpc' => false, 'const' => 'GlobalWorldChannel',),
	array('rpc' => true, 'const' => 'DecMoney',),
	array('rpc' => true, 'const' => 'IncMoney',),
	array('rpc' => true, 'const' => 'IdipGetUserFight',),
	array('rpc' => true, 'const' => 'IdipGetHardLevelStatus',),
	array('rpc' => true, 'const' => 'IdipAddVipExp',),
	array('rpc' => false, 'const' => 'HttpAsync',),
	array('rpc' => true, 'const' => 'AddWorldChannelTplMessage',),
	array('rpc' => false, 'const' => 'WebDebugger',),
	array('rpc' => true, 'const'=>'GetClubHouseMeditation',),
	array('rpc' => false, 'const' => 'GlobalCliquesalary', ),
	array('rpc' => true, 'const'=>'AddFame',),
	array('rpc' => true, 'const'=>'UpdateCliqueKongfuAttrib',),
	array('rpc' => true, 'const'=>'AddAwardExp',),
	array('rpc' => false, 'const' => "ESCORT_BOAT_TIMER", ),
	array('rpc' => true, 'const' => "EscortBoatHijackBattleWin", ),
	array('rpc' => true, 'const' => "EscortBoatRecoverBattleWin", ),
	array('rpc' => true, 'const' => "EscortBoatRecoverBattleLose", ),
	array('rpc' => true, 'const' => "UpdateClubHousePlayer", ),
	array('rpc' => true, 'const' => "PaymentNotify", ),
	array('rpc' => true, 'const' => "XdgmGetUserInfo", ),
	array('rpc' => true, 'const' => "XdgmGetUserRank", ),
	array('rpc' => true, 'const' => "XdgmCliqueList", ),
	array('rpc' => true, 'const' => "XdgmLockPlayer", ),
	array('rpc' => true, 'const' => "XdgmGagPlayer", ),
	array('rpc' => true, 'const' => "XdgmGetNormalEventInfo", ),
	array('rpc' => true, 'const' => "XdgmGetJsonEventInfo", ),
	array('rpc' => true, 'const' => "XdgmGetEventAwardInfo", ),
	array('rpc' => true, 'const' => "XdgmSendMail", ),
	array('rpc' => true, 'const' => "XdgmSendGameLamp", ),
	array('rpc' => true, 'const' => "XdgmSendMailAll", ),
	array('rpc' => true, 'const' => "XdgmUpdateNormalEventInfo", ),
	array('rpc' => true, 'const' => "XdgmUpdateEventAwards", ),
	array('rpc' => true, 'const' => "XdgmGetTextEvents", ),
	array('rpc' => true, 'const' => "XdgmSearchGameLamp", ),
	array('rpc' => true, 'const' => "XdgmDelGameLamp", ),
	array('rpc' => true, 'const' => 'XdgmAddVipExp', ),
	array('rpc' => true, 'const' => 'XdgmSetPaymentsPresent', ),
	array('rpc' => true, 'const' => 'GetPresentRule', ),
	array('rpc' => true, 'const' => 'XdgmReloadGiftCode', ),
	array('rpc' => true, 'const' => 'ArenaInitPlayerArenaRank'),
	array('rpc' => true, 'const' => 'IAPPurchaseDeliver', ),
	array('rpc' => true, 'const' => 'QueryPid', ),
	array('rpc' => true, 'const' => 'WegamesPaymentNotify', ),
	array('rpc' => true, 'const' => 'PlayerUpdateRankDatas'),
	array('rpc' => true, 'const' => 'PlayerRankInGot'),
    
	// todo
	// add more
);

function GenTransConst($output_dir) {
	global $trans_const_config;

	$tran_tag_head     = "TRANS_TAG_";
	$tran_tag_rpc_head = $tran_tag_head . 'RPC_';

	$cols = array();
	$key  = 1 << 16;

	foreach ($trans_const_config as $rs) {
		if (!$rs['rpc']) {
			$key += 1;
			continue;
		}

		$cols[] = $key;
		$cols[] = $tran_tag_rpc_head . 'Call_' . $rs['const'];
		$key += 1;

		$cols[] = $key;
		$cols[] = $tran_tag_rpc_head . 'Serve_' . $rs['const'];
		$key += 1;
	}

	$code = '';
	$code .= "package mdb\n\n";

	$code .= "func (file *SyncFile) initRPCTransInfo() error {\n";
	$code .= "	file.BeginTrans()\n";

	$code .= "	file.WriteBytes([]byte(\"";
	$code .= implode($cols, ',');
	$code .= "\"))\n\n";
	$code .= "	return file.EndTrans()\n";
	$code .= "}\n\n";

	$code .= <<<CODE
	const (

CODE;

	foreach ($trans_const_config as $key => $rs) {
		$Var = ($rs['rpc'] ? $tran_tag_rpc_head : $tran_tag_head . $rs['const']);

		if ($rs['rpc']) {
			$call  = $Var . 'Call_' . $rs['const'];
			$serve = $Var . 'Serve_' . $rs['const'];
			$code .= <<<CODE
		{$call}
		{$serve}

CODE;
		} else {
			$Var = ($key == 0) ? "{$Var} = iota + uint32(1)<<16" : $Var;
			$code .= <<<CODE
		{$Var}

CODE;
		}
	}

	$code .= <<<CODE
	)

CODE;


$code .= <<<CODE
	const (

CODE;

	foreach ($trans_const_config as $key => $rs) {
		if (!$rs['rpc']) {
			continue;
		}

		$Var = 'RPC_Remote_' . $rs['const'] . ' = ' . '"RemoteServe.' . $rs['const'] . '"';
		$code .= <<<CODE
		{$Var}

CODE;
	}

	$code .= <<<CODE
	)

CODE;

	save_code($output_dir . "/zd_trans_const.go", $code);
}
