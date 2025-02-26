#build binary file
FROM golang:1.20.14-alpine as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /bin/shop
#run shop server in container
FROM alpine:latest
COPY --from=builder /bin/shop /bin/shop
EXPOSE 8080
CMD ["/bin/shop"]