run:
	find . \( -name '*.go' -o -name '*.env' \) | entr -r go run cmd/file-server/main.go

test:
	gotestsum --format-icons="text" --format="testdox" --packages $(shell find . -type f -name '*_test.go' -exec dirname {} \; | sort -u)

