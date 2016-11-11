package config_hub

import (
	"config_hub/args"
	"config_hub/eventconf"
	"config_hub/gameconf"
	"config_hub/giftcode"
	"config_hub/httpsrv"
)

func Main() {
	err := args.Parse()
	if err != nil {
		panic(err)
	}

	// load action modules
	giftcode.Load()
	defer giftcode.Unload()
	gameconf.Load()
	defer gameconf.Unload()
	eventconf.Load()
	defer eventconf.Unload()

	err = httpsrv.Start()
	if err != nil {
		panic(err)
	}
}
