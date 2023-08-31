FROM golang:1.21-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download && go mod verify

COPY ../../google   ./google
COPY ../../internal ./internal
COPY ../../proto    ./proto

RUN mkdir -p /usr/local/bin/

RUN go mod tidy

RUN go build -v -o /usr/local/bin/app ./internal/cmd

CMD ["app"]