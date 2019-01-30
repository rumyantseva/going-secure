FROM artifactory/golang-linters

RUN mkdir -p /go/src/github.com/rumyantseva/myapp
ADD . /go/src/github.com/rumyantseva/myapp
WORKDIR /go/src/github.com/rumyantseva/myapp

# Run linters
RUN golangci-lint run \
	--no-config --issues-exit-code=1 \
	--deadline=10m --exclude-use-default=false \
	--enable-all \
	./...

# Run tests
RUN go test -timeout=600s -v --race ./...
