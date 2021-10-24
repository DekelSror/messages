package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    // "strings"
)

type Message struct {
    From string `json:"from"`
    To string	`json:"to"`
    Content string `json:"content"`
}

var myCalls map[string] []string
var routes map[string] func(http.ResponseWriter, *http.Request) 

func PostMessage(w http.ResponseWriter, r *http.Request) {    
	fmt.Println("send message called")

    
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	
	if err != nil {}

    myCalls[m.To] = append(myCalls[m.To], m.Content)
	
	fmt.Fprintf(w, "%+v", m)
}


func GetMessages(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    otherEnd := r.URL.Query()["otherEnd"][0]

    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(myCalls[otherEnd])
}


func Server(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    var url = r.URL.Path

    if routes[url] != nil {
        routes[url](w, r)
    } else {
        w.WriteHeader(http.StatusNotFound)
    }

}

func main() {

    // init histories and router
    myCalls = make(map[string][]string)
    routes = make(map[string]func(http.ResponseWriter, *http.Request))
	
    // init server routes
    routes["/sendMessage"] = PostMessage
    routes["/history"] = GetMessages

    // Serve & Protect
    http.HandleFunc("/", Server)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
