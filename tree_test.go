package rayRoute

import (
	"net/http"
	"testing"
	"fmt"
)

func TestNode_InsertNode(t *testing.T) {

	head := Node{}
	for _, v := range TestRouteCase {
		value := Value(func(w http.ResponseWriter, r *http.Request) {})
		head.InsertNode(v.path, value)
	}

	for _, v := range TestRouteCase {
		result := head.FindNode(v.path)
		if result == nil{
			fmt.Println(v.path)
		}
	}

}
