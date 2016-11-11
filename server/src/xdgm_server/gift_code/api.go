package gift_code

import (
	"math/rand"
)

type CodeType int8

const (
	MONOP_TYPE CodeType = 0
	SHARE_TYPE CodeType = 1
)

var (
	_CODE_LEN      = 8
	_CODE_ELEM_SET = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Gen(codeType CodeType, config string, effectTs, expireTs int64, serverId int, num int16, version int64, content string) ([]GiftCode, error) {
	var (
		ok         bool
		err        error
		codeStrSet map[string]bool
		codeSet    []GiftCode
	)
	codeStrSet = make(map[string]bool, num)
	code := GiftCode{
		Type:            codeType,
		EffectTimestamp: effectTs,
		ExpireTimestamp: expireTs,
		ServerId:        serverId,
		Version:         version,
		Content:         content,
		Config:          config,
	}
	if num <= 0 {
		num = 1
	}

	for num > 0 {
		code.Code = GenCode(codeStrSet)
		ok, err = InsertCode(code)
		if err != nil {
			return codeSet, err
		}
		if ok {
			num--
			codeSet = append(codeSet, code)
		}
	}
	return codeSet, nil
}

func GenCode(set map[string]bool) string {
	buff := make([]byte, _CODE_LEN, _CODE_LEN)
	for {
		idxs := rand.Perm(len(_CODE_ELEM_SET))[:_CODE_LEN]
		for i, idx := range idxs {
			buff[i] = _CODE_ELEM_SET[idx]
		}
		if _, exist := set[string(buff)]; !exist {
			set[string(buff)] = true
			return string(buff)
		}
	}
	return ""
}
