package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        handleGet(w, r)
    case http.MethodPost:
        handlePost(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGet(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    mu.Lock()
    defer mu.Unlock()

    json.NewEncoder(w).Encode(root)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
    var item Item

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Unable to read body", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    if err := json.Unmarshal(body, &item); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    mu.Lock()
    root.Items = append(root.Items, item)
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Item added successfully")
}
