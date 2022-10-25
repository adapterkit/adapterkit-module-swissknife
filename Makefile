test: generate
	go mod tidy
	go test -v .
	@echo "Done."

generate: swissknife.pb.go
.PHONY: generate

swissknife.pb.go: swissknife.proto
	protoc --go_out=./ --go-grpc_out=./ $<

clean:
	rm -f swissknife.pb.goa
