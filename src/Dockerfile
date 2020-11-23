FROM alpine

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$PATH

RUN apk add --no-cache go git yasm && \
    go get -v -u github.com/minio/asm2plan9s && \
	go get -v -u github.com/minio/c2goasm

WORKDIR /work
ENTRYPOINT ["c2goasm"]
