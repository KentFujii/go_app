FROM golang:1.9-stretch
RUN apt-get update && apt-get install --no-install-recommends -y \
    gcc \
    build-essential \
    ca-certificates \
    openssl \
    curl \
    git-core \
    postgresql-client

ENV GOROOT /usr/local/go
ENV GOBIN /usr/local/go/bin
ENV GOPATH /go
ENV PATH $PATH:/usr/local/go/bin

RUN go get -u github.com/derekparker/delve/cmd/dlv
RUN go get -u github.com/lib/pq
RUN go get -u github.com/jinzhu/gorm


WORKDIR /go/src
COPY . /go/src
RUN go build -o storing
CMD ["./storing"]
