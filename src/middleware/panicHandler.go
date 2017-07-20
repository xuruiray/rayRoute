package middleware

import (
	"net/http"
	"fmt"
)

func PanicHandler(next http.HandlerFunc) http.HandlerFunc{

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