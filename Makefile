GOPATH=$(PWD)

all: clean test build
vendor:
	gb vendor restore
clean:
	rm -Rf bin pkg bin
build:
	gb build all
test:
	gb test -v
