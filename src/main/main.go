package main

import (
  "os"
  "encoding/json"
  "log"
  "net/http"
)

type JsonHandler struct {}
type status_struct struct { Status bool }
type shot_struct struct { X string; Y string }

func (this *JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  var shot shot_struct
  err := decoder.Decode(&shot)
  if err != nil {
    panic(err)
  } else {
    response, _ := json.Marshal(&status_struct{Status: true})
    log.Println("X:", shot.X, "Y:", shot.Y)
    log.Pringln(string(response))
    w.WriteHeader(200)
    w.Write([]byte(response))
  }
}

func main() {
  log.Println("Welcome to CBattle")
  log.Println("Listen 0.0.0.0:8080")
  log.Println("PID:", os.Getpid() )

  http.Handle("/", new(JsonHandler))
  http.ListenAndServe(":8080", nil)
}
