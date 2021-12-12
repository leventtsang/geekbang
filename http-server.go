package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	timer := metrics.NewTimer()
    	defer timer.ObserveTotal()
   	delay := randInt(10, 2000)
    	time.Sleep(time.Millisecond * time.Duration(delay))
	
	//从请求中获取HTTP标头并插入response headers
	for k, v := range r.Header {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
		v1 := v[0] //转换v的[]string类型为string
		w.Header().Set(k, v1)
	}

	w.Header().Set("Version", "Go 1.17.1 Windows/AMD64") //在response headers插入version信息

	fmt.Fprintf(w, "GeekBang-k8s-lesson1") //输出页面内容
}

func http4xx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You f*ck up")
}

func http5xx(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I f*ck up")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}

func randInt(min int, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return min + rand.Intn(max-min)
}

func main() {

	http.HandleFunc("/", sayhelloName)   //设置根目录路由
	http.HandleFunc("/404", http4xx)     //设置4xx路由-练手未完成，需设置正则
	http.HandleFunc("/504", http5xx)     //设置5xx路由-练手未完成，需设置正则
	http.HandleFunc("/healthz", healthz) //设置心跳检测路径

	err := http.ListenAndServe(":9090", nil) //创建监听端口
	if err != nil {
		log.Fatal("ListenAndServe:	", err)
	}

}
