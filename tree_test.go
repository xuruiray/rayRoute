package rayRoute

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNode_InsertNode(t *testing.T) {

	head := Node{}
	for _, v := range TestRouteCase {
		value := Value(func(w http.ResponseWriter, r *http.Request) {})
		head.InsertNode(v.path, value)
	}

	for _, v := range TestRouteCase {
		result := head.FindNode(v.path)
		if result == nil {
			fmt.Println(v.path)
		}
	}

}

func Test_findCommonPrefix(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			name:  "case1",
			args:  args{"tianqibucuo", "tianqizhenhao"},
			want:  "tianqi",
			want1: len("tianqi"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findCommonPrefix(tt.args.s1, tt.args.s2)
			if got != tt.want {
				t.Errorf("findCommonPrefix() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findCommonPrefix() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
