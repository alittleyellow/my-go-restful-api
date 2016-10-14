package router

import (
	"net/http"
	"fmt"
	"regexp"
)

type Router struct {
	Url string
}

func NewRouter() *Router {
	return &Router{}
}

func (this *Router) Handler(resp http.ResponseWriter, req *http.Request)  {
	match, _ := regexp.MatchString("/favicon.ico", req.URL.Path)
	if match == true {
		
	} else {
		this.Url = req.URL.Path
	}	
	fmt.Println(this.Url)
}