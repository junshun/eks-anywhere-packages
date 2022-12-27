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
COPY cmd/ cmd/
COPY api/ api/
COPY controllers/ controllers/
COPY config/ config/
COPY pkg/ pkg/

# Build
ENV GOTRACEBACK=single
ARG SKAFFOLD_GO_GCFLAGS
RUN GOPROXY=direct CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -a -o package-manager main.go

CMD ["/workspace/package-manager", "server"]
