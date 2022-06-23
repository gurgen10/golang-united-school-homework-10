package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", NameHandler).Methods("Get")
	router.HandleFunc("/bad", BadHandler).Methods("Get")
	router.HandleFunc("/data", DataHandler).Methods("Post")
	router.HandleFunc("/header", HeaderHandler).Methods("Get")
	http.Handle("/", router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}
func NameHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", param)

}
func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "%v", http.StatusInternalServerError)

}
func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	w.WriteHeader(http.StatusOK)
	if len(body) > 0 {
		fmt.Fprintf(w, "`I got message:\n%v`", string(body))
	}

}
func HeaderHandler(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")

	log.Printf("Headers: a: %s, b: %s", a, b)
	aVal, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	bVal, err1 := strconv.Atoi(b)
	if err1 != nil {
		log.Fatal(err1)
	}

	w.Header().Set("a+b", strconv.Itoa(aVal+bVal))
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
