package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 31 division"
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"Reading Books", "Traveling", "Buying things"},
			Info: Info{
				Affiliation: "Wayne Enterprises",
				Address:     "Gotham City",
			},
		}
		tmpl := template.Must(template.ParseFiles("./view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
