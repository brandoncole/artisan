.PHONY: default

ARTISAN_VERSION?=0.0.1
ARTISAN_IMAGE=artisan:${ARTISAN_VERSION}

default: examples

compiler:
	docker build -f artisan/Dockerfile.compiler artisan/ -t ${ARTISAN_IMAGE}

examples: compiler
	$(MAKE) -C examples/northwind