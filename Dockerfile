# syntax=docker/dockerfile:1

FROM golang:1.18-buster AS build

WORKDIR /usr/src/job

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/job cmd/job/*.go

FROM chromedp/headless-shell:latest

RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/bin/job /usr/local/bin/job
CMD ["/usr/local/bin/job"]



