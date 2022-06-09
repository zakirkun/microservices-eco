# golang alpine 1.15.6-alpine as base image
FROM golang:1.17.7-alpine AS builder

# create appuser.
RUN adduser -D -g '' syx

# create workspace
WORKDIR /opt/app/
COPY go.mod . 
COPY go.sum .

# fetch dependancies
RUN go mod download
RUN go mod verify

# copy the source code as the last step
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/services-auth ./cmd/app

# build a small image
FROM alpine:3.12.3
LABEL language="golang"
LABEL org.opencontainers.image.source github.com/zakirkun/microservices-eco

# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder --chown=syx:1000 /go/bin/services-auth /services-auth 

# use a non-root user
USER syx

ENTRYPOINT ["./services-auth"]