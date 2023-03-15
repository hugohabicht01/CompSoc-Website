package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"
)


func execHandler(w http.ResponseWriter, r *http.Request) {
    db, err := ConnectDatabse(8080, "root", "root", "test", "test")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "An error occurred whilst connecting to the database")
        fmt.Println(err)
        return
    }
    defer db.Close()

    EXEC_ROLES := []string{
        "President",
        "Vice-President",
        "Secretary",
        "Treasurer",
        "Commissioner for Continuity",
        "Education Officer",
        "Events Officer",
        "Publicity Officer",
        "Socials Officer",
        "Technical Officer",
        "Welfare Officer",
    }

    var users []User
    err = db.Select("users", &users)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Couldn't load users")
        fmt.Println(err)
        return
    }

    execMembers := make(map[string]User)
    for _, role := range EXEC_ROLES {
        execMembers[strings.ReplaceAll(strings.ReplaceAll(role, "-", "_"), " ", "_")] = User {
            Forename: "Open",
            Surname: "",
            Role: role,
            EntryYear: 0,
        }
    }
    for _, member := range users {
        if contains(EXEC_ROLES, member.Role) {
            execMembers[strings.ReplaceAll(strings.ReplaceAll(member.Role, "-", "_"), " ", "_")] = member
        }
    }

    templateFuntions := template.FuncMap {
        "formatEntryYear": func(entryYear int) string {
            if entryYear <= 0 {
                return ""
            }
            now := time.Now()
            currentYear := now.Year()
            if now.Month() > 6 {
                currentYear += 1
            }
            duration := currentYear - entryYear
            switch duration {
                case 1:
                    return "First Year"
                case 2:
                    return "Second Year"
                case 3:
                    return "Third Year"
                case 4:
                    return "Fourth Year"
                case 5:
                    return "Fifth Year"
                default:
                    return fmt.Sprintf("%dth", duration)
            }
        },
    }

    t, err := template.New("exec.html").Funcs(templateFuntions).ParseFiles("../pages/exec.html")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "Error loading template")
        fmt.Println(err)
        return
    }
    t.Execute(w, map[string]any {
        "Exec": execMembers,
        "IsHome": false,
        "Site": map[string]any {
            "Title": "LUCompSoc",
        },
        "Page": map[string]any {
            "Title": "Executive Committee",
        },
        "Params": map[string]any {
            "hide_site_title": false,
        },
        "Description": "Meet the Executive Committee of the Lancaster University Computer Science Society",
        "Keywords": []string {
            "Lancaster University",
            "Computer Science",
            "University Society",
            "Executive Committee",
        },
        "Permalink": "https://compsoc.io/exec",
    })
}


func contains[T comparable](s []T, x T) bool {
    for _, y := range s {
        if x == y {
            return true
        }
    }
    return false
}

