package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type UserId int

type User struct {
    name string
    id UserId
}

type Message struct {
    from UserId
    to UserId
    content string
}

var users []User = []User{ {name: "fekel", id: 1}, {name: "odan", id: 0} }
var nextUserId = 2

var messages map[string][]string

func PostUser(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    userName := q["userName"][0]

    if userName == "" {
        http.Error(w, "shoot", http.StatusBadRequest)
        return
    }

    for i := 0; i < len(users); i++ {
        if users[i].name == userName {
            http.Error(w, "username exists", http.StatusBadRequest)
            return
        }

    }

    users = append(users, User{name: userName, id: UserId(nextUserId)})
    nextUserId++
    fmt.Fprintf(w, "user %s added!", userName)
    http.StatusText(http.StatusOK)
}

func StartSession(w http.ResponseWriter, r *http.Request) {



}

func ConversationId (from string, to string) string {
    if from >= to {
        return to + from
    } else {
        return from + to
    }
}



func SendMessage(w http.ResponseWriter, r *http.Request) {    
    var m Message
    error := json.NewDecoder(r.Body).Decode(&m)

    if error != nil {
        fmt.Fprintf(w, error.Error())
        http.Error(w, "bad message", http.StatusBadRequest)
        return
    }

    q := r.URL.Query()
    from := q["fromUser"][0]
    to := q["toUser"][0]
    
    conversationId := ConversationId(from, to)    

    if messages[conversationId] == nil {
        messages[conversationId] = []string{m.content}
    } else {
        messages[conversationId] = append(messages[conversationId], m.content)
    }
}


func GetHistory (w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    from := q["fromUser"][0]
    to := q["toUser"][0]
    
    conversationId := ConversationId(from, to)

    mstr, err := json.Marshal(messages[conversationId])

    if err != nil {return}

    fmt.Fprintf(w, "%s", mstr)    
}

func main() {
    http.HandleFunc("/addUser", PostUser)
    
    http.HandleFunc("/listUsers", func(w http.ResponseWriter, r *http.Request) {
        var res string
        
        for i := 0; i < len(users); i++ {
            res += users[i].name + " "
        }

        fmt.Fprintf(w, "%s", res)
    })


    http.HandleFunc("/viewHistory", GetHistory)

    log.Fatal(http.ListenAndServe(":8080", nil)) // Serve & Protect
}
