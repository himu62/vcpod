MIGRATION_DIRECTORY := db/migrations
MIGRATION_BINDATA := db/bindata.go

.PHONY: dep/tools dep build run test

dep/tools:
	go get -u -v github.com/golang/dep/cmd/dep
	go get -u -v github.com/shuLhan/go-bindata/...

dep:
	dep ensure

build: $(MIGRATION_BINDATA)
	go build -o vcpod vcpod.go

$(MIGRATION_BINDATA): $(MIGRATION_DIRECTORY)
	go-bindata -pkg db -o $@ $</

run:
	go run vcpod.go

test:
	go test -v -count=1 ./...
