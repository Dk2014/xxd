package battle_pet

import (
	"core/net"
	"game_server/api/protocol/battle_pet_api"
	"game_server/module"
)

func init() {
	battle_pet_api.SetInHandler(BattlePetAPI{})
}

type BattlePetAPI struct {
}

func (this BattlePetAPI) GetPetInfo(session *net.Session, in *battle_pet_api.GetPetInfo_In) {
	rsp := &battle_pet_api.GetPetInfo_Out{}
	getPetInfo(module.State(session), rsp)
	session.Send(rsp)
}

func (this BattlePetAPI) SetPet(session *net.Session, in *battle_pet_api.SetPet_In) {
	setPet(module.State(session), in.GridNum, in.PetId)
}

func (this BattlePetAPI) SetPetSwap(session *net.Session, in *battle_pet_api.SetPetSwap_In) {
	setPetSwap(module.State(session), in.FromGridNum, in.ToGridNum)
}

//func (this BattlePetAPI) UpgradeGrid(session *net.Session, in *battle_pet_api.UpgradeGrid_In) {
//	upgradeGrid(session, in.GridNum)
//}

func (this BattlePetAPI) UpgradePet(session *net.Session, in *battle_pet_api.UpgradePet_In) {
	rsp := &battle_pet_api.UpgradePet_Out{}
	rsp.Exp, rsp.Level = upgradePet(module.State(session), in.PetId)
	session.Send(rsp)
}

func (this BattlePetAPI) TrainingPetSkill(session *net.Session, in *battle_pet_api.TrainingPetSkill_In) {
	rsp := &battle_pet_api.TrainingPetSkill_Out{}
	trainingPetSkill(module.State(session), in.PetId)
	session.Send(rsp)
}
