.PHONY: test
test:  ## run go test
	gotest -timeout=300s -race -shuffle=on -count=1 -cover `go list ./...`
