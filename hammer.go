package main

import "bytes"
import "encoding/json"
import "fmt"
import "github.com/google/uuid"
import "github.com/davecgh/go-spew/spew"
import "io/ioutil"
import "net/http"
import "time"

type RenoMessage struct {
    Event string    `json:"event"`
    ID string       `json:"id"`
    Timestamp int64 `json:"timestamp"`
}

func main() {
    fmt.Printf("hello, hammer")
    m := &RenoMessage{ "lsdvrs", uuid.New().String(), time.Now().Unix() }
    spew.Dump(m)
    js, _ := json.Marshal(m)
    fmt.Println(string(js))
    response, _ := http.Post("http://httpbin.org/post", "application/json", bytes.NewReader(js))
    body, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(body))
}
