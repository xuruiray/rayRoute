### 简述
</hr>

扩展了golang自带的HTTP库
添加了中间件层


### demo
</hr>

```go

func main(){
	//创建路由复用器
	re := rcore.CreateNewRemux()

	//添加中间件
	re.AddMiddleware(testMiddleware)

	//设置URL映射
	re.SetHandlerMapping("/",http.HandlerFunc(CallbackFun))
	
	//开始监听并阻塞
	http.ListenAndServe(":80",re)
}


//自主编写的Controller
func CallbackFun(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world\n"))
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

```


![](http://img2-ak.lst.fm/i/u/300x300/6f39caa4a0fa4bc8cbfdf18390146df1.jpg)
