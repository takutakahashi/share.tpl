build:
	go build -o dist/cmd cmd/cmd.go
run: build
	DEBUG=true dist/cmd src
list: build
	dist/cmd list
show: build
	dist/cmd show src
dir: build
	dist/cmd  --output src/dist src/dirtest
