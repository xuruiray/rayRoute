package rayRoute

import (
	"net/http"
	"strings"
)

// Node 字典树节点
type Node struct {
	label    byte
	prefix   string
	children []*Node
	handler  http.HandlerFunc
}

// InsertNode 添加路由节点
func (n *Node) InsertNode(urlStr string, handlerFunc http.HandlerFunc) {

	for _, v := range n.children {
		if v.label == urlStr[0] {
			if strings.HasPrefix(urlStr, v.prefix) {
				if len(urlStr) == len(v.prefix) {
					v.handler = handlerFunc
					n = v
					break
				} else {
					v.InsertNode(urlStr[len(v.prefix):], handlerFunc)
				}
			}
		}
	}

	if n.handler == nil {
		node := Node{label: urlStr[0], prefix: urlStr, handler: handlerFunc}
		if n.children == nil {
			n.children = make([]*Node, 0)
		}
		n.children = append(n.children, &node)
	}

}

// FindNode 查找路由节点
func (n *Node) FindNode(urlStr string) http.HandlerFunc {
	for _, v := range n.children {
		if v.label == urlStr[0] {
			if strings.HasPrefix(urlStr, v.prefix) {
				if len(urlStr) == len(v.prefix) {
					return v.handler
				} else {
					return v.FindNode(urlStr)
				}
			}
		}
	}

	return nil
}
