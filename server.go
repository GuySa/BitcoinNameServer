package main

import (
	"net/http"
    "log"
    "fmt"
    "io/ioutil"
    "time"
    "bytes"
)

func SubmitHandler(writer http.ResponseWriter, req *http.Request) {
    userID := req.Header.Get("bns-submit-id")
    btcAddr := req.Header.Get("bns-address")

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

    fmt.Fprintf(writer, "you tried to submit:\nid: %s\naddress: %s\n", userID, btcAddr)
}

func ResolveHandler(writer http.ResponseWriter, req *http.Request) {
    userID := req.Header.Get("bns-resolve-id")

    if userID == "" {
        fmt.Println("userID is empty!")
    } else {
        fmt.Println("userID is " + userID)
    }

    fmt.Fprintf(writer, "you tried to resolve:\nid: %s\n", userID)
}

func IndexHandler(writer http.ResponseWriter, req *http.Request) {

    indexFile, err := ioutil.ReadFile("index.html")

    if err != nil {
        fmt.Fprint(writer, err)
        return
    }

    http.ServeContent(writer, req, "index.html", time.Now(), bytes.NewReader(indexFile))
}

func main() {

    http.HandleFunc("/submit", SubmitHandler)
    http.HandleFunc("/resolve", ResolveHandler)
    http.HandleFunc("/", IndexHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
