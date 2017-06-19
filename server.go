package main

import (
	"net/http"
    "log"
    "fmt"
    "io/ioutil"
    "time"
    "bytes"
)

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("bns-submit-id")
    btcAddr := r.Header.Get("bns-address")

    if userID == "" {
        fmt.Println("userID is empty!")
    } else {
        fmt.Println("userID is " + userID)
    }

    if btcAddr == "" {
        fmt.Println("btcAddr is empty!")
    } else {
        fmt.Println("btcAddr is " + btcAddr)
    }

    fmt.Fprintf(w, "you tried to submit:\nid: %s\naddress: %s\n", userID, btcAddr)
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("bns-resolve-id")

    if userID == "" {
        fmt.Println("userID is empty!")
    } else {
        fmt.Println("userID is " + userID)
    }

    fmt.Fprintf(w, "you tried to resolve:\nid: %s\n", userID)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

    indexFile, err := ioutil.ReadFile("index.html")

    if err != nil {
        fmt.Fprint(w, err)
        return
    }

    http.ServeContent(w, r, "index.html", time.Now(), bytes.NewReader(indexFile))
}

func main() {

    http.HandleFunc("/submit", SubmitHandler)
    http.HandleFunc("/resolve", ResolveHandler)
    http.HandleFunc("/", IndexHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
