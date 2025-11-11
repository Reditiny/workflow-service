# ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

RUN apk --no-cache add ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o myapp ./cmd/main.go

# ---------- Runtime Stage ----------
FROM public.ecr.aws/amazonlinux/amazonlinux:2023

WORKDIR /app

COPY --from=builder /app/myapp .

RUN chmod +x myapp

EXPOSE 8765

CMD ["./myapp"]

