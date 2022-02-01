# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY src/go.mod ./
COPY src/go.sum ./
RUN go mod download

COPY ./src ./

RUN go build -o /gin-project


EXPOSE 8080

CMD [ "/gin-project" ]
