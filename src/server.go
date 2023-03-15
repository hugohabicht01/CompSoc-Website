package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)



func checkForPublic() bool {
	_, err := os.Stat("../public")
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.Println(err)
		return false
	}

}

func serverStart() {
	http.HandleFunc("/foo/", dynamicHandler)

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	fmt.Println("Running on http://localhost:8000")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	if checkForPublic() {
		serverStart()
	}
}
