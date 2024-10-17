FROM golang:1.23 as build

WORKDIR ~/go/src/github.com/asirolly/grpctest
COPY . /go/src/app

RUN go mod download
RUN CGO_ENABLED=0 go build cmd/main.go -o /go/bin/main

FROM gcr.io/distroless/static-debian12
LABEL authors="Andrew Sirolly"

COPY --from=build /go/bin/main /

HEALTHCHECK CMD curl localhost:8080

CMD ["/main"]