package rayRoute

import (
	"net/http"
	"strings"
)

type Value http.HandlerFunc

// Node 字典树节点
type Node struct {
	label    byte
	prefix   string
	children []*Node
	value    Value
}

// InsertNode 添加路由节点
func (n *Node) InsertNode(urlStr string, value Value) {

	for _, v := range n.children {
		if v.label == urlStr[0] {
			if strings.HasPrefix(urlStr, v.prefix) {
				if len(urlStr) == len(v.prefix) {
					v.value = value
					n = v
					break
				} else {
					v.InsertNode(urlStr[len(v.prefix):], value)
				}
			}
		}
	}

	if n.value == nil {
		node := Node{label: urlStr[0], prefix: urlStr, value: value}
		if n.children == nil {
			n.children = make([]*Node, 0)
		}
		n.children = append(n.children, &node)
	}

}

// FindNode 查找路由节点
func (n *Node) FindNode(urlStr string) Value {
	for _, v := range n.children {
		if v.label == urlStr[0] {
			if strings.HasPrefix(urlStr, v.prefix) {
				if len(urlStr) == len(v.prefix) {
					return v.value
				} else {
					return v.FindNode(urlStr)
				}
			}
		}
	}

	return nil
}

func (n *Node) PrintTree() []string {
	return nil
}