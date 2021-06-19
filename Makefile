build:
	go build -o dist/cmd cmd/cmd.go
run: build
	DEBUG=true dist/cmd --config ./misc/config_test.yaml snippets/single
list: build
	dist/cmd --config ./misc/config_test.yaml list 
show: build
	dist/cmd  --config ./misc/config_test.yaml show snippets
dir: build
	dist/cmd  --output misc/dist snippets/project
