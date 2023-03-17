package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

var samlMiddleware *samlsp.Middleware

func main() {
	sp, err := newSamlMiddleware()
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Handle("/saml/", sp)

	http.Handle("/index", sp.RequireAccount(
		http.HandlerFunc(landingHandler),
	))
	http.Handle("/hello", sp.RequireAccount(
		http.HandlerFunc(helloHandler),
	))

	slo := http.HandlerFunc(logout)

	http.Handle("/logout", slo)

	portString := fmt.Sprintf(":%d", webserverPort)
	fmt.Println("server started at", portString)
	http.ListenAndServe(portString, nil)
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
	name := samlsp.AttributeFromContext(r.Context(), "displayName")
	w.Write([]byte(fmt.Sprintf("Welcome, %s!", name)))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	nameID := samlsp.AttributeFromContext(r.Context(), "urn:oasis:names:tc:SAML:attribute:subject-id")
	url, err := samlMiddleware.ServiceProvider.MakeRedirectLogoutRequest(nameID, "")
	if err != nil {
		panic(err) // TODO handle error
	}

	err = samlMiddleware.Session.DeleteSession(w, r)
	if err != nil {
		panic(err) // TODO handle error
	}

	w.Header().Add("Location", url.String())
	w.WriteHeader(http.StatusFound)
}
