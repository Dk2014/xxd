package module

import (
	"core/log"
	"game_server/mdb"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	HTTP_ASYNC_MAX          = 10000
	HTTP_ASYNC_CONN_TIMEOUT = 5 // second
	HTTP_ASYNC_RW_TIMEOUT   = 5 // second
	HTTP_ASYNC_REQ_TIMEOUT  = 5 // second
)

var UrlServer *UrlAPIServer

type UrlWork struct {
	Req      *http.Request
	Callback func(result string)
}

type UrlAPIServer struct {
	workChan chan UrlWork
}

type TimeoutConn struct {
	conn    net.Conn
	timeout time.Duration
}

func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
	return &TimeoutConn{
		conn:    conn,
		timeout: timeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	c.SetReadDeadline(time.Now().Add(c.timeout))
	return c.conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	c.SetWriteDeadline(time.Now().Add(c.timeout))
	return c.conn.Write(b)
}

func (c *TimeoutConn) Close() error {
	return c.conn.Close()
}

func (c *TimeoutConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TimeoutConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TimeoutConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *TimeoutConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *TimeoutConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (us *UrlAPIServer) Push(work UrlWork) {
	select {
	case us.workChan <- work:
	default:
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Errorf("UrlAPIServer push work failed: %v", err)
				}
			}()
			us.workChan <- work
		}()
	}
}

// 创建GET请求对象
func (us *UrlAPIServer) NewRequestForGet(urlStr string) (req *http.Request, err error) {
	req, err = http.NewRequest("GET", urlStr, nil)

	if err != nil {
		log.Errorf(`[UrlAPIServer] new request for GET fail
url = %s
Error = %v
`, urlStr, err)
	}
	return
}

// 创建POST请求对象
func (us *UrlAPIServer) NewRequestForPost(urlStr string, data string) (req *http.Request, err error) {
	req, err = http.NewRequest("POST", urlStr, strings.NewReader(data))

	if err != nil {
		log.Errorf(`new request for POST fail
url = %s
Error = %v
`, urlStr, err)
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	return
}

func init() {
	UrlServer = &UrlAPIServer{
		workChan: make(chan UrlWork, HTTP_ASYNC_MAX),
	}

	go func() {
		client := &http.Client{
			Transport: &http.Transport{
				Dial: func(netw, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(netw, addr, time.Second*HTTP_ASYNC_CONN_TIMEOUT)

					if err != nil {
						return nil, err
					}

					return NewTimeoutConn(conn, time.Second*HTTP_ASYNC_RW_TIMEOUT), nil
				},
				ResponseHeaderTimeout: time.Second * HTTP_ASYNC_REQ_TIMEOUT,
			},
		}

		for {
			work := <-UrlServer.workChan
			rsp, err := client.Do(work.Req)
			if err != nil {
				log.Errorf(`HTTP Request fail
url = %s
error = %v
`, work.Req.URL, err)
				continue
			}

			func() {
				defer rsp.Body.Close()
				result, err := ioutil.ReadAll(rsp.Body)

				if err != nil {
					log.Errorf(`http read body fail
url = %s
error = %v
`, work.Req.URL, err)
					return
				}

				if work.Callback == nil {
					return
				}

				defer func() {
					if err := recover(); err != nil {
						log.Errorf("http callback failed: %v, req: %v", err, work.Req.URL)
					}
				}()

				mdb.Transaction(mdb.TRANS_TAG_HttpAsync, func() {
					work.Callback(string(result))
				})
			}()
		}
	}()
}
