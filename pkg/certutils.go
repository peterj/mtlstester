package pkg

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

// Loads the key pair from the given file names
func LoadKeyPair(certFile string, keyFile string) (tls.Certificate, error) {
	log.Printf("Loading key pair: %s %s", certFile, keyFile)
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Show server cert information
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return tls.Certificate{}, err
	}

	log.Printf("Using certificate: CN: %s SAN: %s", cert.Leaf.Subject.CommonName, cert.Leaf.DNSNames)
	return cert, nil
}

// Loads the CA from the given file and returns a new cert pool with the CA cert in it
func CreateCACertPool(caFile string) (*x509.CertPool, error) {
	log.Printf("Loading CA: %s", caFile)
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool, nil
}
