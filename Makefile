build:./src/main.go
	rm -rf dist/
	cd src && go get 
	go build -o dist/v2rpc src/main.go
	cp -r ./static ./dist/