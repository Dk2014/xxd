package gift_code

import (
	"fmt"
	"testing"
)

func TestGenCode(t *testing.T) {

	codeStrSet := make(map[string]bool, 1000)
	for i := 0; i < 1000; i++ {
		GenCode(codeStrSet)
	}
	for code, _ := range codeStrSet {
		fmt.Println(code)
	}
}
