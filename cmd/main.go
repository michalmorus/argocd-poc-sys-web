package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		var data []byte

		res, err := http.Get(os.Getenv("API_URL"))
		if err == nil {
			data, _ = io.ReadAll(res.Body)
		}
		tmpl.Execute(w, map[string]string{
			"Data": string(data),
		})
	})

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

	log.Println("listen on ", port)
}
