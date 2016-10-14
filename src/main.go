package main 
	
import (
	"net/http"
	// "fmt"
	"router"
)

func main () {
	route := router.NewRouter();
	http.HandleFunc("/", route.Handler)
	http.ListenAndServe("127.0.0.1:8085", nil)
}















