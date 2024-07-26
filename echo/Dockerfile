FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GOMAXPROCS=8

RUN go get -d -v

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]


# docker run --name myechoapp -d -p 8080:8080 my-echo-app

# docker run -d --cpus="2.0" my-app

# wrk -t12 -c400 -d30s http://localhost:8080/