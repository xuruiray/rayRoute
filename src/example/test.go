package main

import (
	"net/http"
	"rcore"
	"fmt"
)

func main(){
	re := rcore.CreateNewRemux()
	re.AddMiddleware(testMiddleware)
	re.AddMiddleware(panicHandler)

	re.SetHandlerMapping("/",http.HandlerFunc(CallbackFun))
	re.SetHandlerMapping("/panic",http.HandlerFunc(panicTest))

	http.ListenAndServe(":80",re)
}


func CallbackFun(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world\n"))
}

func panicTest(w http.ResponseWriter,req *http.Request){
	panic("err")
}


func testMiddleware(next http.HandlerFunc) http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("forward\n"))
		next.ServeHTTP(w,req)
		w.Write([]byte("backward\n"))
	}
	return http.HandlerFunc(f)
}

func panicHandler(next http.HandlerFunc) http.HandlerFunc{

	f := func(w http.ResponseWriter,req *http.Request){

		defer func(){
			err := recover()
			if err!=nil {
				fmt.Println(err)
			}
		}()

		w.Write([]byte("panic forward\n"))
		next.ServeHTTP(w,req)
		w.Write([]byte("panic backward\n"))
	}
	return http.HandlerFunc(f)
}
