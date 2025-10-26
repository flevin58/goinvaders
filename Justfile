generate:
	@go generate

install: generate
	@go install

