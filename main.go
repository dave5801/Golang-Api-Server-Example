package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Initialize the root node
    root = Root{
        Name:  "My Root Node",
        Items: []Item{},
    }

    http.HandleFunc("/", handler)

    port := "8080"
    fmt.Printf("Starting server at port %s\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
