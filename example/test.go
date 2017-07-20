package main

import (
	"net/http"
	"rcore"
	"middleware"
)

func main(){
	//创建路由复用器
	re := rcore.CreateNewRemux()

	//添加中间件
	re.AddMiddleware(testMiddleware)
	re.AddMiddleware(middleware.PanicHandler)

	//设置URL映射
	re.SetHandlerMapping("/",http.HandlerFunc(CallbackFun))
	re.SetHandlerMapping("/panic",http.HandlerFunc(panicTest))

	//开始监听并阻塞
	http.ListenAndServe(":80",re)
}


//自主编写的Controller
func CallbackFun(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world\n"))
}

func panicTest(w http.ResponseWriter,req *http.Request){
	panic("123912-miss params")
}


//自主编写的middleware
func testMiddleware(next http.HandlerFunc) http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("forward\n"))
		//下一个逻辑
		next.ServeHTTP(w,req)
		w.Write([]byte("backward\n"))
	}
	return http.HandlerFunc(f)
}


