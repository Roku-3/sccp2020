package main

import (
    "fmt"
   	"bytes"
    "net/http"
)

//curl -X POST -d "POSTしたい文字列" http://localhost:8080/todo
//bufbody := new(bytes.Buffer)
//bufbody.ReadFrom(r.Body)
//body := bufbody.String()

//var todoList map[int]string = map[int]string{}
var todoList []string

func main() {

    http.HandleFunc("/", HogeHandler)
    http.HandleFunc("/todo", TodoHandler)
    todoList = append(todoList, "hoge1")
    todoList = append(todoList, "hoge2")
    todoList = append(todoList, "hoge3")
    todoList = append(todoList, "hoge4")


    /*router.POST("/new", func(c *gin.Context) {
        title := c.PostForm("title")
        p := c.PostForm("price")
        price, err := strconv.Atoi(p)
        if err != nil {
            panic(err)
        }
        dbInsert(title, price)
        c.Redirect(302, "/")
    })*/

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
        fmt.Fprint(w, "POST hello!\n")
        bufbody := new(bytes.Buffer)
        bufbody.ReadFrom(r.Body)
        body := bufbody.String()
    switch r.Method {
    case "GET":
        fmt.Fprint(w, "GET hello!\n")
    case "POST":
        todoList = append(todoList, body)
    case "PUT":
    case "DELETE":
        result := []string{}
        for _, v:= range todoList {
            if v != body {
                result = append(result, v)
            }
        }
        todoList = result
    default:
        fmt.Fprint(w, "Method not allowed.\n")
    }
    for _, v:= range todoList {
        fmt.Fprint(w, v + "\n")
    }
}

