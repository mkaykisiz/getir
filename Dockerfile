FROM golang:1.17.2-alpine AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/getir

### Stage 2
FROM scratch
COPY --from=builder /app/getir /bin/getir
ENTRYPOINT ["/bin/getir"]