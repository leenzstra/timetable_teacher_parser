build: 
	go build -o build/parser.exe ./cmd/main.go 

build_docker:
	go build -o build/parser_docker.exe ./cmd/main.go 