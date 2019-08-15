package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/instagram/get_post/{post_id}", IGServiceHandler)

	http.Handle("/", r)

	server := newServer(":"+strconv.Itoa(8080), r)

	panic(server.ListenAndServe())
}

func newServer(addr string, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
}

func IGServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["post_id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from the Instagram service! Fetching post with ID %q\n", post_id)
	fmt.Fprintf(w, "Finding Auth1: %q\n", r.Header.Get("x-ext-auth1"))
	fmt.Fprintf(w, "Finding Auth2: %q\n", r.Header.Get("x-ext-auth2"))
	fmt.Fprintf(w, "Finding Social Network Auth1: %q\n", r.Header.Get("x-ext-sn-auth1"))
	fmt.Fprintf(w, "Finding Social Network Auth2: %q\n", r.Header.Get("x-ext-sn-auth2"))
	fmt.Fprintf(w, "Finding Social Network Auth3: %q\n", r.Header.Get("x-ext-sn-auth3"))
	fmt.Fprintf(w, "Finding Social Network App ID: %q\n", r.Header.Get("x-ext-sn-app-id"))
}
