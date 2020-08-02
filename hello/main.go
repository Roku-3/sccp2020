package main

import (
	"fmt"
//	"bytes"
	"net/http"
)

//curl -X POST -d "POSTしたい文字列" http://localhost:8080/todo
//bufbody := new(bytes.Buffer)
//bufbody.ReadFrom(r.Body)
//body := bufbody.String()

	var todoList map[int]string = map[int]string{}

func main() {

	todoList[0] = "一つ目" 
    todoList[1] = "二つ目"
	todoList[2] = "三つ目"
	todoList[3] = "四つ目"
	http.HandleFunc("/", HogeHandler)
	http.HandleFunc("/hoge", TodoHandler)
    todoList = append(todoList, "hoge")

	// 8000ポートで起動
	http.ListenAndServe(":8080", nil)
}
func HogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("アクセスされたよ/n")
	fmt.Fprint(w, "hoge")
}
//func HelloHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Print("アクセスされたよ/n")
//	fmt.Fprint(w, "hello world")
//}
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	for i:=0; i<4; i++ {
		fmt.Fprint(w, todoList[i] + "\n")
	}
}
