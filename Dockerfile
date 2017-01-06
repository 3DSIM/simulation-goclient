FROM golang:1.7

ENV APPLICATION_NAME=simulation-goclient

# Install glide
WORKDIR $GOPATH/src/github.com/Masterminds
RUN git clone https://github.com/Masterminds/glide.git
RUN cd glide && git checkout v0.12.3 && make bootstrap-dist && make install

# Install application
RUN mkdir -p $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
COPY . $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
WORKDIR $GOPATH/src/github.com/3dsim/$APPLICATION_NAME
RUN glide install --strip-vendor

# Run tests
RUN go test $(glide novendor) -cover
