default: install

build:
	@mkdir -p build/
	@go build -o build/ipl ipl.go
install:
	@go get github.com/kavirajk/ipl
test:
	@go test -v ./...
clean:
	@rm -f ipl
