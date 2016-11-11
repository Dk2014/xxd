package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

/*
 * usage:
 * platform_bench <server ip:port>
 *   stress test server apis
 * platform_bench <platform_server_redis ip:port> <platform_server_mysql ip:port>
 *   test cases and compare the output
 */

/*
 * TODO:
 * 实现上述接口
 * 完善的错误输出
 */
const (
	MaxConcurrentCalls = 300
)

var (
	w       = new(sync.WaitGroup)
	c_ok    = int32(0)
	c_t     = int64(0)
	uri_pre = "http://"
)

func usage() {
	flag.PrintDefaults()
}
func main() {
	var (
		optStress = flag.String("stress", "", "platform_bench <server ip:port>")
		optHelp   = flag.Bool("h", false, "this help")
	)

	flag.Parse()

	if *optHelp {
		usage()
		return
	}

	e := setrlimit(syscall.Rlimit{Cur: 32667, Max: 32667})
	if e != nil {
		fmt.Println("Rlimit error:", e)
	}

	switch {
	case len(*optStress) > 0:
		stressTest(*optStress)
	}

	fmt.Println("Done!")
}

//http request
func request(uri string, body io.Reader) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Errorf("Recovered in %v", r)
		}
		w.Done()
	}()
	tb := time.Now()
	req, _ := http.NewRequest("POST", uri, body)
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		fmt.Println(e)
		return
	}
	result := make(map[string]interface{})
	j := json.NewDecoder(resp.Body)
	j.Decode(&result)
	_, exist := result["error"]
	if !exist {
		atomic.AddInt32(&c_ok, 1)
		atomic.AddInt64(&c_t, int64(time.Since(tb)))
	} else {
		fmt.Println(result)
	}
}
func send(api string, rds [MaxConcurrentCalls]io.Reader) {
	uri := uri_pre + api
	c_ok = 0
	c_t = 0
	ta := time.Now()
	for _, rd := range rds {
		w.Add(1)
		go request(uri, rd)
	}
	w.Wait()
	t := time.Since(ta)
	fmt.Println("==" + api + "==")
	fmt.Println("avg time:", time.Duration(c_t/int64(c_ok)))
	fmt.Println("time used:", t.String())
	fmt.Printf("OK: %v, Total: %v\n", c_ok, MaxConcurrentCalls)
}

//stress test server apis
func stressTest(server string) {
	uri_pre = uri_pre + server
	var (
		jsonServers  [MaxConcurrentCalls]io.Reader
		jsonCreation [MaxConcurrentCalls]io.Reader
		jsonLogon    [MaxConcurrentCalls]io.Reader
		jsonList     [MaxConcurrentCalls]io.Reader
		jsonUpdate   [MaxConcurrentCalls]io.Reader
	)
	//产生随机请求
	for i := 0; i < MaxConcurrentCalls; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := r.Int63()
		jsonServers[i] = strings.NewReader(fmt.Sprintf(`
{
    "OpenId": "_test_id%v",
    "type": 1
}`, id))
		jsonCreation[i] = strings.NewReader(fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
    "RoleId": 1,
    "Nick": "_test_nick%v"
  }`, id, id))
		jsonLogon[i] = strings.NewReader(fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1
  }`, id))
		jsonList[i] = strings.NewReader(fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1
  }`, id))
		jsonUpdate[i] = strings.NewReader(fmt.Sprintf(`
  {
    "OpenId": "_test_id%v",
    "Type": 1,
    "Sid": 1,
    "RoleLevel": 3
  }`, id))
	}

	send("/server/list", jsonServers)
	send("/user/create", jsonCreation)
	send("/user/logon", jsonLogon)
	send("/user/list", jsonList)
	send("/user/update", jsonUpdate)
}

//test cases and compare the output
func compareTest(server_redis, server_mysql string) {
}

func setrlimit(rl syscall.Rlimit) (err error) {
	if rl.Cur > 0 && rl.Max > 0 {
		var lim syscall.Rlimit
		if err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &lim); err != nil {
			fmt.Println("failed to get NOFILE rlimit: ", err)
			return
		}
		fmt.Println(lim)
		lim.Cur = rl.Cur
		if rl.Max > lim.Max {
			lim.Max = rl.Max
		}
		fmt.Println(lim)
		if err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &lim); err != nil {
			fmt.Println("failed to set NOFILE rlimit: ", err)
			return
		}
	}
	return
}
