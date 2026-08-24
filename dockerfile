FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build \
	-trimpath \
	-ldflags="-s -w" \
	-o /out/fsquare ./cmd/app

FROM scratch

COPY --from=builder /out/fsquare /fsquare

USER 65532:65532
ENTRYPOINT ["/fsquare"]
