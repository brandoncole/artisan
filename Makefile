.PHONY: default compiler examples

ARTISAN_VERSION?=0.0.1
ARTISAN_IMAGE_COMPILER=artisan-compiler:${ARTISAN_VERSION}
ARTISAN_IMAGE_CLI=artisan-cli:${ARTISAN_VERSION}
ARTISAN_BINARY=${GOPATH}/bin/artisan
ARTISAN_BINARY_INPUTS=$(find artisan/* -iname *.go)

GLIDE_BINARY=${GOPATH}/bin/glide

default: examples

$(ARTISAN_BINARY): $(GLIDE_BINARY) $(ARTISAN_BINARY_INPUTS)
	glide install
	(cd artisan/cli && go build .)

$(GLIDE_BINARY):
	go get -u github.com/Masterminds/glide

compiler: compiler/Dockerfile
	docker build -f compiler/Dockerfile compiler/ -t ${ARTISAN_IMAGE_COMPILER}

cli: compiler artisan/Dockerfile
	# Compile artisan...
	docker run -v "`pwd`:/go/src/github.com/brandoncole/artisan" ${ARTISAN_IMAGE_COMPILER} make -C /go/src/github.com/brandoncole/artisan/artisan
	docker build -f artisan/Dockerfile artisan/ -t ${ARTISAN_IMAGE_CLI}

examples: cli $(ARTISAN_BINARY)
	$(MAKE) -C examples/northwind