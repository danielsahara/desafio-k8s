package main

import(
	"fmt"
	"log"
	"net/http"
)

func greeting(msg string) string{
	return "<b>" + msg + "</b>"
}

func home(w http.ResponseWriter, r *http.Request){
	msg := r.URL.Query()["msg"]

	if len(msg) < 1 {
		fmt.Fprint(w, greeting("Parâmetro msg não foi informado"))
	} else {
		fmt.Fprint(w, greeting(string(msg[0])))
	}
}

func handle(){
	http.HandleFunc("/", home)
}

func main(){
	handle()

	log.Fatal(http.ListenAndServe(":8000", nil))
}