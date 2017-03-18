.PHONY: northwind

ARTISAN_VERSION?=0.0.1
ARTISAN_IMAGE=artisan:${ARTISAN_VERSION}

default: northwind

northwind:
	docker run --rm -v "`pwd`:/northwind" -w /northwind ${ARTISAN_IMAGE} make build

build:
	rm -rf bindings
	mkdir -p bindings
	protoc -o bindings/description.pb \
		--include_imports \
		--include_source_info \
		--go_out=plugins=grpc:./bindings/ \
		--proto_path=/artisan/protos \
		--proto_path=protos \
		-I $(shell find protos/ -iname *.proto)
