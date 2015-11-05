package main

import (
	"fmt"
	"net/http"
	"net"
	"os"
)


func handler(w http.ResponseWriter,r *http.Request){
     h, _ := os.Hostname()
     addrs, _ := net.LookupIP(h)

     fmt.Fprintf(w, "Hi there, I'm served from %s with IP %s !!!", h, addrs[0].To4())
}
func main() {
     http.HandleFunc("/", handler)
     http.ListenAndServe(":8484", nil)
}