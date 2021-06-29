build:
	go build -o dist/cmd cmd/cmd.go
run: build
	DEBUG=true dist/cmd --config ./example/global_config.yaml snippets/single
list: build
	dist/cmd --config ./example/global_config.yaml list 
show: build
	dist/cmd  --config ./example/global_config.yaml show snippets/single
dir: build
	dist/cmd --config ./example/global_config.yaml snippets/project
update: build
	dist/cmd --config ./example/global_config.yaml update
exec: build
	dist/cmd --config ./example/global_config.yaml exec snippets2/exec-test
