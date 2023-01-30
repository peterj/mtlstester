package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"mtlstester/config"
	"mtlstester/pkg"
	"net/http"
)

func Run(cfg *config.Config) {
	log.Println("Running client")

	// Read the certs and make a request to the server
	cert, err := pkg.LoadKeyPair(cfg.ClientCertFile, cfg.ClientKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      cfg.CACertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	// Create a request with the host header
	req, err := http.NewRequest("GET", cfg.RequestURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make the request
	log.Printf("Making request to %s", cfg.RequestURL)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received response: %s", string(body))
}
