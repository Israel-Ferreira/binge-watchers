FROM golang:1.19-buster

WORKDIR /usr/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/binge-watchers-app

EXPOSE 8080

CMD [ "bin/binge-watchers-app" ]


