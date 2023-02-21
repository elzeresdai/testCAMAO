FROM golang:1.19.0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /camao-test

EXPOSE 8080

CMD [ "/camao-test" ]