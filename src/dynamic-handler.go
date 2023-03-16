package main

import(
  "fmt"
  "strings"
  "net/http"
)

func dynamicHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	if path[2] == "bar" {
		fmt.Fprint(w, "you have reach foo bar!!!")
	} else {
		fmt.Fprintln(w, "error")
	}

}
