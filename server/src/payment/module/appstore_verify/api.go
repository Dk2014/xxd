package appstore_verify

import (
	"github.com/dogenzaka/go-iap/appstore"
	"payment/config"
	"payment/module"
	"time"
)

var (
	g_devClient     *appstore.Client
	g_productClient *appstore.Client
)

func init() {
	module.AppStoreVerify = AppStoreVerifyMod{}
}

type AppStoreVerifyMod struct{}

func (mod AppStoreVerifyMod) Init(timeoutSecond int) {
	cfg := appstore.Config{
		IsProduction: true,
		TimeOut:      time.Duration(timeoutSecond) * time.Second,
	}
	client := appstore.NewWithConfig(cfg)
	g_productClient = &client

	cfgDev := appstore.Config{
		IsProduction: false,
		TimeOut:      time.Duration(timeoutSecond) * time.Second,
	}
	clientDev := appstore.NewWithConfig(cfgDev)
	g_devClient = &clientDev
}

func (mod AppStoreVerifyMod) Verify6(receiptStr string, pid int64) (appstore.IAPResponse6, []byte, error) {
	receipt := appstore.IAPRequest{
		ReceiptData: receiptStr,
	}
	var isProduct = true
	serverId := int(pid >> 32 / 10)
	for _, sid := range config.CFG.AppStoreVerify.SandServerList {
		if sid == serverId {
			isProduct = false
			break
		}
	}
	if isProduct {
		result, rawRsp, err := g_productClient.Verify6(receipt)
		return result, rawRsp, err
	} else {
		result, rawRsp, err := g_devClient.Verify6(receipt)
		return result, rawRsp, err
	}
}
