FROM golang:1.17-alpine

ARG APP_PORT

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

EXPOSE ${APP_PORT}

ENTRYPOINT ["./entrypoint.sh"]
