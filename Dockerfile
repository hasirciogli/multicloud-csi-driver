FROM golang:1.23.4 AS builder

WORKDIR /app

COPY . .

RUN go build -o csi-driver .

FROM golang:1.23.4

WORKDIR /app

COPY --from=builder /app/csi-driver .

CMD ["./csi-driver"]
