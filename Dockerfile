FROM golang:1.13 as builder

ARG APP_NAME=scheduler
RUN mkdir -p /app
ADD . /app/
WORKDIR /app
RUN go get -d && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /
COPY --from=builder /app .

CMD ["/app"]
EXPOSE 8088
