package main

import (
  "net/http"
  "log"
  "encoding/json"
  "strings"
  "strconv"
  "sync"
  "bytes"
)

func main() {

  var shortenedurls = make(map[int]string)
  var mutex = &sync.Mutex{}
  count := 0
  
  http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {

    var message map[string] interface{} // empty interface

    if err := json.NewDecoder(r.Body).Decode(&message); err == nil {
    
      for key := range message {
        if key == "url" {
            
          s := message["url"].(string)

          mutex.Lock()
          count += 1
          shortenedurls[count] = s
          i := strconv.Itoa(count)
          mutex.Unlock()
          
          var buffer bytes.Buffer
          buffer.WriteString("http://localhost/")
          buffer.WriteString(i)
          response, _ := json.Marshal(map[string]string{"short": buffer.String()})
          w.Write([]byte(response))

        }
      }
    }

  })

  http.HandleFunc("/expand", func(w http.ResponseWriter, r *http.Request) {

    var message map[string] interface{} // empty interface

    if err := json.NewDecoder(r.Body).Decode(&message); err == nil {
    
      for key := range message {
        if key == "short" {
            
          s := strings.Replace(message["short"].(string),"http://localhost/","", -1)
          i, _ := strconv.Atoi(s)
          
          var buffer bytes.Buffer
          buffer.WriteString(shortenedurls[i])
          response, _ := json.Marshal(map[string]string{"expand": buffer.String()})
          w.Write([]byte(response))
        }
      }
    }
    
  })

  log.Fatal(http.ListenAndServe(":8080", nil))

}
