package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

    "github.com/fatih/structs"
)


func dynamicHandler(w http.ResponseWriter, r *http.Request) {
    path := strings.Split(r.URL.Path, "/")

    db, err := ConnectDatabse(8080, "root", "root", "test", "test")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "An error occurred whilst connecting to the database")
        return
    }
    defer db.Close()

    switch path[2] {
        case "get": {
            var users []User
            err = db.Select("users", &users)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, "An error occurred getting table users")
                return
            }
            usersJson, err := json.Marshal(users)
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, "An error occurred getting table users")
                return
            }
            w.Header().Set("Content-Type", "application/json")
            w.Write(usersJson)
        }; break

        case "del": {
            err = db.Delete("users")
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                fmt.Fprintf(w, "An error occurred getting table users")
                return
            }
        }; break

        case "new": {
            var users []interface{}
            userData, err := db.db.Create("users", structs.Map(User {
                Forename: "Jonathan",
                Surname: "Leeming",
                Role: President,
                EntryYear: 2022,
                Interests: []Interest { "programming", "software", "hardware" },
            }))
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            users = append(users, userData)

            userData, err = db.db.Create("users", structs.Map(User {
                Forename: "Aaron",
                Surname: "Kelly",
                Role: VicePresident,
                EntryYear: 2022,
                Interests: []Interest { "dev-ops", "networking", "hardware" },
            }))
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            users = append(users, userData)

            userData, err = db.db.Create("users", structs.Map(User {
                Forename: "Max",
                Surname: "Friedrich",
                Role: Secretary,
                EntryYear: 2022,
                Interests: []Interest { "dev-ops", "networking", "programming", "software" },
            }))
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            users = append(users, userData)

            userData, err = db.db.Create("users", structs.Map(User {
                Forename: "Aaron",
                Surname: "Kelly",
                Role: VicePresident,
                EntryYear: 2022,
                Interests: []Interest { "dev-ops", "networking", "hardware" },
            }))
            if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            users = append(users, userData)
        }; break
    }
}

