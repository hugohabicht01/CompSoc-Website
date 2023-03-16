package main

import (
	"net/http"
	"strings"
)

func dynamicRoute(w http.ResponseWriter, path string, localPath string) bool {
	if strings.HasPrefix(path, "/mytest") {
		localPath, err := checkPath("/mytest")
		if err != nil {
			staticHandler(w, ERRORPAGE, 404)
			return true
		}
		var myPageData struct {
			MyContent string
		}
		myPageData.MyContent = "Hello, World!"

		err = templateFill(myPageData, localPath, w)
		if err != nil {
			staticHandler(w, ERRORPAGE, 404)
			return true
		}
	} else {
		return false
	}
	return true
}
