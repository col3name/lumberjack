package main

import (
        "log"
        "net/http"
        "os"
)

func main() {
        fs := http.FileServer(http.Dir("./web"))
        http.Handle("/", fs)

    port := os.Getenv("PORT")
        if len(port) == 0 {
                port = "3000"
        }
        log.Println("Listening on :"+ port + "...")
        err := http.ListenAndServe(":" + port, nil)
        if err != nil {
                log.Fatal(err)
        }
}