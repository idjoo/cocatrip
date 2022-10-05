FROM golang:alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...


FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /usr/src/app

COPY . .

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

ENTRYPOINT ["app"]
