FROM golang:1.25-alpine AS build

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/server ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates wget

WORKDIR /app
COPY --from=build /out/server /app/server
COPY migrations /app/migrations

ENV PORT=8000
EXPOSE 8000

USER 65532:65532

HEALTHCHECK --interval=10s --timeout=3s --start-period=15s --retries=3 \
    CMD wget -qO- "http://127.0.0.1:${PORT}/health" >/dev/null || exit 1

ENTRYPOINT ["/app/server"]