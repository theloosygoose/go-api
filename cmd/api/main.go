package main

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi"
    "github.com/theloosygoose/go-api/internal/handlers"
    log "github.com/sirupsen/logrus"

)

func main(){
    log.SetReportCaller(true)
    var r *chi.Mux = chi.NewRouter()

    handlers.Handler(r)
    
    fmt.Println("Starting Go API")

    err := http.ListenAndServe("localhost:8080", r)
    if err != nil {
        log.Error(err)
    }
}
