build:
	go build -o dist/cmd cmd/cmd.go
run: build
	DEBUG=true dist/cmd --config ./misc/config_test.yaml snippets/single
list: build
	dist/cmd --config ./misc/config_test.yaml list 
show: build
	dist/cmd  --config ./misc/config_test.yaml show snippets/single
dir: build
	dist/cmd --config ./misc/config_test.yaml --output misc/dist snippets/project
