FROM --platform=linux/amd64 golang:latest AS builder

WORKDIR /go/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/articleArena-server .

FROM debian:11.5

RUN apt update \
    && apt install -y --no-install-recommends \
    ca-certificates \
    git \
    bash \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /usr/app/

COPY --from=builder /go/src/app/bin .

EXPOSE 8010

ENTRYPOINT ["./articleArena-server"]
