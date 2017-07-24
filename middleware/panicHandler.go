package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func PanicHandler(next http.HandlerFunc) http.HandlerFunc {

	f := func(w http.ResponseWriter, req *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				str := fmt.Sprintf("", err)
				errParams := strings.Split(str[16:len(str)-2], "-")
				errStr := fmt.Sprintf("{\"errno\":%s,\"errmsg\":\"%s\",\"data\":null}", errParams[0], errParams[1])
				w.Write([]byte(errStr))
			}
		}()

		w.Write([]byte("panic forward\n"))
		next.ServeHTTP(w, req)
		w.Write([]byte("panic backward\n"))
	}
	return http.HandlerFunc(f)
}
