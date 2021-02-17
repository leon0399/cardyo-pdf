FROM golang:1.15.7-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ENV PORT=8080 \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app .


FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release

WORKDIR /root/
COPY --from=builder /go/src/app/app .

EXPOSE 8080

CMD ["./app"]


FROM builder as development

RUN go get -u github.com/beego/bee

CMD ["bee", "run"]