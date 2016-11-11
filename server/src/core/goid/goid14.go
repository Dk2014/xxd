// +build go1.4

package goid

import "runtime"

//go1.4以及以后，只能在runtime包内修改代码提供改函数
func GetGoroutineId() int64 {
	return runtime.GetGoroutineId()
}
