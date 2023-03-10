package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
)

func newSamlMiddleware() (*samlsp.Middleware, error) {
	keyPair, err := tls.LoadX509KeyPair(samlCertificatePath, samlPrivateKeyPath)
	if err != nil {
		return nil, err
		// panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		return nil, err
		// panic(err) // TODO handle error
	}

	idpMetadataURL, err := url.Parse(samlIDPMetadata)
	if err != nil {
		return nil, err
		// panic(err) // TODO handle error
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		return nil, err
		// panic(err) // TODO handle error
	}

	rootURL, err := url.Parse(webserverRootURL)
	if err != nil {
		return nil, err
		// panic(err) // TODO handle error
	}

	sp, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})
	if err != nil {
		return nil, err
	}

	return sp, nil
	// app := http.HandlerFunc(hello)
	// http.Handle("/hello", samlSP.RequireAccount(app))
	// http.Handle("/saml/", samlSP)
	// http.ListenAndServe(":8000", nil)
}
