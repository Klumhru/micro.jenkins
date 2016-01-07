build:
	@cd cmd/jenkins-server && go build

install: build
	@cd cmd/jenkins-server && go install

run: install
	@jenkins-server

provision:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

generate: provision
	swagger generate server
