package config

import (
	"crypto/x509"
	"flag"
	"log"
	"mtlstester/pkg"
)

type Config struct {
	ServerPort string
	RunMode    string
	CACertFile string

	// The client certificate and key files
	ClientCertFile string
	ClientKeyFile  string

	// The server certificate and key files
	ServerCertFile string
	ServerKeyFile  string

	// Client request URL
	RequestURL string

	// The server certificate pool
	CACertPool *x509.CertPool
}

func (c *Config) IsServer() bool {
	return c.RunMode == "server"
}

func (c *Config) IsClient() bool {
	return c.RunMode == "client"
}

func (c *Config) IsRunModeValid() bool {
	return c.IsServer() || c.IsClient()
}

// Parses the command line arguments and returns a Config struct
func Parse() *Config {
	// Parse the config from the command line
	cfg := Config{}

	flag.StringVar(&cfg.ServerPort, "serverPort", "8443", "Server port number")
	flag.StringVar(&cfg.RunMode, "run", "client", "Run mode (server|client)")
	flag.StringVar(&cfg.CACertFile, "caCertFile", "ca-crt.pem", "CA certificate file")

	flag.StringVar(&cfg.ClientCertFile, "clientCertFile", "client-crt.pem", "Client certificate file")
	flag.StringVar(&cfg.ClientKeyFile, "clientKeyFile", "client-key.pem", "Client key file")

	flag.StringVar(&cfg.ServerCertFile, "serverCertFile", "server-crt.pem", "Server certificate file")
	flag.StringVar(&cfg.ServerKeyFile, "serverKeyFile", "server-key.pem", "Server key file")

	flag.StringVar(&cfg.RequestURL, "requestURL", "https://localhost:8443/hello", "Request URL")

	flag.Parse()

	// Load the CA's certificate
	caCertPool, err := pkg.CreateCACertPool(cfg.CACertFile)
	if err != nil {
		log.Fatal(err)
	}
	cfg.CACertPool = caCertPool

	return &cfg
}
