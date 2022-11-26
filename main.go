package main

import (
	"net/http"
)

// func main() {
// 	fmt.Println(uuid.New().String())
// }

func main() {
	//create routes
	http.HandleFunc("/", homehandler)
	http.HandleFunc("/ivomiho", userhandler)

	//start the server
	http.ListenAndServe(":3000", nil)
}

func homehandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func userhandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ivo Miho ZONDA37"))
}
