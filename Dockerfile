FROM golang:1.11-alpine AS builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -a -tags netgo -ldflags '-w -extldflags "-static"' -o main .


FROM scratch as final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /
ENV GO_CONFIG /config/pom.yaml
ENV GO_SECRET /secret/pom.yaml
USER nobody:nobody
CMD ["./main"]

