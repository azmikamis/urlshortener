package main

import (
  "net/http"
  "log"
)
func main() {

  http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {

    log.Println(r.Body)

  })

  http.HandleFunc("/original", func(w http.ResponseWriter, r *http.Request) {

    log.Println(r.Body)

  })

  http.ListenAndServe(":8080", nil)

}
