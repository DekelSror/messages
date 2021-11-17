package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

func MockDecode(encrypted_string string) string {
	var mask uint64
	mask = 0x123456789abcdef0
	var result string

	for i := 0; i < len(encrypted_string)/8; i++ {
		piece, err := strconv.ParseUint(encrypted_string[i*8:i*8+8], 16, 64)

		if err != nil {

		}

		decrypted_piece := piece ^ mask
		result += strconv.FormatUint(decrypted_piece, 16)

	}

	fmt.Print(fmt.Sprintf("result %s \n", result))

	return result
}

func PostMessage(w http.ResponseWriter, r *http.Request, server Server) {
	var m Message

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	encrypted_string := buf.String()
	MockDecode(encrypted_string)

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
