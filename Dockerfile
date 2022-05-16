# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

WORKDIR /usr/src/job

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/job cmd/job/*.go

FROM chromedp/headless-shell:latest

RUN apk chromedp/headless-shell:latest
ENV TZ=Europe/Stockholm

COPY --from=build /usr/local/bin/job /usr/local/bin/job

CMD ["/usr/local/bin/job"]



