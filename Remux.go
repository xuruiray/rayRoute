package rayRoute

import (
	"net/http"
)

const(
	ERROR_URL = 1000
)

type Remux struct {
	HandlerMapping map[string]http.Handler
	middleChan []func(http.ResponseWriter,*http.Request)
}

func(re *Remux) ServeHTTP(w http.ResponseWriter, r *http.Request){

}

func(re *Remux) SetHandlerMapping (urlStr string,handler http.Handler){
	re.HandlerMapping[urlStr] = handler
}


func(re *Remux) getHandlerMapping (urlStr string) http.Handler{
	if v,ok:=re.HandlerMapping[urlStr];ok{
		return v
	}else {
		panic(string(ERROR_URL)+"-this url hasn't be set")
	}
}

func CreateNewRemux() Remux{
	return Remux{make(map[string]http.Handler),	make([]func(http.ResponseWriter,*http.Request),5)}
}