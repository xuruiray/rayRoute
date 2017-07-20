package rcore

import (
	"net/http"
)

const(
	ERROR_URL = 1000	//未匹配到URL
)

type Remux struct {
	handlerMapping map[string]http.HandlerFunc
	middleHandler http.HandlerFunc
}

func(re *Remux) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if _,ok:=re.handlerMapping[r.RequestURI];ok {
		re.middleHandler.ServeHTTP(w, r)
	}
}

func(re *Remux) SetHandlerMapping (urlStr string,handlerFunc http.HandlerFunc){
	re.handlerMapping[urlStr] = handlerFunc
}

func(re *Remux) getHandlerMapping (urlStr string) (http.Handler) {
	return re.handlerMapping[urlStr]
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

func CreateNewRemux() *Remux{
	re := Remux{handlerMapping:make(map[string]http.HandlerFunc)}
	return &re
}