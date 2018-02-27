package rayRoute

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNode_InsertNode(t *testing.T) {

	tests := []struct {
		name string
		path string
		want http.HandlerFunc
	}{
		{
			name: "正常流程 01",
			path: "/a1/b1/c1",
			want: func(w http.ResponseWriter, r *http.Request) {},
		}, {
			name: "正常流程 02",
			path: "/a1/b2/c1",
			want: func(w http.ResponseWriter, r *http.Request) {},
		}, {
			name: "正常流程 03",
			path: "/a2/b2/c1",
			want: func(w http.ResponseWriter, r *http.Request) {},
		}, {
			name: "正常流程 04",
			path: "/1a/1b/1c",
			want: func(w http.ResponseWriter, r *http.Request) {},
		},
	}

	head := Node{}
	for _, v := range tests {
		head.InsertNode(v.path, v.want)
		result := head.FindNode(v.path)
		assert.Equal(t, fmt.Sprintf("%v", v.want), fmt.Sprintf("%v", result), "name: %v", v.name)
	}

}
