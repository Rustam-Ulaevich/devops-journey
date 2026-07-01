package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Привет, мир DevOps! Мой первый контейнер на Go 🐳")
    })

    fmt.Println("Сервер запущен на порту 8080")
    http.ListenAndServe(":8080", nil)
}
