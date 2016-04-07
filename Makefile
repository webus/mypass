GOPATH=$(PWD)

all: clean test build
vendor:
	gb vendor restore
clean:
	rm -RIf bin pkg bin
build:
	gb build all
test:
	gb test -v
