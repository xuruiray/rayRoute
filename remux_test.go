package rayRoute

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"
)

func initServer(listener net.Listener, handler http.Handler) {
	svr := http.Server{Handler: handler}
	svr.Serve(listener)
}

func TestSetHandlerMapping(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		wantBody string
		wantErr  error
	}{
		{
			name:     "正常流程",
			path:     "/hello",
			wantBody: "hello",
			wantErr:  nil,
		},
		//TODO 测试用例待完善
	}

	re := CreateNewRemux()
	for _, v := range tests {

		f := func(context.Context, *http.Request) string {
			return v.wantBody
		}

		re.SetHandlerMapping(v.path, f)
	}

	//启动服务
	listener, _ := net.Listen("tcp", ":8001")
	go initServer(listener, re)

	for _, v := range tests {

		//发送测试请求
		resp, err := http.Get("http://localhost:8001" + v.path)
		assert.Equal(t, v.wantErr, err)
		if err != nil {
			continue
		}

		// 读取响应
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			assert.Fail(t, err.Error())
		}

		assert.Equal(t, v.wantBody, string(body))
	}
}

// TODO 测试应该不用这样写，写的手酸。。
func TestAddMiddleware(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		body         string
		wantForward  string
		wantBackward string
		want         string
		wantErr      error
	}{
		{
			name:         "正常流程",
			path:         "/test",
			body:         "test",
			wantForward:  "F",
			wantBackward: "B",
			want:         "FtestB",
			wantErr:      nil,
		}, {
			name:         "正常流程",
			path:         "",
			wantForward:  "Forward",
			wantBackward: "Backward",
			wantErr:      nil,
		},
	}

	for _, v := range tests {

		re := CreateNewRemux()

		// 绑定一个 middleware
		midf := func(next http.HandlerFunc) http.HandlerFunc {
			f := func(w http.ResponseWriter, req *http.Request) {
				w.Write([]byte(v.wantForward))
				next.ServeHTTP(w, req)
				w.Write([]byte(v.wantBackward))
			}
			return http.HandlerFunc(f)
		}
		re.AddMiddleware(midf)

		// 绑定一个 Handler
		handf := func(context.Context, *http.Request) string {
			return "test"
		}
		re.SetHandlerMapping("/test", handf)

		//启动服务
		listener, _ := net.Listen("tcp", ":8001")
		go initServer(listener, re)

		//发送测试请求
		resp, err := http.Get("http://localhost:8001" + v.path)
		assert.Equal(t, v.wantErr, err)
		if err != nil {
			continue
		}

		// 读取响应
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			assert.Fail(t, err.Error())
		}

		// 比较测试结果
		assert.Equal(t, v.wantForward+"test"+v.wantBackward, string(body))

		//关闭服务
		listener.Close()
	}
}

func TestServer(t *testing.T) {
	re := CreateNewRemux()
	//绑定一个 Handler
	f := func(context.Context, *http.Request) string {
		return "test"
	}
	re.SetHandlerMapping("/test", f)
	//启动服务
	listener, _ := net.Listen("tcp", ":8001")
	go func() {
		time.Sleep(1 * time.Second)
		//关闭服务
		listener.Close()
	}()
	initServer(listener, re)
}
