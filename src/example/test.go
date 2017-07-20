package main

import (
	"net/http"
	"rcore"
)

func main(){
	re := rcore.CreateNewRemux()
	re.AddMiddleware(testMiddleware)
	re.SetHandlerMapping("/",http.HandlerFunc(CallbackFun))
	http.ListenAndServe(":80",re)
}


func CallbackFun(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world\n"))
}


func testMiddleware(next http.HandlerFunc) http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte("forward\n"))
		next.ServeHTTP(w,req)
		w.Write([]byte("backward\n"))
	}
	return http.HandlerFunc(f)
}

