FROM k1low/alpine-pandoc-ja
MAINTAINER karino2 <hogeika2@gmail.com>
RUN apk --no-cache add bash go bzr git mercurial subversion openssh-client ca-certificates musl-dev
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
WORKDIR /workspace
