package rayRoute

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func init() {
	//创建路由复用器
	re := CreateNewRemux()

	//添加中间件
	re.AddMiddleware(testMiddleware)

	//设置URL映射
	re.SetHandlerMapping("/hello", Hello)

	//开始监听并阻塞
	err := http.ListenAndServe(":8001", re)
	if err != nil {
		fmt.Println(err)
	}
}

//自主编写的Controller
func Hello(conntext context.Context, req *http.Request) string {
	return "hello world\n"
}

//自主编写的middleware
func testMiddleware(next http.HandlerFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("forward\n"))
		//下一个逻辑
		next.ServeHTTP(w, req)
		w.Write([]byte("backward\n"))
	}
	return http.HandlerFunc(f)
}

func TestCreateNewRemux(t *testing.T) {

}
