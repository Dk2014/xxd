package daily_sign_in

import (
	"game_server/dat/daily_sign_in_dat"
)

func isNewPlayer(firstLoginDate int, date int) bool {
	return date-firstLoginDate < daily_sign_in_dat.NewPlayerSignInDuration
}
