package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Student struct {
	Id    string
	Name  string
	Grade int8
}

var data = []Student{
	Student{"1", "Ahmad", 1},
	Student{"2", "Mardiana", 1},
	Student{"3", "Ahmad Mardiana", 2},
	Student{"4", "Mardiana Ahmad", 3},
}

func main() {
	var port = flag.String("p", "8000", "server port")
	fmt.Println(port)

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)
	http.ListenAndServe(string(":"+*port), nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var res []byte
		var err error

		for _, each := range data {
			if each.Id == id {
				res, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(res)
				return
			}
		}

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
	return
}
