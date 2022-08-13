ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine
ENV CGO_ENABLED=0
WORKDIR /app
RUN apk --no-cache add curl
COPY ./ ./
RUN go get -d -v ./...
CMD ["/bin/sh", "startup.sh"]