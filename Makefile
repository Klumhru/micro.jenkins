build:
	@cd cmd/jenkins-server && go build

install: build
	@cd cmd/jenkins-server && go install

run: install
	@jenkins-server --port 10000

provision:
	go get github.com/go-swagger/go-swagger/cmd/swagger

generate:
	swagger generate server -a jenkins ./swagger.json

clean:
	rm -rf cmd/ models/ restapi/ tmp/
