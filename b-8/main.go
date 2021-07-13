package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var funcMap = template.FuncMap{
	"unescape": func(s string) template.HTML {
		return template.HTML(s)
	},
	// "unescape": func(s string) string {
	// 	return s
	// },
	"avg": func(n ...int) int {
		total := 0
		for _, each := range n {
			total += each
		}
		return total / len(n)
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("view.html").Funcs(funcMap).
			ParseFiles("view.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
