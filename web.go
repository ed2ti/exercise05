//# *************************** #
//# College: Trebas Institute 
//# Professor: Ghazal G.Fard
//# Class: Edward
//# Day: 2022-06-21
//# *************************** #

package main
import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

