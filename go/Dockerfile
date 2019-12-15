FROM golang:1.7-alpine as go-env
RUN mkdir /app

ADD main.go /app

WORKDIR /app

RUN go build main.go

FROM golang:1.7-alpine

COPY --from=go-env /app/main /app/main

ENTRYPOINT ["/app/main"]

EXPOSE 8000
