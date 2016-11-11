<?php

$sample_pet_id=$this->NewQuery("
select min(`pet_id`) `id` from `battle_pet_skill_training`;
");
$pet_id=-1;
if($row=$sample_pet_id->GoNext()){
	$pet_id=$row['id'];
}
$this->DropQuery($sample_pet_id);
$this->AddSQL("
delete from `battle_pet_skill_training` where `pet_id`!={$pet_id};
");

