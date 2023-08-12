build:
	@go build -buildmode=exe -o sigma

install:
	@mkdir -p /usr/local/bin/
	@mv sigma /usr/local/bin/sigma
