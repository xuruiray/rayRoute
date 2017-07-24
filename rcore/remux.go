package rcore

import (
	"net/http"
	"context"
)

type ReHandlerFun func(context.Context,*http.Request) (string)

type Remux struct {
	tree Node
	handlerMapping map[string]http.HandlerFunc
	middleHandler http.HandlerFunc
}

func(re *Remux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	re.middleHandler.ServeHTTP(w, r)
}

func(re *Remux) SetHandlerMapping (urlStr string,handlerFunc func(context.Context,*http.Request) (string)){
	//re.handlerMapping[urlStr] = handlerFunc
	re.tree.InsertNode(urlStr,packagefun(handlerFunc))
}

func(re *Remux) getHandlerMapping (urlStr string) (http.Handler) {
	//return re.handlerMapping[urlStr]
	return re.tree.FindNode(urlStr)
}

func(re *Remux) AddMiddleware(f func(handlerFunc http.HandlerFunc)http.HandlerFunc){
	if re.middleHandler == nil{
		re.middleHandler = f(re.defaultMiddleware())
	}else {
		re.middleHandler = f(re.middleHandler)
	}
}

func (re *Remux)defaultMiddleware() http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		fun := re.getHandlerMapping(req.RequestURI)
		fun.ServeHTTP(w,req)
	}
	return http.HandlerFunc(f)
}

func packagefun(fun ReHandlerFun) http.HandlerFunc{
	f := func(w http.ResponseWriter,req *http.Request){
		w.Write([]byte(fun(req.Context(),req)))
	}
	return http.HandlerFunc(f)
}

func CreateNewRemux() *Remux{
	re := Remux{handlerMapping:make(map[string]http.HandlerFunc)}
	return &re
}