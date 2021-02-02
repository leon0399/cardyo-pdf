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


FROM builder as development

RUN go get -u github.com/beego/bee

CMD ["bee", "run"]


FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /build/app .

EXPOSE 4000

CMD ./app
