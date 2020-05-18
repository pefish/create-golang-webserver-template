FROM golang:1.14
WORKDIR /app
ENV GO111MODULE=on
COPY ./ ./
RUN GOMAXPROCS=4 go test -timeout 90s -race ./...
RUN make
CMD ["./build/bin/linux/main", "--help"]

# docker build -t pefish/main:v1.2.4 .