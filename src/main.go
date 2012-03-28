package main

import (
    "os"
    "fmt"
    "http"
)

func hello(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
}

func main(){
    ip := os.Getenv("OPENSHIFT_INTERNAL_IP")
    http.HandleFunc("/", hello)
    http.ListenAndServe(fmt.Sprintf("%s:8080", ip), nil)
}
