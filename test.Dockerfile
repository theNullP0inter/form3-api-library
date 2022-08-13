ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine
ENV CGO_ENABLED=0
WORKDIR /app
COPY ./ ./
RUN go get -d -v ./...
CMD [ "go", "test", "-v", "./..." ]