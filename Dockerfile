FROM golang:1.16-alpine AS builder
ARG TAGGED
LABEL tagged=$TAGGED
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOGC=off GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o ./app ./cmd/main.go

FROM alpine:3.14
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /build/app /app/app
COPY --from=builder /build/.env /app/.env
WORKDIR /app
ENTRYPOINT ["/app"]