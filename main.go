package main

import (
	"html/template"
	"net/http"
	"os"
	// "github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1> Hello world$$$</h1>")) // response data to client
	tpl.Execute(w, nil)
}
func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }
	port := os.Getenv("PORT") // set up port
	if port == "" {
		port = "3000"
	}
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)

}
