package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"time"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/facebook/get_post/{post_id}", FBServiceHandler)

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

func FBServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_id := vars["post_id"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello from the Facebook service! Fetching post with ID %q\n", post_id)
	fmt.Fprintf(w, "Finding test header value: %q\n", r.Header.Get("x-ext-auth-token"))
}
