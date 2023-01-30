package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"mtlstester/config"
	"mtlstester/pkg"
	"net/http"
)

func Run(cfg *config.Config) {
	log.Println("Running server")

	// Load the server's certificate and private key
	cert, err := pkg.LoadKeyPair(cfg.ServerCertFile, cfg.ServerKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs:    cfg.CACertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
	}
	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on provided port with the TLS config
	server := &http.Server{
		Addr:      fmt.Sprintf(":%s", cfg.ServerPort),
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		// Show client cert information
		cert := r.TLS.PeerCertificates[0]
		log.Printf("\tClient certificate: CN: %s SAN: %s", cert.Subject.CommonName, cert.DNSNames)

		fmt.Fprintf(w, "Hello world")
	})

	// Listen to HTTPS connections with the server certificate and wait
	log.Printf("Listening for requests on port %s", cfg.ServerPort)
	log.Fatal(server.ListenAndServeTLS("", ""))
}
