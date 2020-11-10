package main

import "os"

func getEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}


func main() {
	
}

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
	})
	
	http.HandleFunc("/basepoint/update/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello")
    })

    http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Path[len("/greet/"):]
        fmt.Fprintf(w, "Hello %s\n", name)
    })

    http.ListenAndServe(":9990", nil)
}