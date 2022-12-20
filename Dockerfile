# Build the manager binary
FROM golang:1.18 as builder
ENV GOTRACEBACK=all
ARG SKAFFOLD_GO_GCFLAGS

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
ENV GOPROXY=direct
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY config/ config/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -a -o package-manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/package-manager .
USER 65532:65532

CMD ["/package-manager", "server"]
