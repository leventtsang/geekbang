package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "GeekBang-k8s-lesson1")
}

func sayno(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "no hello")
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问路由
	//创建监听端口
	http.HandleFunc("/no", sayno) //设置访问路由
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}
}