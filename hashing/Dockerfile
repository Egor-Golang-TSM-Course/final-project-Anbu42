FROM golang:1.21 AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hashing ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /build/hashing .

EXPOSE 5051

ENTRYPOINT ["./hashing"]