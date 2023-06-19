.PHONY:

build: 
	echo "🔨 Building executables for http and cli apps..."
	go build -o bin/http ./cmd/http
	go build -o bin/cli ./cmd/cli

