package main

import "sync"

// Item represents each JSON object sent by clients
type Item map[string]interface{}

// Root represents the root node with a string parameter and a list of Items
type Root struct {
    Name  string `json:"name"`
    Items []Item `json:"items"`
}

// Global root node and mutex to protect it
var (
    root Root
    mu   sync.Mutex
)
