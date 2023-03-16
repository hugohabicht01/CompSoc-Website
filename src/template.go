package main

import (
	"html/template"
	"net/http"

)

func templateFill(pageData interface{}, path string, out http.ResponseWriter) error {

	templ, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	// debugFile, err := os.OpenFile("debug.html", os.O_WRONLY, os.ModePerm)
	// defer debugFile.Close()
	// if err != nil {
	// 	return err
	// }
	err = templ.Execute(out, pageData)
	if err != nil {
		return err
	}
	return nil
}
