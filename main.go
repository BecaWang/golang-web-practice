package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func InputHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	fmt.Println("hello world")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 启动静态文件服务

	http.HandleFunc("/input", InputHandler)
	err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
