ARG GO_VERSION=1.13
FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git
ENV CGO_ENABLED=0
WORKDIR /app
COPY ./ ./
RUN go get -d -v ./...
RUN go build \
    -installsuffix 'static' \
    -o /form3 .

FROM scratch AS final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /form3 /form3
USER nobody:nobody
ENTRYPOINT ["/form3"]