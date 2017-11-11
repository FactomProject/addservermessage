FROM golang:1.9.2-alpine as builder

# Get git
RUN apk add --no-cache curl git

# Get glide
RUN go get github.com/Masterminds/glide

# Where addservermessage sources will live
WORKDIR $GOPATH/src/github.com/FactomProject/addservermessage

# Get the dependencies
COPY glide.yaml glide.lock ./

# Install dependencies
RUN glide install -v

# Populate the rest of the source
COPY . .

ARG GOOS=linux

# Build and install addservermessage
RUN go install

# Now squash everything
FROM alpine:3.6

# Get git
RUN apk add --no-cache ca-certificates curl git

RUN mkdir -p  /go/bin
COPY --from=builder /go/bin/addservermessage /go/bin/addservermessage

ENTRYPOINT ["/go/bin/addservermessage"]

