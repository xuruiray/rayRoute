package main

import (
	"rcore"
	"net/http"
)

func main() {
	re := rcore.CreateNewRemux()
	re.SetHandlerMapping("/",CollectPCommentHandler)
}


func CollectPCommentHandler(w http.ResponseWriter,req *http.Request){
	w.Write([]byte("hello world\n"))
}