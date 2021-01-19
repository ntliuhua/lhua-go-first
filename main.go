package main

import (
	"fmt"
	"net/http"

	"github.com/ntliuhua/lhua-go-first/controller"
)

func main() {
	http.HandleFunc("/index", controller.Index)
	http.ListenAndServe(":8100", nil)
	fmt.Println("liuhua-go-web")
}
