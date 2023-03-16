package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cubewise-code/go-mime"
)

const ERRORPAGE = "../public/404/index.html"

func checkPath(path string) (string, error) {
	path = filepath.Join("../public", path)
	pathStat, err := os.Stat(path)
	if err != nil {
		return "", nil
	} else if pathStat.IsDir() {
		return filepath.Join(path, "index.html"), nil
	}
	return path, nil
}

func staticHandler(w http.ResponseWriter, localPath string, statusCode int) {
	data, err := ioutil.ReadFile(localPath)
	if err != nil {
		staticHandler(w, ERRORPAGE, 404)
	}

	mime := gomime.TypeByExtension(filepath.Ext(localPath))
	w.Header().Set("Content-Type", mime)
	if statusCode == 404 {

		w.WriteHeader(statusCode)
	}
	fmt.Fprint(w, string(data))
}

func router(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	localPath, err := checkPath(path)
	if err != nil {
		staticHandler(w, ERRORPAGE, 404)
	} else if !dynamicRoute(w, path, localPath) {
		staticHandler(w, localPath, 200)
	}
}
