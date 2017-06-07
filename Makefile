.PHONY: all linux darwin

all: linux darwin

clean:
	rm -f jsonctl-Linux-* jsonctl-Darwin-*

linux:
	env CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jsonctl-Linux-x86_64 ./cmd

darwin:
	env CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o jsonctl-Darwin-x86_64 ./cmd
