package router

import (
	"net/http"
	"regexp"
	"encoding/json"
	"fmt"
)

type Router struct {
	Url string
}

func NewRouter() *Router {
	return &Router{}
}

func (this *Router) Handler(resp http.ResponseWriter, req *http.Request) {
	match, _ := regexp.MatchString("/favicon.ico", req.URL.Path)
	if match == true {

	} else {
		Url = req.URL.Path

		w, err := this.Url();
		if err == nil {
			fmt.Fprintf(resp, string(w))
		}
		
	}
}

func (this *Router) GetJson() ([]byte, error) {
	pass := make(map[string]string)
	pass["name"] = "huangyang"
	pass["age"] = "18"
	return json.MarshalIndent(pass, "", "")
}