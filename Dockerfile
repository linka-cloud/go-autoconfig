FROM golang:alpine as builder

RUN apk add git ca-certificates

WORKDIR /go-autoconfig

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o autoconfig .

FROM alpine

COPY --from=builder /etc/ca-certificates /etc/ca-certificates
COPY --from=builder /go-autoconfig/autoconfig /usr/bin/autoconfig

EXPOSE 1323

ENTRYPOINT ["/usr/bin/autoconfig"]