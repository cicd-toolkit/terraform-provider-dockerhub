terraform-provider-dockerhub: *.go */*.go go.mod docs/index.md test
	go build .

test:
	terraform fmt -recursive
	go fmt ./...
	go vet .
	go test ./...

testacc: test
	TF_ACC=1 go test ./...

install: terraform-provider-dockerhub
	-rm .terraform.lock.hcl
	terraform init

docs/index.md: $(shell find -name "*.go" -or -name "*.tmpl" -or -name "*.tf")
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
