package main

import (
       "log"
    "net/http"
    "strings"
    "path"
)

func Handler(publicDir string) http.Handler {
	handler := http.FileServer(http.Dir(publicDir))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_path := req.URL.Path

		// static files
		if strings.Contains(_path, ".") || _path == "/" {
			handler.ServeHTTP(w, req)
			return
		}

		// the all 404 gonna be served as root
		http.ServeFile(w, req, path.Join(publicDir, "/index.html"))
	})
}

func main() {
     server := &http.Server{
     	    Addr:	":3003",
	    Handler:	Handler("./dist"),
     }
     err := server.ListenAndServe()
     if err != nil {
     	log.Println(err)
     }
}
