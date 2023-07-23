FROM golang:1.18

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
COPY Makefile .

# Install app dependencies
RUN make install

# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.GitCommit=$GIT_COMMIT" app/main.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose

# Run
# this EXPOSE is just docs, not expose any at all
EXPOSE 8080 9090
ENTRYPOINT CONFIG_PATH=$CONFIG_PATH ./main