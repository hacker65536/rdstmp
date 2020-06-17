

test: fmt vet 
	go test -race -coverprofile=coverage.txt -covermode=atomic

vet:
	go vet ./...


fmt:
	go fmt ./...
