package main

import "net/http"

func main() {
	s := MakeServer()

	s.Teach("/sendMessage", PostMessage)
	s.Teach("/history", GetMessages)

	http.ListenAndServe(":8080", s)

}
