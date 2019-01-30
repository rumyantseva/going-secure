package main

import (
	"fmt"
	"net/http"
	"html/template" // HL
)

func main() {
	tpl := template.Must(template.ParseFiles("examples/xss/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		data := struct {
			Name string
		}{
			Name: name,
		}
		tpl.Execute(w, data)
	})
	fmt.Println("Ready to listen and serve...")
	http.ListenAndServe(":3000", nil)
}
