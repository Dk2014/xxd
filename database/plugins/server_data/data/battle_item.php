<?php

$generate = function($db) {
	$configlist = db_query($db, "select item.name,bc.* from battle_item_config bc left join item on item.id = bc.id");

	$code = "package battle\n\n";
	$code.= generateTypeComment();
	$code.= generateCode_CreateBattleItem($configlist);
	save_code(dirname(__FILE__)."/../../../../server/src/game_server/battle/battle_item_auto.go", $code);
};

function generateTypeComment() {
	global $battle_item_target, $battle_item_effect_type,$battle_item_effect_normal_mod,$battle_item_effect_buff_mod;

	$code = "	目标对象\n";
	foreach ($battle_item_target as $key => $value) {
		$code .= "	{$key} -- {$value['name']}\n";
	}

	$code .= "\n\n	产生效果\n";
	foreach ($battle_item_effect_type as $key => $value) {
		$code .= "	{$key} -- {$value['name']}\n";
	}

	$code .= "\n\n	基础属性类型\n";
	foreach ($battle_item_effect_normal_mod as $key => $value) {
		$name = is_array($value) ? $value['name']:$value;
		$code .= "	{$key} -- {$name}\n";
	}

	$code .= "\n\n	Buff增益类型(对应buff类型)\n";
	foreach ($battle_item_effect_buff_mod as $key => $value) {
		$code .= "	{$key} -- ({$value['const_var']}){$value['name']}\n";
	}

	return <<<C

	/*
{$code}
	*/

C;
}

function generateCode_CreateBattleItem(&$configlist) {
	$code = "";
	$itemVarList = "";
	foreach ($configlist as $key => $item) {
		$code .= "	case {$item['id']}:\n";
		$code .= "		return battle_item_{$item['id']} // {$item['name']}\n";
		$itemVarList .= generateCode_BattleItem($item);
	}

	return <<<C

func createBattleItem(id int32) *Item {
	switch id {
{$code}		
	}
	return nil
}

{$itemVarList}

C;
}

function generateCode_BattleItem($item) {
	global $battle_item_target, $battle_item_effect_type;

	$config = json_decode($item['config']);
	$callFn = $battle_item_effect_type[$item['effect_type']]['call'];

	$effectCode = "";
	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		$effectCode .= <<<EC
		&ItemEffect{
			Mode:  {$rs->mod},
			Value: {$rs->val},
		},
EC;
	}

	$applyCode = "";
	$targetCode = $battle_item_target[$item['target_type']]['code'];
	$applyFuncCode = $callFn($config, $item['can_use_level'], $applyCode);
	return <<<C

//{$item['name']}
var battle_item_{$item['id']} = &Item{
	ItemId:      {$item['id']},
	Target:      {$item['target_type']},
	EffectType:  {$item['effect_type']},
	Keep:        {$item['keep']},
	CostPower: 	 {$item['cost_power']},
	MaxOverride: {$item['max_override']},
	Effect: []*ItemEffect{
{$effectCode}
	},
{$applyCode}
	_GetTargets: func(user *Fighter) []*Fighter {
		return {$targetCode}
	},
	_applyFunc: func(item *Item, user *Fighter, target *Fighter) []*Buff {
{$applyFuncCode}
	},
}

C;
}

function applyItemForNormalCode($config, &$applyCode) {
	global $battle_item_effect_normal_mod;
	$arr = array();
	for ($i=0; $i < count($config); $i++) {
		$rs = $config[$i];
		$attr = $battle_item_effect_normal_mod[$rs->mod]['var'];
		$arr[] = "		attr.{$attr} += ({$rs->val})";
	}
	$code = implode(',', $arr);

	$applyCode = <<<C
	ApplyItemForNormal: func(attr *FighterAttribute) {
{$code}
	},
C;
}

function applyItemForBuffCode($config, &$applyCode) {
	$arr = array();
	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		$arr[] = "fighter.addBuff(fighter, {$rs->mod}, {$rs->val}, item.Keep, 0, item.MaxOverride, false)";
	}
	$code = implode("\n		", $arr);

	$applyCode = <<<C
	ApplyItemForBuff: func(fighter *Fighter, item *Item) {
		{$code}
	},
C;
}

function catchCode($config) {
return <<<C
		catchPet(user, target, item.Effect[0])
		return nil
C;
}

function configBuffCode($config, $can_use_level, &$applyCode) {
	if ($can_use_level) {
		applyItemForBuffCode($config, $applyCode);
	}

	$arr = array();
	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		$arr[] = "target.addBuff(user, {$rs->mod}, {$rs->val}, item.Keep, 0, item.MaxOverride, false)";
	}
	$code = implode(',', $arr);

return <<<C
		return []*Buff{{$code}}
C;
}

function configNormalCode($config, $can_use_level, &$applyCode) {
	global $battle_item_effect_normal_mod;

	if ($can_use_level) {
		applyItemForNormalCode($config, $applyCode);
	}
	
	$arr = array();
	for ($i=0; $i < count($config); $i++) { 
		$rs = $config[$i];
		if (array_key_exists('call', $battle_item_effect_normal_mod[$rs->mod])) {
			$call = $battle_item_effect_normal_mod[$rs->mod]['call'];
			$arr[] = "		{$call}(user, target, item.Effect[{$i}])";
		} else {
			$max = $battle_item_effect_normal_mod[$rs->mod]['max'];
			$min = $battle_item_effect_normal_mod[$rs->mod]['min'];
			$attr = $battle_item_effect_normal_mod[$rs->mod]['var'];
			$fixCode = attrFixCode($attr, $max, $min, $rs->val);
			$arr[] = "		target.{$attr} += ({$rs->val}){$fixCode}";
		}
	}
	$code = implode("\n", $arr);	
return <<<C
{$code}
		return nil
C;
}


function attrFixCode($attr, $max, $min, $val) {
	$fixCode = "";
	$sign = "";
	$fixVar =  "";

	if ($val > 0) {
		$sign = ">";
		$fixVar = $max;
	} else {
		$sign = "<";
		$fixVar = $min;
	}

	$fixCode = "\n		if target.{$attr} {$sign} {$fixVar} {\n";
	$fixCode .= "			target.{$attr} = {$fixVar}\n";
	$fixCode .= "		}\n";
	return $fixCode;
}

?>
