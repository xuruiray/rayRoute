package main

import "net/http"

func main(){
	http.HandleFunc("/",CallbackFun)
	http.ListenAndServe(":80",CreateNew)
}

func CallbackFun(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world"))
}
