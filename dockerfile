FROM golang:1.20.0-alpine

WORKDIR /app

LABEL version="0.5"
LABEL maintainer="gibran.devops@gmail.com"

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o bin/app

EXPOSE 8000

CMD [ "bin/app" ]
