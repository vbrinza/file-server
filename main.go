package main

import (
  "flag"
  auth "github.com/abbot/go-http-auth"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "time"
)

func main()  {
  var passwordFile, staticDir string
  flag.StringVar(&staticDir, "static-dir", "", "The directory containing static files")
  flag.StringVar(&passwordFile, "password-file", "htpasswd", "The path to the password file")
  flag.Parse()
  if len(staticDir) == 0 {
    log.Fatalf("static-dir must ne provided")
  }
  log.Printf("staticDir = %s", staticDir)
  secrets := auth.HtpasswdFileProvider(passwordFile)
  authenticator := auth.NewBasicAuthenticator("gitrest", secrets)

  r := mux.NewRouter().StrictSlash(false)
  r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))
  http.Handle("/", authenticator.Wrap(func(w http.ResponseWriter, ar *auth.AuthenticatedRequest) {
    r.ServeHTTP(w, &ar.Request)
  }))
  s := &http.Server{
    Addr:           ":8000",
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Fatal(s.ListenAndServe())
}
