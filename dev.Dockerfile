FROM golang:latest

WORKDIR /app
ENV STAGE="dev"

# Downloading golang modules on the container build
COPY go.mod .
COPY go.sum .
RUN go mod download

# Downloading necessary tool for interactive development
RUN apt update -y 
RUN apt upgrade -y
RUN apt install -y git
RUN go install github.com/githubnemo/CompileDaemon@latest

# Starting interactive development
ENTRYPOINT CompileDaemon -polling -build="go build -buildvcs=false -o ./bin/serv ./cmd/server" -command="./bin/serv"
