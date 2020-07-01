FROM golang:1.7
EXPOSE 8080
RUN mkdir -p /app
WORKDIR /app
ADD . /app
RUN go build ./app.go
CMD ["./app"]