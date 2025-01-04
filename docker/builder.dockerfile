ARG BUILDER_IMAGE=golang:alpine

FROM ${BUILDER_IMAGE} as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Set working directory
WORKDIR /src/app

# Copy dependencies and source code
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
    -ldflags='-w -s' -o /go/bin/app ./cmd/main.go

RUN chmod +x /go/bin/app

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy compiled binary
COPY --from=builder /go/bin/app /go/bin/app

# Use unprivileged user
USER appuser:appuser

# Set entrypoint
ENTRYPOINT ["/go/bin/app"]
