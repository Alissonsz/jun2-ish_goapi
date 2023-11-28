FROM golang:1.20-alpine

WORKDIR /usr/jun2-ish/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN apk add --no-cache make

COPY . .
RUN go build -o jun2-ish_goapi
RUN chmod +x jun2-ish_goapi
RUN chmod +x entrypoint.sh

EXPOSE 8080

ENTRYPOINT [ "./entrypoint.sh" ]
