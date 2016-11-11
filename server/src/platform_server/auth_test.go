package platform_server

import (
	"core/log"
	"encoding/json"
	"fmt"
	"testing"
)

// token 会过期，所以本test case不输出Error

func TestAuthBefore(t *testing.T) {
	log.Setup("./log", false)
}

func TestCheckLogin(t *testing.T) {
	fmt.Println("==CheckToken==")

	//由于token检查在沙盒环境下有很大的网络延迟，故平常暂不测试
	fmt.Println("pass token check")
	return

	//ios 手Q
	out, err := CheckToken(2,
		"3CB0518F3C69B822A5E0610340AB7D60",
		"_TEST_1C8ED77F2B3B1314FBF7CDE4214082F4",
		"127.0.0.1")
	if err != nil {
		fmt.Println("error:", err)
	}
	b, err := json.Marshal(out)
	fmt.Printf("check ios 手Q token result:%v\n", string(b))

	//android 手Q
	out, err = CheckToken(18,
		"3CB0518F3C69B822A5E0610340AB7D60",
		"_TEST_1C8ED77F2B3B1314FBF7CDE4214082F4",
		"127.0.0.1")
	if err != nil {
		fmt.Println("error:", err)
	}
	b, err = json.Marshal(out)
	fmt.Printf("check 安卓 手Q token result:%v\n", string(b))

	//ios 微信
	out, err = CheckToken(1,
		"o1KVXt2ib_YsTJt8MqJMKXL9Djik",
		"_TEST_OezXcEiiBSKSxW0eoylIeBA6PSvR_UvHV-Jmo7pAiw5mfTwtHU7h-TGLozmMPVjokXsPP-0EWvS2yJjCnuneG9Qpjf2JkO-q5O7w5j3gOp2rDi6usLFOrAk2Pgb1JPYmM_SLXZUbjRpZVF7bPYQ4fA",
		"127.0.0.1")
	if err != nil {
		fmt.Println("error:", err)
	}

	b, err = json.Marshal(out)
	fmt.Printf("check ios 微信 token result:%v\n", string(b))

	//android 微信
	out, err = CheckToken(17, "o1KVXt2ib_YsTJt8MqJMKXL9Djik",
		"_TEST_OezXcEiiBSKSxW0eoylIeBA6PSvR_UvHV-Jmo7pAiw5mfTwtHU7h-TGLozmMPVjoHAqZQ6EGkNtHf_TuGlP_Ixx5VZbT1YMm9_JTeZ5RMnPlwpwqvXxtcfy0lGh9h16oiH3jEcR95aAs3S8KYjXOjw",
		"127.0.0.1")
	if err != nil {
		fmt.Println("error:", err)
	}

	b, err = json.Marshal(out)
	fmt.Printf("check 安卓 微信 token result:%v\n", string(b))

	//ios 游客
	out, err = CheckToken(5, "G_74YCaivFGz0B96LgLGB5lrDxFXxxtcUs",
		"_TEST_O2SSHQ9Zrw4dwJKwP6eWkn5fgY2wO2f4ABQb5FXbssY=",
		"127.0.0.1")
	if err != nil {
		fmt.Println("error:", err)
	}

	b, err = json.Marshal(out)
	fmt.Printf("check ios 游客 token result:%v\n", string(b))
}
