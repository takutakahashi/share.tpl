build:
	go build -o dist/cmd cmd/cmd.go
run: build
	dist/cmd src/test.txt
list: build
	dist/cmd list
show: build
	dist/cmd show src/test.txt