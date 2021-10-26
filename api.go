package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func PostMessage(w http.ResponseWriter, r *http.Request, server Server) {
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
	}

	server.myCalls[m.To] = append(server.myCalls[m.To], m.Content)

	fmt.Fprintf(w, "%+v", m)
}

func GetMessages(w http.ResponseWriter, r *http.Request, server Server) {
	w.Header().Set("Content-Type", "application/json")
	otherEnd := r.URL.Query()["otherEnd"][0]

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(server.myCalls[otherEnd])
}
