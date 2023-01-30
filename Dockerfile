FROM golang:1.19.5-alpine3.17 AS builder

# Copy the go source and build it
WORKDIR /go/src/mtlstester
COPY . .
RUN go build -o /go/bin/mtlstester

# Copy the binary to the production image from the builder stage.
FROM alpine:3.17
COPY --from=builder /go/bin/mtlstester /app/mtlstester
ENTRYPOINT ["/app/mtlstester"]

