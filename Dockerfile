####################################################################
# Builder Stage                                                    #
####################################################################
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine AS builder

MAINTAINER saad<saad.ali@d4interactive.io>

# Create WORKDIR using project's root directory
WORKDIR /go/src/github.com/goose-rest-api
# Copy the local package files to the container's workspace
# in the above created WORKDIR

COPY . .
RUN apk add --no-cache git
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get github.com/advancedlogic/GoOse
# Build the go-API-template command inside the container
RUN cd main && go build -o main


#####################################################################
# Final Stage                                                       #
#####################################################################
# Pull golang alpine image (very small image, with minimum needed to run Go)
FROM alpine:3.7

RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        ca-certificates \
    && update-ca-certificates 2>/dev/null || true

# Create WORKDIR
WORKDIR /app

# Copy app binary from the Builder stage image
COPY --from=builder /go/src/github.com/goose-rest-api/main .

# Run the userServer command by default when the container starts.
CMD ["./main"]

# Document that the service uses port 8080
EXPOSE 8000


