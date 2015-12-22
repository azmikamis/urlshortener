package main

import (
  "net/http"
  "log"
  "encoding/json"
  "strings"
)

func main() {

  var shortenedurls = make(map[int]string)
  count := 0
  
  http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {

    var message map[string] interface{} // empty interface

    if err := json.NewDecoder(r.Body).Decode(&message); err == nil {
    
      for key := range message {
        if key == "url" {
            
          s := message["url"].(string)

          count += 1
          shortenedurls[count] = s

        }
      }
    }

  })

  http.HandleFunc("/expand", func(w http.ResponseWriter, r *http.Request) {

    log.Println(r.Body)

  })

  http.ListenAndServe(":8080", nil)

}
