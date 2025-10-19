FROM golang:1.22-alpine AS builder

WORKDIR /src
RUN apk add --no-cache git ca-certificates

COPY go.mod ./
RUN go mod download

COPY . .

ARG VERSION=dev
ARG TARGETOS=linux
ARG TARGETARCH=amd64
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN go build -trimpath -ldflags="-s -w -X main.version=$VERSION" -o /out/ycl ./cmd/ycl

FROM gcr.io/distroless/static:nonroot

USER nonroot:nonroot
COPY --from=builder /out/ycl /ycl

ENTRYPOINT ["/ycl"]
