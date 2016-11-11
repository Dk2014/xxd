package battle

import (
	"testing"
)

func TestGenEnemyFormPos(t *testing.T) {
	for i := 0; i < 1000; i++ {
		posForm := genEnemyFormPos()
		for j := len(reserve); j > 0; j-- {
			if posForm[len(posForm)-j] != reserve[len(reserve)-j] {
				t.Fatal(posForm)
			}
		}
	}
}
