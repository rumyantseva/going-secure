FROM artifactory/golang:1.11

# Change here to update
ENV VERSION 1.12.3
ENV CHECKSUM c531688661b500d4c0c500fcf57f829388a4a9ba79697c2e134302aedef0cd46

# Make sure we have a fixed golangci-lint script with a chekcsum check
RUN echo "${CHECKSUM}  golangci-lint-${VERSION}-linux-amd64.tar.gz" > CHECKSUM

# Download from Github the specified release and extract into the go/bin folder
RUN curl -L "https://github.com/golangci/golangci-lint/releases/download/v${VERSION}/golangci-lint-${VERSION}-linux-amd64.tar.gz" \
	-o golangci-lint-${VERSION}-linux-amd64.tar.gz \
	&& shasum -a 256 -c CHECKSUM \
	&& tar xvzf golangci-lint-${VERSION}-linux-amd64.tar.gz \
		--strip-components=1 \
		-C ./bin \
		golangci-lint-${VERSION}-linux-amd64/golangci-lint

# Clean up
RUN rm -rf CHECKSUM "golangci-lint-${VERSION}-linux-amd64.tar.gz"
