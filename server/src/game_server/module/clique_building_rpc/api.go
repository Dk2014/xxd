package clique_building_rpc

import (
	"core/net"
	"game_server/api/protocol/clique_building_api"
)

type CliqueBuildingAPI struct {
}

func init() {
	clique_building_api.SetInHandler(&CliqueBuildingAPI{})
}

// 帮派建筑物 总舵 捐赠铜钱
func (this *CliqueBuildingAPI) CliqueBaseDonate(session *net.Session, in *clique_building_api.CliqueBaseDonate_In) {
	cliqueBuildingBaseDonate(session, in.Money)
}

// 所有帮派建筑物状态 TODO 以后再加建筑物即修改这个协议增加建筑物状态
func (this *CliqueBuildingAPI) CliqueBuildingStatus(session *net.Session, in *clique_building_api.CliqueBuildingStatus_In) {
	cliqueBuildingStatus(session)
}

func (this *CliqueBuildingAPI) CliqueBankDonate(session *net.Session, in *clique_building_api.CliqueBankDonate_In) {
	cliqueBuildingBankDonate(session, in.Money)
}

func (this *CliqueBuildingAPI) CliqueBankBuy(session *net.Session, in *clique_building_api.CliqueBankBuy_In) {
	cliqueBankBuy(session, in.Kind, in.Num)
}

func (this *CliqueBuildingAPI) CliqueBankSold(session *net.Session, in *clique_building_api.CliqueBankSold_In) {
	cliqueBankSold(session, in.Kind)
}

func (this *CliqueBuildingAPI) CliqueKongfuDonate(session *net.Session, in *clique_building_api.CliqueKongfuDonate_In) {
	kongfuBuildingDonate(session, in.Building, in.Money)
}

func (this *CliqueBuildingAPI) CliqueKongfuInfo(session *net.Session, in *clique_building_api.CliqueKongfuInfo_In) {
	kongfuInfo(session, in.Building)
}

func (this *CliqueBuildingAPI) CliqueKongfuTrain(session *net.Session, in *clique_building_api.CliqueKongfuTrain_In) {
	kongfuTrain(session, in.KongfuId)
}

func (this *CliqueBuildingAPI) CliqueTempleWorship(session *net.Session, in *clique_building_api.CliqueTempleWorship_In) {
	cliqueTempleWorship(session, in)
}
func (this *CliqueBuildingAPI) CliqueTempleDonate(session *net.Session, in *clique_building_api.CliqueTempleDonate_In) {
	cliqueTempleDonate(session, in)
}

func (this *CliqueBuildingAPI) CliqueTempleInfo(session *net.Session, in *clique_building_api.CliqueTempleInfo_In) {
	cliqueTempleInfo(session, in)
}
