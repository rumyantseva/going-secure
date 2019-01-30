# Stage 0. Build the binary
FROM artifactory/golang:1.11

# add a non-privileged user
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/rumyantseva/myapp
ADD . /go/src/github.com/rumyantseva/myapp
WORKDIR /go/src/github.com/rumyantseva/myapp

# build the binary with go build
RUN CGO_ENABLED=0 go build \
	-o bin/myapp github.com/rumyantseva/myapp/cmd/myapp

# Stage 1. Run the binary
FROM scratch

ENV PORT 8080

# certificates to interact with other services
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# don't forget /etc/passwd from previous stage
COPY --from=0 /etc/passwd /etc/passwd
USER myapp

# and finally the binary
COPY --from=0 /go/src/github.com/rumyantseva/myapp/bin/myapp /myapp
EXPOSE $PORT

CMD ["myapp"]
