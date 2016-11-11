package i18l

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	T *Translation
)

type Translation struct {
	filepath     string
	CurrentLocal string
	Locales      map[string]map[string]string
}

const (
	DEFAULTLAN = "zh"
)

func NewLocale(filepath string, defaultlocal string) *Translation {
	i18n := make(map[string]map[string]string)
	if filepath != "" {
		file, err := os.Open(filepath)
		if err != nil {
			panic("open " + filepath + " err :" + err.Error())
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			panic("read " + filepath + " err :" + err.Error())
		}
		err = json.Unmarshal(data, &i18n)
		if err != nil {
			panic("json.Unmarshal " + filepath + " err :" + err.Error())
		}
	}
	return &Translation{
		filepath:     filepath,
		CurrentLocal: defaultlocal,
		Locales:      i18n,
	}
}

func (t *Translation) SetLocale(local string) {
	t.CurrentLocal = local
}

func (t *Translation) TranL(key string, local string) string {
	if local == "" {
		local = t.CurrentLocal
	}
	if ct, ok := t.Locales[local]; ok {
		if v, o := ct[key]; o {
			return v
		}
	}
	return key
}

func (t *Translation) Tran(key string) string {
	if ct, ok := t.Locales[t.CurrentLocal]; ok {
		if v, o := ct[key]; o {
			return v
		}
	}
	return key
}
