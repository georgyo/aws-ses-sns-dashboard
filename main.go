package main

import (
    "net/http"
)

func main() {
    router := NewRouter()
    http.ListenAndServe(":8000", router)
}
