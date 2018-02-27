# rayRoute

简易的路由框架，使用了基数树作为路由结构，添加了中间件层


### 安装
```go
go get github.com/xuruiray/rayRoute
```

### 示例
</hr>

```go
package main

import "github.com/xuruiray/rayRoute"
import "net/http"
import "context"

func main(){
	//创建路由复用器
	mux := rayRoute.CreateNewRemux()

	//添加中间件
	mux.AddMiddleware(testMiddleware)

	//绑定 controller
	mux.SetHandlerMapping("/hello",helloHandler)

	//开始监听并阻塞
	http.ListenAndServe(":8001",mux)
}


//自主编写的Controller
func helloHandler(conntext context.Context, req *http.Request) (string){
	return "hello world\n"
}

//自主编写的middleware
func testMiddleware(next http.HandlerFunc) http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("forward\n"))
		//下一个中间件逻辑
		next.ServeHTTP(w,req)
		w.Write([]byte("backward\n"))
	}
	return http.HandlerFunc(f)
}

```

### Radix Tree
![](http://photo.rhymecode.com/illustrations/radixTree.jpg)
