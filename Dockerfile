FROM golang:1.12-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build host/main.go

# Start fresh from a smaller image
FROM alpine:3.9 
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/main /app/main

RUN mkdir -p /root/oscal_processing_space/tmp
RUN mkdir -p /root/oscal_processing_space/uploads
RUN mkdir -p /root/oscal_processing_space/downloads
COPY testprofile.xml /root/oscal_processing_space/downloads/testprofile.xml
COPY testssp.xml /root/oscal_processing_space/downloads/testssp.xml
COPY container_root_files /root

# This container exposes port 8080 to the outside world
EXPOSE 9050

CMD ["/app/main"]
