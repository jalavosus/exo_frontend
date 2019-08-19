package main

import (
    "net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, "dist/index.html");
    });
  http.ListenAndServe(":3000", nil)
}
