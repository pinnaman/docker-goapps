package main

import 
(
"fmt"
"net/http"
"log"
"github.com/gorilla/mux" 
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    //fmt.Println(w, string(9786))
    w.Write([]byte(string(9786)))
    fmt.Fprintln(w, "Smiley,")
    w.Write([]byte("Docker Gorilla!\n"))
}

func main() {

// welcome with a smiley
//fmt.Println(string(9786))

 r := mux.NewRouter()
  // Routes consist of a path and a handler function.
 r.HandleFunc("/", indexHandler)

  // Bind to a port and pass our router in
 log.Fatal(http.ListenAndServe(":8080", r))
}
