package main

import "net/http"
import "./router"
import "log"
//import "io/ioutil"

func main() {
	
	http.HandleFunc("/gpa", router.GPAQuery) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	

}