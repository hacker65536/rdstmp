

test: fmt vet 
	go test ./... -coverprofile cover.out



vet:
	go vet ./...


fmt:
	go fmt ./...
