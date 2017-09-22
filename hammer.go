package main

import "bytes"
import "encoding/json"
import "fmt"
import "github.com/google/uuid"
import "io/ioutil"
import "net/http"
import "os"
import "sync"
import "time"

type RenoMessage struct {
    Event string    `json:"event"`
    ID string       `json:"id"`
    Timestamp int64 `json:"timestamp"`
}

func send(url string) {
    for {
        m := &RenoMessage{ "lsdvrs", uuid.New().String(), time.Now().Unix() }
        js, _ := json.Marshal(m)
        fmt.Println("Sending json to ReNo: ", string(js))
        response, _ := http.Post(url, "application/json", bytes.NewReader(js))
        body, _ := ioutil.ReadAll(response.Body)
        fmt.Println("    => ", string(body))
        time.Sleep(3 * time.Second)
    }
}

func main() {
    url := os.Getenv("RENO_URL")
    fmt.Println("Using RENO_URL: ", url)
    fmt.Printf("hello, hammer")
    var wg sync.WaitGroup 
    wg.Add(1)
    go send(url)
    wg.Wait()
}
