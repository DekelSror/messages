package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"

)

type Message struct {
    From string `json:"from"`
    To string	`json:to`
    Content string `json:content`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {    
	fmt.Println("send message called")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	
	if err != nil {}
	
	fmt.Fprintf(w, "%+v", m)
}


func main() {
	fmt.Println("hi!")
    http.HandleFunc("/sendMessage", SendMessage)
    log.Fatal(http.ListenAndServe(":8080", nil)) // Serve & Protect
}
