test:
	go test ./...

test-coverage:
	go test -cover ./... -v
	go test -coverprofile=coverage.out ./...

test-report: test-coverage
	go tool cover -html=coverage.out -o coverage.html
