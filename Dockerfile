# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/rushtehrani/healthz

# Build the healthz command inside the container.
RUN go install github.com/rushtehrani/healthz

# Run the healthz command by default when the container starts.
ENTRYPOINT /go/bin/healthz

EXPOSE 5000