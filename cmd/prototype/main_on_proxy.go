package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// os.Setenv("HTTP_PROXY", "")
	fmt.Println(os.Args)
	r,e := http.Get("http://www.google.com")
	fmt.Println(r,e)
}
