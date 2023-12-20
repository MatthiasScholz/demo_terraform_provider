# NOTE Terraform Provider Naming convention (prevents usage of "-")
provider_name := demotpgen
provider_src := ./internal/provider

setup: prerequisites init
prerequisites: prerequisites-basics-check prerequisites-partial-setup

scaffold: scaffold-provider

# -- General go stuff

init:
	go mod init $(provider_name)

deps:
	go mod tidy

binary := build/$(provider_name)
build: deps
	@echo "INFO :: Build destination: $(binary)"
	go build -o $(binary)

# -- Generator specifics

prerequisites-basics-check:
	go version
	terraform -version

prerequisites-partial-setup:
	go install github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework@latest
	go install github.com/hashicorp/terraform-plugin-codegen-openapi/cmd/tfplugingen-openapi@latest

scaffold-provider:
	mkdir -p $(provider_src)
	tfplugingen-framework scaffold provider --name $(provider_name) --output-dir $(provider_src)
