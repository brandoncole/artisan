.PHONY: northwind build

ARTISAN_VERSION?=0.0.1
ARTISAN_IMAGE_CLI=artisan-cli:${ARTISAN_VERSION}

NORTHWIND_PACKAGE=github.com/brandoncole/artisan/examples/northwind
NORTHWIND_PACKAGE_DIR="/go/src/${NORTHWIND_PACKAGE}"

default: northwind

northwind:
	docker run --rm -v "`pwd`:${NORTHWIND_PACKAGE_DIR}" -w ${NORTHWIND_PACKAGE_DIR} ${ARTISAN_IMAGE_CLI} make build

build:
	artisan generate -f artisan.yaml
	protoc -o artisan/schema/description.pb \
		--include_imports \
		--include_source_info \
		--go_out=plugins=grpc:./artisan/grpc/ \
		--proto_path=/artisan/include \
		--proto_path=protos \
		-I $(shell find protos/ -iname *.proto)
