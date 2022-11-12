FROM golang:1.19-alpine as builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./


FROM golang:1.19-alpine

ENV GIN_MODE=release

WORKDIR /usr/src/app

COPY . .

ENV READINESS_CHECK_PORT=7531
ENV PORT=7531
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.5.0 /lambda-adapter /opt/extensions/lambda-adapter

COPY --from=builder /usr/local/bin/app /

ENTRYPOINT ["/app"]
