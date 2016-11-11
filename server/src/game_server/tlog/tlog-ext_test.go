// 本文件由 tlog-parser 根据 日志xml规范生成
// 请勿手动改写！
package tlog

import (
	"os"
	"testing"
	"time"
)

var logs []Log

func Test(t *testing.T) {

	logs = append(logs, NewGameSvrState("test string"))

	logs = append(logs, NewPlayerRegister("test string", "test string", "test string", "test string", "test string", "test string", 123, 123, 123.456, 123, "test string", 123, "test string", "test string", "test string"))

	logs = append(logs, NewPlayerLogin("test string", 123, 123, "test string", "test string", "test string", "test string", "test string", 123, 123, 123.456, 123, "test string", 123, "test string", "test string", "test string"))

	logs = append(logs, NewPlayerLogout("test string", 123, 123, 123, "test string", "test string", "test string", "test string", "test string", 123, 123, 123.456, 123, "test string", 123, "test string", "test string", "test string", 123))

	logs = append(logs, NewMoneyFlow(123, "test string", 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewItemFlow(123, "test string", 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewItemMoneyFlow(123, "test string", 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerExpFlow("test string", 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerMissionFlow("test string", 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerPhysicalFlow("test string", 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerPvpFlow("test string", 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerGuideFlow("test string", 123, 123, 123, 123))

	logs = append(logs, NewPlayerSystemModuleFlow("test string", 123, 123, 123))

	logs = append(logs, NewPlayerSwordDrawFlow("test string", 123, 123, 123, 123, 123))

	logs = append(logs, NewPlayerChestDrawFlow("test string", 123, 123, 123, 123, "test string", 123))

	logs = append(logs, NewPlayerFightNumFlow("test string", 123, 123, 123))

	logs = append(logs, NewPlayerQuestFlow("test string", 123, 123, 123, 123, 123))

	logs = append(logs, NewBusinessManFlow("test string", 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewTradeBuyFlow("test string", 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewEquipRefineFlow("test string", 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewEquipDecomposeFlow("test string", 123, 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewGhostTrainFlow("test string", 123, 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewGhostUpgradeFlow("test string", 123, 123, 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewGhostSwordUpgradeFlow("test string", 123, "test string", 123, 123, 123, 123))

	logs = append(logs, NewSendHeartFlow("test string", 1234567890, 123, 123))

	logs = append(logs, NewAddPetFlow("test string", 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123))

	logs = append(logs, NewUpgradePetGridFlow("test string", 123, 123, 123, 123, 123))

	logs = append(logs, NewIDIPFLOW(1234567, 1234567, 1234567, "test string", 1234567, 1234567, 1234567, "test string", 1234567, 1234567))

	f, err := os.Create("../../testdata/log.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, log := range logs {
		debugWriteToFile(log, f)
	}
}
