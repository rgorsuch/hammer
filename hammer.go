package main

import "fmt"
import "time"
import "github.com/google/uuid"
import "github.com/davecgh/go-spew/spew"

type RenoMessage struct {
    Event string
    ID string
    Timestamp int64
}


func main() {
    fmt.Printf("hello, hammer")
    m := RenoMessage{ "lsdvrs", uuid.New().String(), time.Now().Unix() }
    spew.Dump(m)
}
