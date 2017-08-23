package main

import "encoding/json"
import "net/http"
import "fmt"

type people struct {
    Number string
    Name  string
    Position string
}

var data = []people{
    people{"10", "Tsubasa", "Shadow Striker"},
    people{"11", "Misaki", "Mid"},
    people{"9", "Hyuga", "Striker"},
    people{"1", "Wakabayashi", "Keeper"},
}

func players(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func player(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var number = r.FormValue("number")
        var result []byte
        var err error

        for _, each := range data {
            if each.Number == number {
                result, err = json.Marshal(each)

                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                w.Write(result)
                return
            }
        }

        http.Error(w, "User not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func main() {
    http.HandleFunc("/players", players)
    http.HandleFunc("/player", player)

    fmt.Println("starting web server at http://localhost:8765/")
    http.ListenAndServe(":8765", nil)
}