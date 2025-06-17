FROM golang:1.24

# install nftables and netcat for debugging
RUN apt-get update && apt-get install -y nftables netcat-openbsd

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "--build.cmd", "go build -o /usr/local/bin/efw ./cmd/efw/", "--build.bin", "/usr/local/bin/efw", "--build.args_bin", "status"]