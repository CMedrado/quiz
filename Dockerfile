FROM golang:1.23.5
LABEL maintainer=rafaelmedrado
WORKDIR app
COPY . .
RUN go mod tidy
RUN go build -o api app/main.go
CMD ["./api"]
