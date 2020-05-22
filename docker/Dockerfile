FROM golang
WORKDIR /work_lyft/tasks/docker
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/golang/example/outyet
ENTRYPOINT /go/bin/outyet
EXPOSE 8080

