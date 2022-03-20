// package main

// import "net/http"

// func main() {
// 	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("hello world !!!!"))
// 	})
// 	http.ListenAndServe(":80", nil)
// }

package main

import (
	"net/http"

	"github.com/mabukar/golang-learning/projects/HTTP_Server/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
