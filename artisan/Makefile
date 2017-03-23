SOURCES=$(find . -iname *.go)
PROTOS=$(find . -iname *.proto)

bin\artsian: $(SOURCES) $(PROTOS)
	mkdir -p bin
	protoc 	--go_out=plugins=grpc:$(GOPATH)/src \
			--proto_path=api \
			api/artisan.proto
	go build -o bin/artisan .