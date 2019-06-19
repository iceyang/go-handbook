OUTPUT_DIR="_output"

all: test build
build:
	go build -i -o ${OUTPUT_DIR}/bin/server -v cmd/server.go
test:
	go test -v ./...
clean:
	go clean
	rm -rf ${OUTPUT_DIR}
run:
	go build -i -o ${OUTPUT_DIR}/bin/server -v cmd/server.go
	./${OUTPUT_DIR}/bin/server
