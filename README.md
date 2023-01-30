# mTLS Tester

This repo contains a simple Go app that can run in client or server mode, for the purpose of showing how mTLS works.

Both modes accept the cert and private key, as well as the CA certificate.

To test the mTLS, both client and server need to use certificates issued by the CA. To generate all certificates, run the following (make sure to provide a password when prompted):

```shell
# CA cert
openssl req -x509 -sha256 -days 90 -newkey rsa:2048 -addext basicConstraints=critical,CA:TRUE,pathlen:1 -subj '/O=ACME CA Inc./CN=CA' -keyout ca-key.pem -out ca-crt.pem

# Server cert
openssl genrsa -out server-key.pem 4096
openssl req -new -subj "/CN=server" -key server-key.pem -out server-csr.pem
openssl x509 -req -days 90 -in server-csr.pem -CA ca-crt.pem -CAkey ca-key.pem -CAcreateserial -out server-crt.pem   -extensions SAN -extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:localhost"))

# Client cert
openssl genrsa -out client-key.pem 4096
openssl req -new -subj "/CN=client" -key client-key.pem -out client-csr.pem
openssl x509 -req -days 90 -in client-csr.pem -CA ca-crt.pem -CAkey ca-key.pem -CAcreateserial -out client-crt.pem   -extensions SAN -extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:localhost"))
```

If you use the above snippet and the default file names, you can run the app without providing any extra flags.

To run the server:

```shell
go run main.go --run server
```

```console
2023/01/30 15:01:29 Loading CA: ca-crt.pem
2023/01/30 15:01:29 Running client
2023/01/30 15:01:29 Loading key pair: server-crt.pem server-key.pem
2023/01/30 15:01:29 Using certificate: CN: server SAN: [localhost]
2023/01/30 15:01:29 Listening for requests on port 8443
```

In a separate terminal, run the client that will make an mTLS request with the server:

```shell
go run main.go --run client
```

```console
2023/01/30 15:02:28 Loading CA: ca-crt.pem
2023/01/30 15:02:28 Loading key pair: client-crt.pem client-key.pem
2023/01/30 15:02:28 Using certificate: CN: client SAN: [localhost]
2023/01/30 15:02:28 Making request to https://localhost:8443/hello
Received response: Hello world%
```

On the server-side, you'll notice the request was received and the client certificate details:

```console 
2023/01/30 15:02:28 Request received
2023/01/30 15:02:28     Client certificate: CN: client SAN: [localhost]
```

## Running using Docker

You can also run the app in the Docker containers. We use `host` network, so the client can call the server using `localhost`. Theoretically, you could use other host name, but then you'll have to make sure to include the DNS names in the `subjectAltName` when creating the certificates.


The commands assume you have the certs created in the same folder where you're running these commands.

To run the server:

```shell
docker run --net host -v $PWD:/certs ghcr.io/peterj/mtlstester:latest --run server --caCertFile certs/ca-crt.pem --serverCertFile certs/server-crt.pem --serverKeyFile certs/server-key.pem
```

The above command exposes the server on the default port (`8443`) from the host.

You can run the client like this:

```shell
docker run --net host -v $PWD:/certs ghcr.io/peterj/mtlstester:latest --run client --caCertFile certs/ca-crt.pem --clientCertFile certs/client-crt.pem --clientKeyFile certs/client-key.pem
```
