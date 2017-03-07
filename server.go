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
    var userID []string = r.Header["bns-submit-id"]
    var btcAddr []string = r.Header["bns-address"]

    log.Println("AHHHH!!!!! submit")
    fmt.Fprintf(w, "you tried to submit:\nid: %s\naddress: %s", userID, btcAddr)
}

func ResolveHandler(w http.ResponseWriter, r *http.Request) {
    var userID []string = r.Header["id"]

    log.Println("AHHHH!!!!! resolve")
    fmt.Fprintf(w, "you tried to resolve:\nid: %s", userID[0])
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
