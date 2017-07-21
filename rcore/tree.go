package rcore

import (
	"net/http"
	"strings"
)

type Node struct {
	label byte
	prefix string
	children []*Node
	handler http.HandlerFunc
}

func (n *Node)InsertNode(urlStr string,handlerFunc http.HandlerFunc){

	for _,v := range n.children{
		if v.label==urlStr[0] {
			if strings.HasPrefix(urlStr,v.prefix){
				if(len(urlStr)==len(v.prefix)){
					v.handler = handlerFunc
					n = v
					break
				}else {
					v.InsertNode(urlStr[len(v.prefix):], handlerFunc)
				}
			}
		}
	}

	if n.handler==nil{
		node := Node{label:urlStr[0],prefix:urlStr,handler:handlerFunc}
		n.children = append(n.children,&node)
	}

}

func (n *Node)FindNode(urlStr string)http.HandlerFunc{
	for _,v := range n.children{
		if v.label==urlStr[0] {
			if strings.HasPrefix(urlStr,v.prefix){
				if(len(urlStr)==len(v.prefix)){
					return v.handler
				}else{
					return v.FindNode(urlStr)
				}
			}
		}
	}

	return nil
}