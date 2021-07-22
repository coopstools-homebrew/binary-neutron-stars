package main

import (
	"encoding/json"
	"fmt"
	"github.com/coopstools-homebrew/dev-environment-controller/kube"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"net/http"
	"os"
)

var k8sctl = &kube.Kubectl{}

func main() {
	args := os.Args
	prefix := ""
	if len(args) > 2 {
		prefix = "/" + args[2]
	}

	mux := http.NewServeMux()
	mux.HandleFunc(prefix + "/", GetNamespaces)
	handler := logRequestHandler(mux)
	handler = cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "https://home.coopstools.com"},
	}).Handler(handler)
	addr := ":" + args[1]
	fmt.Println(addr)
	err := http.ListenAndServe(addr, handler)
	fmt.Printf("%+v\n", errors.Wrap(err, "could not start server"))
}

func GetNamespaces(w http.ResponseWriter, r *http.Request) {
	namespaces, err := k8sctl.ListNamespaces()
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("\nServer error getting namespaces: %+v", errors.WithStack(err))
		fmt.Fprint(w, "Internal server error")
		return
	}
	data, err := json.Marshal(namespaces)
	if err != nil {
		w.WriteHeader(500)
		fmt.Printf("\nServer error serializing namespaces: %+v", errors.WithStack(err))
		fmt.Fprint(w, "Internal server error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	count, _ := w.Write(data)
	fmt.Printf("\n%d bytes returned", count)
}

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		uri := r.URL.String()
		method := r.Method
		fmt.Printf("\n%v: %v", method, uri)
	}
	return http.HandlerFunc(fn)
}
