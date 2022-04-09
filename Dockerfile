FROM golang:1.18 AS builder

ENV GO111MODULE=on \
    GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"

WORKDIR $GOPATH/src/app

# manage dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app .

FROM alpine:latest  
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache --update --virtual .build-deps \
    ca-certificates \
    curl \
    git \
    go \
    linux-headers \
    make \
    openssl \
    pcre \
    zlib \
    chromium

WORKDIR /root/
COPY --from=builder /app ./
EXPOSE 8080
CMD ["/app"]