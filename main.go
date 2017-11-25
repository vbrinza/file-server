package main

import (
	"flag"
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP server")
	var staticDir, passwordFile, listenPort string
	flag.StringVar(&staticDir, "static-dir", "", "The directory containing static files to be served.")
	flag.StringVar(&passwordFile, "password-file", "", "The path to password file.")
	flag.StringVar(&listenPort, "p", ":8080", "The listening port.")
	flag.Parse()
	if len(staticDir) == 0 {
		log.Fatalf("static-dir must be provided")
	}
	if len(passwordFile) == 0 {
		log.Fatalf("password-file must be provided")
	}
	log.Printf("static-dir = %s", staticDir)
	log.Printf("password-file = %s", passwordFile)
	secrets := auth.HtpasswdFileProvider(passwordFile)
	authenticator := auth.NewBasicAuthenticator("Authentication", secrets)
	http.HandleFunc("/", authenticator.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(staticDir)).ServeHTTP(res, &req.Request)
	}))

	http.ListenAndServe(listenPort, nil)
}
