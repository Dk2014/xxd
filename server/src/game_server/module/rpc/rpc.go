package rpc

import (
	. "game_server/rpc"
)

var Remote *RPC = NewRPC()

type RemoteServe int

func init() {
	Remote.Register(new(RemoteServe))
}
