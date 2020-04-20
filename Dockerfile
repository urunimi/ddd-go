# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.12.0

# Change working dir -> for locating configrations
WORKDIR /gorest

# COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# Copy the local package files to the container's workspace.
COPY . .

# Run test
# RUN go test $(go list ./...)

# Build the project inside the container.
RUN go install ./cmd/gorest

# Run by default when the container starts.
CMD /go/bin/gorest

# Document that the service listens on port 8081
EXPOSE 8079